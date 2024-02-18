package app

import (
	"errors"
	"time"

	"github.com/Exca-DK/pegism/core/clock"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Clock *clock.Clock
	Db    struct {
		Username string
		Password string
		Db       string
		Endpoint string
	}
	IssueSecret     string
	AccessDuration  time.Duration
	RefreshDuration time.Duration
	Rpc             struct {
		Host           string
		Port           int
		GatewayPort    int
		MaxMsgSizeInMb uint64
	}
	NodeEndpoint   string
	FactoryAddress common.Address
	RouterAddress  common.Address
}

func (cfg *Config) valid() error {
	if cfg.Clock == nil {
		return errors.New("clock not set")
	}
	if len(cfg.NodeEndpoint) == 0 {
		return errors.New("node endpoint not set")
	}
	if cfg.FactoryAddress == (common.Address{}) {
		return errors.New("factory address not set")
	}
	if cfg.RouterAddress == (common.Address{}) {
		return errors.New("router address not set")
	}
	return nil
}
