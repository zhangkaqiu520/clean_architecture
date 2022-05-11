package main

import "my-clean-rchitecture/cmd/wire"

func main() {
	s := wire.I()
	s.Run(`:9991`)
}
