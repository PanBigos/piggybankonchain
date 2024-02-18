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
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestBlockCatchingUp(t *testing.T) {
	client := mockBlockchain.NewBlockchainClient(t)
	calledHeaders := 0
	calledBlocks := 0
	var wg sync.WaitGroup
	wg.Add(1)
	ts := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client.EXPECT().
		HeaderByNumber(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, i *big.Int) (*types.Header, error) {
			defer func() {
				calledHeaders += 1
			}()
			if calledHeaders == 0 {
				return &types.Header{Number: big.NewInt(10)}, nil
			} else if calledHeaders == 1 {
				require.NotNil(t, i)
				require.Equal(t, uint64(10), i.Uint64())
				return &types.Header{Time: uint64(ts.Add(-1 * time.Second).Unix())}, nil
			} else if calledHeaders == 2 {
				require.NotNil(t, i)
				require.Equal(t, uint64(9), i.Uint64())
				return &types.Header{Time: uint64(ts.Unix())}, nil
			} else if calledHeaders == 3 {
				require.Nil(t, i)
				return &types.Header{Number: big.NewInt(12)}, nil
			} else if calledHeaders == 4 {
				require.Nil(t, i)
				return &types.Header{Number: big.NewInt(14)}, nil
			} else if calledHeaders == 5 {
				require.Nil(t, i)
				wg.Done()
				cancel()
				return &types.Header{Number: big.NewInt(14)}, nil
			}
			return nil, errors.New("unexpected call to header by number")
		}).Times(6)
	client.EXPECT().
		BlockByNumber(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, i *big.Int) (*types.Block, error) {
			defer func() {
				calledBlocks += 1
			}()
			if calledBlocks == 0 {
				require.NotNil(t, i)
				require.Equal(t, uint64(10), i.Uint64())
				return types.NewBlockWithHeader(&types.Header{Number: big.NewInt(10)}), nil
			} else if calledBlocks == 1 {
				require.NotNil(t, i)
				require.Equal(t, uint64(11), i.Uint64())
				return types.NewBlockWithHeader(&types.Header{Number: big.NewInt(11)}), nil
			} else if calledBlocks == 2 {
				require.NotNil(t, i)
				require.Equal(t, uint64(12), i.Uint64())
				return types.NewBlockWithHeader(&types.Header{Number: big.NewInt(12)}), nil
			} else if calledBlocks == 3 {
				require.NotNil(t, i)
				require.Equal(t, uint64(13), i.Uint64())
				return types.NewBlockWithHeader(&types.Header{Number: big.NewInt(13)}), nil
			} else if calledBlocks == 4 {
				require.NotNil(t, i)
				require.Equal(t, uint64(14), i.Uint64())
				return types.NewBlockWithHeader(&types.Header{Number: big.NewInt(14)}), nil
			}
			return nil, errors.New("unexpected call to block by number")
		}).Times(5)

	feed := event.FeedOf[*types.Block]{}
	listener, err := newHttpBlockListener(client, func(block *types.Block) { feed.Send(block) })
	listener = listener.WithMinimal(0)
	require.NoError(t, err)

	wg.Add(1)
	defer wg.Wait()

	blockCh := make(chan *types.Block)
	sub := feed.Subscribe(blockCh)
	defer sub.Unsubscribe()
	go func() {
		require.NoError(t, listener.Start(ctx))
		wg.Done()
	}()

	for i := 10; i <= 12; i++ {
		block := <-blockCh
		require.Equal(t, uint64(i), block.NumberU64())
	}

	for i := 13; i <= 14; i++ {
		block := <-blockCh
		require.Equal(t, uint64(i), block.NumberU64())
	}
}

func TestReorg(t *testing.T) {
	client := mockBlockchain.NewBlockchainClient(t)
	calledHeaders := 0
	calledBlocks := 0
	var wg sync.WaitGroup
	wg.Add(1)
	ts := time.Now()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	client.EXPECT().
		HeaderByNumber(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, i *big.Int) (*types.Header, error) {
			defer func() {
				calledHeaders += 1
			}()
			if calledHeaders == 0 {
				return &types.Header{Number: big.NewInt(10)}, nil
			} else if calledHeaders == 1 {
				require.NotNil(t, i)
				require.Equal(t, uint64(10), i.Uint64())
				return &types.Header{Time: uint64(ts.Add(-1 * time.Second).Unix())}, nil
			} else if calledHeaders == 2 {
				require.NotNil(t, i)
				require.Equal(t, uint64(9), i.Uint64())
				return &types.Header{Time: uint64(ts.Unix())}, nil
			} else if calledHeaders == 3 {
				require.Nil(t, i)
				return &types.Header{Number: big.NewInt(8)}, nil
			} else if calledHeaders == 4 {
				require.Nil(t, i)
				return &types.Header{Number: big.NewInt(8)}, nil
			} else if calledHeaders == 5 {
				require.Nil(t, i)
				wg.Done()
				cancel()
				return &types.Header{Number: big.NewInt(10)}, nil
			}
			return nil, errors.New("unexpected call to header by number")
		}).Times(6)
	client.EXPECT().
		BlockByNumber(mock.Anything, mock.Anything).
		RunAndReturn(func(ctx context.Context, i *big.Int) (*types.Block, error) {
			defer func() {
				calledBlocks += 1
			}()
			if calledBlocks == 0 {
				require.NotNil(t, i)
				require.Equal(t, uint64(8), i.Uint64())
				return types.NewBlockWithHeader(&types.Header{Number: big.NewInt(8)}), nil
			} else if calledBlocks == 1 {
				require.NotNil(t, i)
				require.Equal(t, uint64(9), i.Uint64())
				return types.NewBlockWithHeader(&types.Header{Number: big.NewInt(9)}), nil
			} else if calledBlocks == 2 {
				require.NotNil(t, i)
				require.Equal(t, uint64(10), i.Uint64())
				return types.NewBlockWithHeader(&types.Header{Number: big.NewInt(10)}), nil
			} else if calledBlocks == 3 {
				require.NotNil(t, i)
				require.Equal(t, uint64(8), i.Uint64())
				return types.NewBlockWithHeader(&types.Header{Number: big.NewInt(8)}), nil
			} else if calledBlocks == 4 {
				require.NotNil(t, i)
				require.Equal(t, uint64(9), i.Uint64())
				return types.NewBlockWithHeader(&types.Header{Number: big.NewInt(9)}), nil
			} else if calledBlocks == 5 {
				require.NotNil(t, i)
				require.Equal(t, uint64(10), i.Uint64())
				return types.NewBlockWithHeader(&types.Header{Number: big.NewInt(10)}), nil
			}
			return nil, errors.New("unexpected call to block by number")
		}).Times(6)
	feed := event.FeedOf[*types.Block]{}
	listener, err := newHttpBlockListener(client, func(block *types.Block) { feed.Send(block) })
	listener = listener.WithMinimal(0)
	require.NoError(t, err)

	wg.Add(1)
	defer wg.Wait()

	blockCh := make(chan *types.Block)
	sub := feed.Subscribe(blockCh)
	defer sub.Unsubscribe()
	go func() {
		require.NoError(t, listener.Start(ctx))
		wg.Done()
	}()

	for i := 8; i <= 10; i++ {
		block := <-blockCh
		require.Equal(t, uint64(i), block.NumberU64())
	}

	for i := 8; i <= 10; i++ {
		block := <-blockCh
		require.Equal(t, uint64(i), block.NumberU64())
	}
}
