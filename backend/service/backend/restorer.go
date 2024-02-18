package backend

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/Exca-DK/pegism/contracts/piggy/router"
	"github.com/Exca-DK/pegism/service/types"

	cryptotypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
)

type calldata []byte

func (c calldata) Valid() bool { return len(c) > 4 }

func (c calldata) Data() []byte { return c[4:] }

var (
	routerFilterer, _ = router.NewRouterFilterer(common.Address{}, nil)
)

var (
	tokenTransferSignature = common.HexToHash(
		"0xbba8a6f1ace6d0ccb2089d879d1bf044d9153802c1d010c514711798d413828c",
	)
	ethTransferSignature = common.HexToHash(
		"0x7fa9aafeb8bb803d77de5d84bc2f2edbd842ca91b20cd5020aa21dfe26ab0be9",
	)
)

var (
	errFeeOverflows    = errors.New("fee overflows")
	errAmountOverflows = errors.New("amount overflows")
)

type routerRestorer struct {
	address common.Address
}

func newRouterRestorer(address common.Address) routerRestorer {
	return routerRestorer{address: address}
}

func (restorer routerRestorer) Restore(
	msg *types.Message,
	txData calldata,
	logs []*cryptotypes.Log,
) error {
	if !txData.Valid() {
		return errors.New("illegal calldata")
	}
	var (
		amount         *big.Int
		fee            *big.Int
		to             common.Address
		token          common.Address
		calldataParser func(calldata) (string, string, error)
	)

OUTER:
	for _, log := range logs {
		if log.Address != restorer.address {
			continue
		}
		for _, topic := range log.Topics {
			switch topic {
			case tokenTransferSignature:
				calldataParser = tokenTransferToContent
				parsed, err := routerFilterer.ParseTokenTransfer(*log)
				if err != nil {
					return err
				}
				fee = parsed.Fee
				amount = parsed.Amount
				to = parsed.To
				token = parsed.Token
				break OUTER
			case ethTransferSignature:
				calldataParser = ethTransferToContent
				parsed, err := routerFilterer.ParseTransfer(*log)
				if err != nil {
					return err
				}
				fee = parsed.Fee
				amount = parsed.Amount
				to = parsed.To
				token = common.Address{}
				break OUTER
			}
		}
	}

	if calldataParser == nil {
		return errors.New("message not detected in transaction")
	}

	nick, content, err := calldataParser(txData)
	if err != nil {
		return err
	}

	msg.Address = to
	msg.Token = token
	msg.Amount = amount
	msg.Fee = fee
	msg.Nick = nick
	msg.Content = content
	return nil
}

func tokenTransferToContent(encoded calldata) (string, string, error) {
	data := encoded.Data()
	if len(data) < 3*32 {
		return "", "", errors.New("donation message has not been found")
	}
	// skip over eth.
	encodedPart := data[3*32-32-4:]
	return ethTransferToContent(encodedPart)
}

func ethTransferToContent(encoded calldata) (string, string, error) {
	data := encoded.Data()
	if len(data) < 32 {
		return "", "", errors.New("donation message has not been found")
	}
	type template struct {
		Nick    string `json:"nick"`
		Message string `json:"msg"`
	}
	t := template{}
	err := json.Unmarshal(data[32:], &t)
	if err != nil {
		return "", "", err
	}
	return t.Nick, t.Message, nil
}
