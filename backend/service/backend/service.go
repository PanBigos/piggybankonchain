package backend

import (
	"fmt"
	"sync"

	"github.com/Exca-DK/pegism/core/blockchain"
	"github.com/Exca-DK/pegism/core/clock"
)

type Service struct {
	wg sync.WaitGroup

	clock *clock.Clock

	issuer      *tokenIssuer
	session     *sessionController
	profile     *profileController
	listener    *blockchainListener
	transaction *transactionController
	piggy       *piggyController
}

func New(cfg Config) (*Service, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid profile config provided. %w", err)
	}

	service := &Service{clock: cfg.Clock}
	blockchain, err := blockchain.NewBlockchain(cfg.EthClient)
	if err != nil {
		return nil, fmt.Errorf("failed bootstraping from remote blockchain. %v", err)
	}

	issuer, err := newTokenController(
		cfg.IssueSecret,
		cfg.AccessDuration,
		cfg.RefreshDuration,
		cfg.Clock,
	)
	if err != nil {
		return nil, err
	}
	service.issuer = issuer
	service.session = newSessionController(cfg.Database, issuer)
	service.profile = newProfile(cfg.Database)
	service.piggy = newPiggyController(cfg.Database)
	service.listener = newListener(
		service.Api(),
		blockchain,
		cfg.FactoryAddress,
		cfg.RouterAddress,
	)
	service.transaction, err = newTransactionController(
		service.piggy,
		cfg.FactoryAddress,
		cfg.RouterAddress,
		cfg.EthClient,
	)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (b *Service) Api() *BackendApi {
	return &BackendApi{backend: b}
}

func (b *Service) Start() error {
	return nil
}
func (b *Service) Stop() error {
	b.wg.Wait()
	return nil
}
func (b *Service) Status() error { return nil }
