package app

import (
	"context"
	"fmt"

	"github.com/Exca-DK/pegism/core/registry"
	"github.com/Exca-DK/pegism/service/backend"
	"github.com/Exca-DK/pegism/service/rpc"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Exca-DK/pegism/service/db"

	"github.com/Exca-DK/pegism/core/log"
)

const name = "peggism-service"

type App struct {
	cfg    Config
	ctx    context.Context
	cancel func()
	log    log.Logger
	db     *db.Database

	registry *registry.ServiceRegistry
}

func New() *App {
	ctx, cancel := context.WithCancel(context.Background())
	application := &App{
		log:      log.Root(),
		ctx:      ctx,
		cancel:   cancel,
		registry: registry.NewServiceRegistry(name),
	}
	return application
}

func (app *App) Stop() {
	app.cancel()
}

func (app *App) GetService(srvc interface{}) error {
	return app.registry.FetchService(srvc)
}

func (app *App) Start() {
	app.registry.Start()
}

func (app *App) Setup(cfg Config) error {
	if err := cfg.valid(); err != nil {
		return fmt.Errorf("invalid config provided. %v", err)
	}
	var err error
	err = app.startDB(cfg)
	if err != nil {
		return err
	}

	err = app.registerProfileService(cfg)
	if err != nil {
		return err
	}

	err = app.registerRpcService(cfg)
	if err != nil {
		return err
	}

	return nil
}

func (app *App) startDB(cfg Config) error {
	db, err := db.NewStore(db.Config{
		Username: cfg.Db.Username,
		Password: cfg.Db.Password,
		Db:       cfg.Db.Db,
		Endpoint: cfg.Db.Endpoint,
		Clock:    cfg.Clock,
	})
	if err != nil {
		return err
	}
	app.db = db
	return nil
}

func (app *App) registerProfileService(cfg Config) error {
	var (
		service *backend.Service
		err     error
	)
	var (
		accessDuration  = cfg.AccessDuration
		refreshDuration = cfg.RefreshDuration
		issueSecret     = cfg.IssueSecret
		rpcEndpoint     = cfg.NodeEndpoint
		factoryAddress  = cfg.FactoryAddress
		routerAddress   = cfg.RouterAddress
	)

	ethclient, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		return err
	}

	service, err = backend.New(backend.Config{
		IssueSecret:     issueSecret,
		AccessDuration:  accessDuration,
		RefreshDuration: refreshDuration,
		Database:        app.db,
		Clock:           cfg.Clock,
		EthClient:       ethclient,
		FactoryAddress:  factoryAddress,
		RouterAddress:   routerAddress,
	})
	if err != nil {
		return err
	}
	return app.registry.RegisterService(service)
}

func (app *App) registerRpcService(cfg Config) error {
	var (
		err        error
		profile    *backend.Service
		rpcService *rpc.RpcService
	)
	err = app.GetService(&profile)
	if err != nil {
		return err
	}

	rpcService, err = rpc.NewService(&rpc.Config{
		GrpcConfig: rpc.GrpcConfig{
			Host:        cfg.Rpc.Host,
			Port:        cfg.Rpc.Port,
			GatewayPort: cfg.Rpc.GatewayPort,
			Clock:       cfg.Clock,
		},
		Api: profile.Api(),
	})
	if err != nil {
		return err
	}
	return app.registry.RegisterService(rpcService)
}
