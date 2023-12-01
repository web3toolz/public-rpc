package run

import (
	"github.com/urfave/cli/v2"
	"public-rpc/internal/config"
	"public-rpc/internal/service"
)

var Command = &cli.Command{
	Name:  "run",
	Usage: "Run main server and worker",
	Flags: []cli.Flag{},
	Action: func(cliCtx *cli.Context) error {
		cfg, err := config.LoadConfigFromEnv("")

		if err != nil {
			return cli.Exit(err, 1)
		}

		return service.RunApplication(*cfg)
	},
}
