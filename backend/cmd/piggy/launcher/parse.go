package launcher

import (
	"fmt"

	"github.com/Exca-DK/pegism/core/clock"
	"github.com/Exca-DK/pegism/core/log"
	"github.com/Exca-DK/pegism/service/app"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli/v2"
)

func prepareLogging(ctx *cli.Context) error {
	cli_lvl := ctx.String(loggingLevel.Name)
	var lvl log.Lvl
	if err := lvl.FromString(cli_lvl); err != nil {
		return err
	}

	log.ChangeLvl(lvl)
	return nil
}

func parseConfig(ctx *cli.Context) (app.Config, error) {
	cfg := app.Config{Clock: clock.NewClock()}
	parsers := []cfgParser{parseRpc, parseDb, parseProfile}
	for _, parser := range parsers {
		if err := parser(ctx, &cfg); err != nil {
			return app.Config{}, fmt.Errorf("config parsing failure. %v", err)
		}
	}
	return cfg, nil
}

type cfgParser func(ctx *cli.Context, cfg *app.Config) error

func parseRpc(ctx *cli.Context, cfg *app.Config) error {
	var (
		port        int
		gatewayPort int
		host        string
	)
	if !ctx.IsSet(serverPort.Name) {
		log.Root().Info("grpc.port not provided, fallback to 5050")
		port = 5050
	} else {
		port = ctx.Int(serverPort.Name)
	}
	if !ctx.IsSet(serverGatewayPort.Name) {
		log.Root().Info("gateway port not provided, fallback to grpc.port + 1")
		gatewayPort = port + 1
	} else {
		gatewayPort = ctx.Int(serverGatewayPort.Name)
	}

	if !ctx.IsSet(serverHost.Name) {
		log.Root().Info("grpc host not provided, fallback to localhost")
		host = "127.0.0.1"
	} else {
		host = ctx.String(serverHost.Name)
	}

	cfg.Rpc.Port = port
	cfg.Rpc.GatewayPort = gatewayPort
	cfg.Rpc.Host = host
	return nil
}

func parseDb(ctx *cli.Context, cfg *app.Config) error {
	cfg.Db.Db = ctx.String(databseName.Name)
	cfg.Db.Endpoint = ctx.String(databseEndpoint.Name)
	cfg.Db.Username = ctx.String(databseUsername.Name)
	cfg.Db.Password = ctx.String(databsePassword.Name)
	return nil
}

func parseProfile(ctx *cli.Context, cfg *app.Config) error {
	cfg.IssueSecret = ctx.String(issueSecretFlag.Name)
	cfg.AccessDuration = ctx.Duration(issueAccessDurationFlag.Name)
	cfg.RefreshDuration = ctx.Duration(issueRefreshDurationFlag.Name)
	cfg.NodeEndpoint = ctx.String(nodeEndpointFlag.Name)
	cfg.FactoryAddress = common.HexToAddress(ctx.String(contractFactoryAddressFlag.Name))
	cfg.RouterAddress = common.HexToAddress(ctx.String(contractRouterAddressFlag.Name))
	return nil
}
