package launcher

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/Exca-DK/pegism/service/app"

	"github.com/Exca-DK/pegism/core/log"
	"github.com/urfave/cli/v2"
)

var (
	cli_app = &cli.App{
		Name:        "piggy",
		Usage:       "This is a microservice implementing piggy service",
		Version:     app.Version(),
		Writer:      os.Stdout,
		HideVersion: false,
	}
)

func init() {
	// Initialize the CLI app
	cli_app.Action = func(ctx *cli.Context) error {
		if err := profile(ctx); err != nil {
			return cli.Exit(err.Error(), 1)
		}
		return nil
	}
	cli_app.Flags = append(cli_app.Flags, GetCliFlags()...)
	cli_app.Commands = make([]*cli.Command, 0)
	cli_app.Before = func(ctx *cli.Context) error { return nil }
	cli_app.After = func(ctx *cli.Context) error { return nil }
}

// main entry
func profile(ctx *cli.Context) error {
	if args := ctx.Args().Slice(); len(args) > 0 {
		return fmt.Errorf("invalid command: %q", args[0])
	}
	if err := prepareLogging(ctx); err != nil {
		return fmt.Errorf("logging preparation failure. %v", err)
	}
	cfg, err := parseConfig(ctx)
	if err != nil {
		return err
	}
	app := app.New()
	if err := app.Setup(cfg); err != nil {
		return err
	}

	app.Start()
	return nil
}

func Launch(args []string) error {
	defer func() {
		if x := recover(); x != nil {
			log.Root().Error(fmt.Sprintf("Runtime panic: %v\n%v", x, string(debug.Stack())))
			panic(x)
		}
	}()
	return cli_app.Run(args)
}
