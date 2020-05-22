package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	app = cli.NewApp()
)

func init() {
	app.Name = "addrtool"
	app.Usage = "a powerful tool that can generate [mnemonic, seed, address] and supports multi coins"
	app.Copyright = "Copyright 2019-2020 tpkeeper"
	app.Commands = cli.Commands{
		&genMnemonicCommand,
		&genSeed,
		&genAddr,
	}
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
