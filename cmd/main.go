package main

import (
	"github.com/urfave/cli"
	"my-clean-rchitecture/cmd/handler"
	"os"
)

var (
	app = &cli.App{}
)

func main() {
	NewApp()
	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}

func NewApp() {
	app.Name = `ddd`
	app.Commands = []cli.Command{
		{
			Name:   "api",
			Usage:  "cover api",
			Action: handler.Api,
			Before: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name:   "migrate",
			Action: handler.Up,
			Before: func(c *cli.Context) error {
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:   "rollback",
					Action: handler.Down,
				},
			},
		},
	}
}
