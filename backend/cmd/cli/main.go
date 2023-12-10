package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"public-rpc/cmd/cli/dumpdata"
	"public-rpc/cmd/cli/loaddata"
	"public-rpc/cmd/cli/run"
)

func main() {
	app := &cli.App{
		Name:                 "public-rpc",
		Usage:                "Public RPC cli.",
		EnableBashCompletion: true,
		Suggest:              true,
		Commands: []*cli.Command{
			loaddata.Command,
			dumpdata.Command,
			run.Command,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
