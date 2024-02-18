package blockchain

import (
	"context"
	"errors"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/Exca-DK/pegism/core/backoff"
	"github.com/Exca-DK/pegism/core/log"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

type subBlockListener struct {
	client blockClient

	log                log.Logger
	backoff            backoff.Backoff
	minimalBackoff     time.Duration
	currentBlockHeight atomic.Uint64

	blockFunc func(*types.Block)
}

func newSubBlockListener(client blockClient, notify func(*types.Block)) (*subBlockListener, error) {
	currentHeader, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	listener := &subBlockListener{
		client:         client,
		log:            log.Root(),
		minimalBackoff: 5 * time.Second,
		backoff:        backoff.NoopBackoff,
		blockFunc:      notify,
	}
	listener.currentBlockHeight.Store(currentHeader.Number.Uint64())
	return listener, nil
}

func (listener *subBlockListener) Start(ctx context.Context) error {
	verifyError := func(err error, ctx context.Context) error {
		if errors.Is(err, ctx.Err()) {
			return nil
		} else {
			return err
		}
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ch := make(chan error, 1)
	go func() { ch <- listener.catchupBlocks(ctx) }()
	select {
	case <-ctx.Done():
		cancel()
	case err := <-ch:
		return verifyError(err, ctx)
	}
	return verifyError(ctx.Err(), ctx)
}

func (listener *subBlockListener) catchupBlocks(ctx context.Context) error {
	ch := make(chan *types.Header, 512)
	sub := event.Resubscribe(5*time.Minute, func(ctx context.Context) (event.Subscription, error) {
		listener.log.Info("Reconnecting to client")
		return listener.client.SubscribeNewHead(ctx, ch)
	})
	defer sub.Unsubscribe()
	for {
		select {
		case err := <-sub.Err():
			return err
		case header := <-ch:
			go func(header *types.Header) {
				if err := listener.catchupBlock(ctx, header.Number.Uint64()); err != nil {
					listener.log.Warn(
						"Failed catching block",
						"block", header.Number.Uint64(),
						"err", err,
					)
				}
			}(header)
		}
	}

}

func (listener *subBlockListener) catchupBlock(ctx context.Context, number uint64) error {
	listener.log.Trace("Calling block by numer", "number", number)
	block, err := listener.client.BlockByNumber(ctx, big.NewInt(int64(number)))
	if err != nil {
		return err
	}
	listener.notifyNewBlock(ctx, block)
	return nil
}

func (listener *subBlockListener) notifyNewBlock(ctx context.Context, block *types.Block) {
	listener.log.Trace("Sending new block notification", "number", block.NumberU64())
	listener.currentBlockHeight.Store(block.NumberU64())
	listener.blockFunc(block)
}
