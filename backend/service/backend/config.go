package backend

import (
	"errors"
	"time"

	"github.com/Exca-DK/pegism/core/clock"
	"github.com/Exca-DK/pegism/service/db"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Config struct {
	IssueSecret     string
	AccessDuration  time.Duration
	RefreshDuration time.Duration
	Database        db.Store
	Clock           *clock.Clock
	EthClient       *ethclient.Client
	FactoryAddress common.Address
	RouterAddress   common.Address
}

// checks whether config is okay
func (cfg Config) Validate() error {
	if cfg.Database == nil {
		return errors.New("db not provided")
	}
	if cfg.AccessDuration >= cfg.RefreshDuration {
		return errors.New("access token has to be lower than refresh token")
	}
	if cfg.Clock == nil {
		return errors.New("clock not provided")
	}
	if cfg.EthClient == nil {
		return errors.New("ethcllient not provided")
	}
	if cfg.FactoryAddress == (common.Address{}) {
		return errors.New("factory address not provided")
	}
	if cfg.RouterAddress == (common.Address{}) {
		return errors.New("router address not provided")
	}
	return nil
}
