package backend

import (
	"time"

	"github.com/Exca-DK/pegism/core/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

type blockchainListener struct {
	notifier       blockchainNotifier
	api            api
	factoryAddress common.Address
	routerAddress  common.Address

	done chan struct{}

	log log.Logger
}

type blockchainNotifier interface {
	SubscribeNewBlock(ch chan<- *types.Block) event.Subscription
}

type api interface {
	NotifyNewTransaction(hash common.Hash) error
}

func newListener(
	api api,
	notifier blockchainNotifier,
	factoryAddress common.Address,
	routerAddress common.Address,
) *blockchainListener {
	return &blockchainListener{
		notifier:       notifier,
		api:            api,
		factoryAddress: factoryAddress,
		routerAddress:  routerAddress,
		done:           make(chan struct{}),
		log:            log.Root(),
	}
}

func (listener *blockchainListener) Start() {
	ch := make(chan *types.Block)
	sub := listener.notifier.SubscribeNewBlock(ch)
	defer sub.Unsubscribe()
	for {
		select {
		case <-listener.done:
			return
		case block := <-ch:
			for _, tx := range block.Transactions() {
				if tx.To() == nil {
					continue
				}
				if *tx.To() == listener.factoryAddress || *tx.To() == listener.routerAddress {
					ts := time.Now()
					err := listener.api.NotifyNewTransaction(tx.Hash())
					if err == nil {
						listener.log.Debug(
							"Added transaction from chain",
							"tx", tx.Hash().Hex(),
							"elapsed", time.Since(ts).String(),
						)
						continue
					}
					_, ok := err.(ErrorWithCode)
					if !ok {
						listener.log.Warn(
							"Failed notifying new transaction",
							"tx", tx.Hash().Hex(),
							"err", err,
						)
						continue
					}
					listener.log.Trace(
						"rejected transaction",
						"tx", tx.Hash().Hex(),
						"err", err,
					)
				}
			}
		}
	}
}

func (listener *blockchainListener) Stop() {
	close(listener.done)
}
