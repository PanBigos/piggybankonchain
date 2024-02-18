package launcher

import (
	"fmt"
	"time"

	"github.com/Exca-DK/pegism/core/log"
	"github.com/urfave/cli/v2"
)

var (
	generalFlags = []cli.Flag{
		loggingLevel,
		nodeEndpointFlag,
		contractFactoryAddressFlag,
		contractRouterAddressFlag,
		serverPort,
		serverHost,
		serverGatewayPort,
	}
	loggingLevel = &cli.StringFlag{
		Name:  "logging",
		Usage: fmt.Sprintf("Options are: %v", log.Levels()),
	}
	nodeEndpointFlag = &cli.StringFlag{
		Name:     "node.endpoint",
		EnvVars:  []string{"node.endpoint"},
		Required: true,
	}
	contractFactoryAddressFlag = &cli.StringFlag{
		Name:     "contract.factory.address",
		EnvVars:  []string{"contract.factory.address"},
		Required: true,
	}
	contractRouterAddressFlag = &cli.StringFlag{
		Name:     "contract.router.address",
		EnvVars:  []string{"contract.router.address"},
		Required: true,
	}
	serverPort        = &cli.StringFlag{Name: "rpc.server.port", Required: true}
	serverHost        = &cli.StringFlag{Name: "rpc.server.host", Required: true}
	serverGatewayPort = &cli.StringFlag{Name: "rpc.gateway.port"}

	// db flags
	databaseFlags   = []cli.Flag{databseEndpoint, databseUsername, databsePassword, databseName}
	databseEndpoint = &cli.StringFlag{
		Name:     "db.endpoint",
		Required: true,
		EnvVars:  []string{"db.endpoint"},
	}
	databseUsername = &cli.StringFlag{
		Name:     "db.username",
		Required: true,
		EnvVars:  []string{"db.username"},
	}
	databsePassword = &cli.StringFlag{
		Name:     "db.password",
		Required: true,
		EnvVars:  []string{"db.password"},
	}
	databseName = &cli.StringFlag{Name: "db.name", Required: true, EnvVars: []string{"db.name"}}

	//authService flags
	profileFlags    = []cli.Flag{issueSecretFlag, issueAccessDurationFlag, issueRefreshDurationFlag}
	issueSecretFlag = &cli.StringFlag{
		Name:     "auth.issuer.secret",
		Required: true,
		EnvVars:  []string{"auth.issuer.secret"},
	}
	issueAccessDurationFlag = &cli.DurationFlag{
		Name:    "auth.issuer.access.duration",
		Value:   10 * time.Minute,
		EnvVars: []string{"auth.issuer.access.duration"},
	}
	issueRefreshDurationFlag = &cli.DurationFlag{
		Name:    "auth.issuer.refresh.duration",
		Value:   24 * time.Hour * 31,
		EnvVars: []string{"auth.issuer.refresh.duration"},
	}
)

func GetCliFlags() []cli.Flag {
	flags := make([]cli.Flag, 0)
	flags = append(flags, generalFlags...)
	flags = append(flags, databaseFlags...)
	flags = append(flags, profileFlags...)
	return flags
}
