package blockchain

import (
	"context"
	"errors"
	"math/big"
	"sync"
	"testing"
	"time"

	mockBlockchain "github.com/Exca-DK/pegism/mocks/core/blockchain"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBoostrapSub(t *testing.T) {
	feedCh := make(chan<- *types.Block)
	feed := event.FeedOf[*types.Block]{}

	client := mockBlockchain.NewBlockchainClient(t)
	var wg sync.WaitGroup
	wg.Add(1)
	client.EXPECT().
		SubscribeNewHead(mock.Anything, mock.Anything).
		Return(feed.Subscribe(feedCh), nil)
	client.EXPECT().ChainID(mock.Anything).Return(big.NewInt(10), nil)
	client.EXPECT().
		HeaderByNumber(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, i *big.Int) (*types.Header, error) {
			wg.Done()
			return &types.Header{
				Number: big.NewInt(10),
				Time:   uint64(time.Now().Unix()),
			}, nil
		})
	blockchain, err := NewBlockchain(client)
	require.NoError(t, err)
	go blockchain.Start()
	wg.Wait()
	blockchain.Stop()
}

func TestBootstrapHttp(t *testing.T) {
	client := mockBlockchain.NewBlockchainClient(t)
	var wg sync.WaitGroup
	wg.Add(3)
	times := 0
	ts := time.Now()
	client.EXPECT().
		SubscribeNewHead(mock.Anything, mock.Anything).
		Return(nil, rpc.ErrNotificationsUnsupported)
	client.EXPECT().ChainID(mock.Anything).Return(big.NewInt(10), nil)
	client.EXPECT().
		HeaderByNumber(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, i *big.Int) (*types.Header, error) {
			defer func() {
				wg.Done()
				times += 1
			}()
			if times == 0 {
				return &types.Header{
					Number: big.NewInt(10),
					Time:   uint64(ts.Unix()),
				}, nil
			} else if times == 1 {
				return &types.Header{
					Number: big.NewInt(10),
					Time:   uint64(ts.Unix()),
				}, nil
			} else if times == 2 {
				return &types.Header{
					Number: big.NewInt(9),
					Time:   uint64(ts.Add(-1).Unix()),
				}, nil
			}
			return nil, errors.New("unexpected call")
		}).Times(3)

	blockchain, err := NewBlockchain(client)
	require.NoError(t, err)
	go blockchain.Start()
	wg.Wait()
	blockchain.Stop()
}
