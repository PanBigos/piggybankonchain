package backend

import (
	"context"
	"errors"
	"time"

	"github.com/Exca-DK/pegism/contracts/piggy/factory"

	"github.com/Exca-DK/pegism/core/blockchain"
	"github.com/Exca-DK/pegism/core/log"
	"github.com/Exca-DK/pegism/service/types"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

var (
	errTransactionNotFound = ErrorWithCode{
		Code: NotFound,
		Err:  errors.New("transaction not found"),
	}
	errValidationFailure = ErrorWithCode{
		Code: Unprocessable,
		Err:  errors.New("failed transaction validation"),
	}
	errInternalError = ErrorWithCode{
		Code: Internal,
		Err:  errors.New("something went wrong"),
	}
)

type transactionController struct {
	controller      *piggyController
	factory         common.Address
	router          common.Address
	factoryContract *factory.FactoryFilterer
	restorer        routerRestorer
	client          blockchain.BlockchainClient
	log             log.Logger
}

func newTransactionController(
	piggyController *piggyController,
	factoryAddress common.Address,
	routerAddress common.Address,
	client blockchain.BlockchainClient,
) (*transactionController, error) {
	factoryContract, err := factory.NewFactoryFilterer(factoryAddress, nil)
	if err != nil {
		return nil, err
	}
	return &transactionController{
		controller:      piggyController,
		factory:         factoryAddress,
		router:          routerAddress,
		log:             log.Root(),
		factoryContract: factoryContract,
		restorer:        newRouterRestorer(routerAddress),
		client:          client,
	}, nil
}

func (t *transactionController) OnNewTx(hash common.Hash) error {
	receipt, err := t.client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		log.Root().Debug("TransactionReceipt failed", "tx", hash.Hex(), "err", err)
		return errTransactionNotFound
	}
	tx, _, err := t.client.TransactionByHash(context.Background(), hash)
	if err != nil {
		log.Root().Debug("TransactionByHash failed", "tx", hash.Hex(), "err", err)
		return errTransactionNotFound
	}

	if tx.To() == nil {
		return errValidationFailure
	}
	to := *tx.To()

	if to == t.factory {
		return t.onNewFactoryTx(tx, receipt)
	} else if to == t.router {
		return t.onNewRouterTx(tx, receipt)
	}
	return errValidationFailure
}

func (t *transactionController) onNewFactoryTx(
	tx *ethTypes.Transaction,
	receipt *ethTypes.Receipt,
) error {
	var err error
	var piggyEvent *factory.FactoryCreatedPiggyBank
	for _, log := range receipt.Logs {
		if log.Address != t.factory {
			continue
		}
		piggyEvent, err = t.factoryContract.ParseCreatedPiggyBank(*log)
		break
	}

	if err != nil {
		t.log.Warn("Failed parsing piggy event", "tx", tx.Hash(), "err", err)
		return errValidationFailure
	} else if piggyEvent == nil {
		return errValidationFailure
	}

	adapted, err := t.adaptPiggyFactoryEvent(piggyEvent)
	if err != nil {
		t.log.Warn("Failed adapting piggy event", "tx", tx.Hash(), "err", err)
		return errValidationFailure
	}

	_, err = t.controller.AddPiggy(adapted)
	if err != nil {
		t.log.Warn("Failed creating piggy", "tx", tx.Hash(), "err", err)
		return errInternalError
	}
	return nil
}

func (t *transactionController) onNewRouterTx(
	tx *ethTypes.Transaction,
	receipt *ethTypes.Receipt,
) error {
	msg := types.Message{Hash: tx.Hash()}
	if err := t.restorer.Restore(&msg, tx.Data(), receipt.Logs); err != nil {
		t.log.Warn("Failed restoring data", "tx", tx.Hash().Hex(), "err", err)
		return err
	}
	_, err := t.controller.AddPiggyMessage(msg)
	if err != nil {
		t.log.Warn("Failed adding piggy message", "tx", tx.Hash().Hex(), "err", err)
		return err
	}
	return nil
}

func (t *transactionController) adaptPiggyFactoryEvent(
	ev *factory.FactoryCreatedPiggyBank,
) (types.Piggy, error) {
	if !ev.UnlockDate.IsInt64() {
		return types.Piggy{}, errors.New("unlock date is not an valid date")
	}
	if !ev.CreatedAt.IsInt64() {
		return types.Piggy{}, errors.New("created at date is not an valid date")
	}
	unlockDate := time.Unix(ev.UnlockDate.Int64(), 0)
	createdAt := time.Unix(ev.CreatedAt.Int64(), 0)

	return types.Piggy{
		Address:        ev.PiggyBank,
		FromAddress:    ev.Creator,
		ProfileAddress: ev.Owner,
		CreatedAt:      createdAt,
		UnlocksAt:      unlockDate,
	}, nil
}
