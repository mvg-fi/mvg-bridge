package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

const VERSION = "0.0.1"

func main() {
	app := &cli.App{
		Name:                 "mvg-bridge",
		Usage:                "MVG Bridge is an ultimate cross-chain bridge powered by MTG",
		Version:              VERSION,
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:   "boot",
				Usage:  "Boot a MTG node",
				Action: bootCmd,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Value:   "~/.mvg/bridge/config.toml",
						Usage:   "The configuration file path",
					},
					&cli.StringFlag{
						Name:    "dir",
						Aliases: []string{"d"},
						Value:   "~/.mvg/bridge/data",
						Usage:   "The database directory path",
					},
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   5670,
						Usage:   "The RPC server http port",
					},
				},
			},
			{
				Name:   "run",
				Usage:  "Run the application",
				Action: runCmd,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Value:   "~/.mvg/bridge/config.toml",
						Usage:   "The configuration file path",
					},
					&cli.StringFlag{
						Name:    "dir",
						Aliases: []string{"d"},
						Value:   "~/.mvg/bridge/data",
						Usage:   "The database directory path",
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
