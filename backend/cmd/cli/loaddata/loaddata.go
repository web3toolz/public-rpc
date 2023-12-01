package loaddata

import (
	"github.com/urfave/cli/v2"
	"public-rpc/internal/config"
	"public-rpc/internal/tools/loaddata"
)

var Command = &cli.Command{
	Name:  "loaddata",
	Usage: "Load data from yaml to database",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "file",
			Aliases: []string{"f"},
			Value:   "rpc-data.yaml",
			Usage:   "Load data to storage from `FILE`",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		cfg, err := config.LoadConfigFromEnv("")

		if err != nil {
			return cli.Exit(err, 1)
		}

		filepath := cliCtx.String("file")

		if filepath == "" {
			return cli.Exit("file is required", 1)
		}

		err = loaddata.LoadData(*cfg, filepath)

		if err != nil {
			return cli.Exit(err, 1)
		}

		return nil
	},
}
