package dumpdata

import (
	"github.com/urfave/cli/v2"
	"public-rpc/internal/config"
	"public-rpc/internal/tools/dumpdata"
)

var Command = &cli.Command{
	Name:  "dumpdata",
	Usage: "Save data from storage to yaml",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "filepath",
			Aliases: []string{"f"},
			Value:   "db-data.yaml",
			Usage:   "Dump data from storage to `FILEPATH`",
		},
	},
	Action: func(cliCtx *cli.Context) error {
		cfg, err := config.LoadConfigFromEnv("")

		if err != nil {
			return cli.Exit(err, 1)
		}

		filepath := cliCtx.String("filepath")

		if filepath == "" {
			return cli.Exit("filepath is required", 1)
		}

		err = dumpdata.DumpData(*cfg, filepath)

		if err != nil {
			return cli.Exit(err, 1)
		}

		return nil
	},
}
