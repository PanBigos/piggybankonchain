package blockchain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync/atomic"
	"time"

	"github.com/Exca-DK/pegism/core/backoff"
	"github.com/Exca-DK/pegism/core/log"
	"github.com/ethereum/go-ethereum/core/types"
)

type httpBlockListener struct {
	client blockClient

	log                 log.Logger
	backoff             backoff.Backoff
	minimalBackoff      time.Duration
	currentBlockHeight  atomic.Uint64
	heighestBlockHeight atomic.Uint64

	blockNotify func(*types.Block)
}

func newHttpBlockListener(
	client blockClient,
	notify func(block *types.Block),
) (*httpBlockListener, error) {
	currentHeader, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	listener := &httpBlockListener{
		client:         client,
		log:            log.Root(),
		minimalBackoff: 2 * time.Second,
		backoff:        backoff.NoopBackoff,
		blockNotify:    notify,
	}
	listener.log.Info(fmt.Sprintf("Starting head is at %v", currentHeader.Number.Uint64()))
	listener.currentBlockHeight.Store(currentHeader.Number.Uint64())
	return listener, nil
}

func (listener *httpBlockListener) WithBackoff(backoff backoff.Backoff) *httpBlockListener {
	listener.backoff = backoff
	return listener
}

func (listener *httpBlockListener) WithMinimal(t time.Duration) *httpBlockListener {
	listener.minimalBackoff = t
	return listener
}

func (listener *httpBlockListener) Start(ctx context.Context) error {
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
	return verifyError(<-ch, ctx)
}

func (listener *httpBlockListener) catchupBlocks(ctx context.Context) error {
	base := listener.getBlockTimeDiff()
	wait := base
	timer := time.NewTimer(wait)
	defer timer.Stop()

	firstIteration := true

	diffTicker := time.NewTicker(1 * time.Hour)
	defer diffTicker.Stop()
	for {
		select {
		case <-diffTicker.C:
			base = listener.getBlockTimeDiff()
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			current := listener.currentBlockHeight.Load()
			highest, err := listener.fetchLatest(ctx)
			if highest == current {
				timer.Reset(wait)
				continue
			}
			listener.log.Trace("Fetched latest block head", "number", highest)
			// reorg
			if highest < current {
				temp := current
				current = highest
				highest = temp
				listener.log.Info("Reorg detected",
					"block.current", current,
					"block.latest.known", highest)
				firstIteration = true
			}
			if err != nil {
				listener.log.Warn(
					"Failed fetching latest block",
					"block.current", current,
					"block.latest.known", listener.heighestBlockHeight.Load(),
					"err", err,
				)
				wait = listener.backoff(wait)
				timer.Reset(wait)
				continue
			}
			if !firstIteration {
				current += 1
			}
		OUTER:
			for i := current; i <= highest; i++ {
				if err := listener.catchupBlock(ctx, i); err != nil {
					listener.log.Warn(
						"Failed catching up block",
						"block.current", listener.currentBlockHeight.Load(),
						"block.latest.known", listener.heighestBlockHeight.Load(),
						"block.failed", i,
						"err", err,
					)
					break OUTER
				}
			}
			wait = base
			firstIteration = false
			timer.Reset(wait)
		}
	}
}

func (listener *httpBlockListener) catchupBlock(ctx context.Context, number uint64) error {
	listener.log.Trace("Calling block by numer", "number", number)
	block, err := listener.client.BlockByNumber(ctx, big.NewInt(int64(number)))
	if err != nil {
		return err
	}
	listener.notifyNewBlock(ctx, block)
	return nil
}

func (listener *httpBlockListener) notifyNewBlock(ctx context.Context, block *types.Block) {
	listener.log.Trace("Sending new block notification", "number", block.NumberU64())
	listener.currentBlockHeight.Store(block.NumberU64())
	listener.blockNotify(block)
}

func (listener *httpBlockListener) fetchLatest(ctx context.Context) (uint64, error) {
	height, err := listener.client.HeaderByNumber(ctx, nil)
	if err != nil {
		return 0, err
	}
	number := height.Number.Uint64()
	listener.heighestBlockHeight.Store(height.Number.Uint64())
	return number, nil
}

func (listener *httpBlockListener) getBlockTimeDiff() time.Duration {
	start := listener.currentBlockHeight.Load()
	if start == 0 {
		start = 1
	}
	startBlock, err := listener.client.HeaderByNumber(
		context.Background(),
		big.NewInt(int64(start)),
	)
	if err != nil {
		listener.log.Warn(
			"Failed getting init header",
			"number", start,
			"err", err,
		)
		return listener.getDefaultBlockDiff()
	}
	prevBlock, err := listener.client.HeaderByNumber(
		context.Background(),
		big.NewInt(int64(start-1)),
	)
	if err != nil {
		listener.log.Warn(
			"Failed getting init header",
			"number", start-1,
			"err", err,
		)
		return listener.getDefaultBlockDiff()
	}

	tsFirst := time.Unix(int64(startBlock.Time), 0)
	tsPrev := time.Unix(int64(prevBlock.Time), 0)
	diff := tsFirst.Sub(tsPrev)
	if diff < 0 {
		diff *= -1
	}

	if diff < listener.minimalBackoff {
		diff = listener.minimalBackoff
	}
	listener.log.Debug("Fetched latest block diff", "diff", diff.String())
	return diff
}

func (listener *httpBlockListener) getDefaultBlockDiff() time.Duration {
	return 30 * time.Second
}
