package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"public-rpc/internal/config"
	"public-rpc/internal/service"
)

func main() {
	app := &cli.App{
		Name:                 "public-rpc",
		EnableBashCompletion: true,
		Commands:             []*cli.Command{},
		Action: func(context *cli.Context) error {
			cfg, err := config.LoadConfigFromEnv("")

			if err != nil {
				log.Fatal("error to load config: ", err)
			}

			return service.RunApplication(*cfg)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
