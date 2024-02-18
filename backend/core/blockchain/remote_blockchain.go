package blockchain

import (
	"context"
	"errors"
	"math/big"
	"sync"

	"github.com/Exca-DK/pegism/core/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
)

type blockListener interface {
	Start(context.Context) error
}

type txClient interface {
	TransactionByHash(
		ctx context.Context,
		hash common.Hash,
	) (tx *types.Transaction, isPending bool, err error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}

type blockClient interface {
	ChainID(ctx context.Context) (*big.Int, error)
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)
}

type BlockchainClient interface {
	blockClient
	txClient
}

type RemoteBlockchain struct {
	client BlockchainClient
	signer types.Signer

	blockListener blockListener

	done chan struct{}
	stop chan struct{}
	wg   sync.WaitGroup

	blockSubs *event.SubscriptionScope

	blockFeed event.FeedOf[*types.Block]

	log log.Logger
}

func NewBlockchain(client BlockchainClient) (*RemoteBlockchain, error) {
	id, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &RemoteBlockchain{
		client:    client,
		signer:    types.LatestSignerForChainID(id),
		done:      make(chan struct{}),
		stop:      make(chan struct{}),
		blockSubs: &event.SubscriptionScope{},
		log:       log.Root(),
	}, err
}

func (b *RemoteBlockchain) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan *types.Header, 1024)
	sub, err := b.client.SubscribeNewHead(ctx, ch)
	if err == nil {
		sub.Unsubscribe()
		err = b.runSubBlockListener(ctx)
	}
	if err != nil {
		if errors.Is(err, rpc.ErrNotificationsUnsupported) {
			err = b.runHttpBlockListener(ctx)
		}
	}
	if err != nil {
		cancel()
		close(b.done)
		return err
	}

	go func() {
		defer cancel()
		<-b.stop
		close(b.done)
		// close all tracked subs
		subs := b.blockSubs.Count()
		b.log.Info("Closing block subscriptions", "amount", subs)
		b.blockSubs.Close()
	}()
	return nil
}

func (b *RemoteBlockchain) runHttpBlockListener(ctx context.Context) error {
	b.log.Info("Running blockchain with pull based mode")
	listener, err := newHttpBlockListener(b.client, func(block *types.Block) {
		b.blockFeed.Send(block)
	})
	if err != nil {
		return err
	}
	b.blockListener = listener
	b.wg.Add(1)
	go func() {
		if err := listener.Start(ctx); err != nil {
			b.log.Warn("http listener start error", "err", err)
		} else {
			b.log.Debug("http listener went offline")
		}
		b.wg.Done()
	}()
	return nil
}

func (b *RemoteBlockchain) runSubBlockListener(ctx context.Context) error {
	b.log.Info("Running blockchain with subscription based mode")
	listener, err := newSubBlockListener(
		b.client,
		func(block *types.Block) { b.blockFeed.Send(block) },
	)
	if err != nil {
		return err
	}
	b.blockListener = listener
	b.wg.Add(1)
	go func() {
		if err := listener.Start(ctx); err != nil {
			b.log.Warn("Sub listener start error", "err", err)
		}
		b.wg.Done()
	}()
	return nil
}

func (b *RemoteBlockchain) Stop() {
	select {
	case <-b.done:
	case b.stop <- struct{}{}:
	}
	b.wg.Wait()
}

func (b *RemoteBlockchain) IsMined(ctx context.Context, hash common.Hash) bool {
	_, err := b.client.TransactionReceipt(ctx, hash)
	return err == nil
}

func (b *RemoteBlockchain) FetchTransaction(
	ctx context.Context,
	hash common.Hash,
) (*types.Transaction, error) {
	tx, _, err := b.client.TransactionByHash(ctx, hash)
	return tx, err
}

func (b *RemoteBlockchain) FetchReceipt(
	ctx context.Context,
	hash common.Hash,
) (*types.Receipt, error) {
	receipt, err := b.client.TransactionReceipt(ctx, hash)
	return receipt, err
}

func (b *RemoteBlockchain) Sender(tx *types.Transaction) (common.Address, error) {
	return b.signer.Sender(tx)
}

func (b *RemoteBlockchain) Header(
	ctx context.Context,
	blockHash common.Hash,
) (*types.Header, error) {
	header, err := b.client.HeaderByHash(ctx, blockHash)
	if err != nil {
		return nil, err
	}
	return header, nil
}

func (b *RemoteBlockchain) SubscribeNewBlock(ch chan<- *types.Block) event.Subscription {
	return b.blockSubs.Track(b.blockFeed.Subscribe(ch))
}
