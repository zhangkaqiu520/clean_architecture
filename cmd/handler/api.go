package handler

import (
	"github.com/urfave/cli"
	"my-clean-rchitecture/cmd/wire"
)

func Api(*cli.Context) error {
	s := wire.I()

	return s.Run(`:9991`)
}
