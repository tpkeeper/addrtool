package main

import (
	"addrtool"
	"encoding/hex"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	app                = cli.NewApp()
	genMnemonicCommand = cli.Command{
		Name:  "genmnemonic",
		Usage: "generate mnemonic",
		Action: func(c *cli.Context) error {
			protocol := c.String("protocol")
			size := c.Int("size")
			switch size {
			case 12, 15, 18, 21, 24:
			default:
				fmt.Println("size only support: 12, 15, 18, 21, 24")
				return nil
			}

			switch protocol {
			case "bip39":
				bitSize := size*11 - size/3
				mnemonic, err := addrtool.Bip39GenMnemonic(bitSize)
				if err != nil {
					return err
				}
				fmt.Println(mnemonic)
			case "decred":
				fmt.Println("not implement yet")
			default:
				fmt.Println("protocol only support: bip39, decred")
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "protocol",
				Value: "bip39",
				Usage: "protocol for generate mnemonic, support: \"bip39\",\"decred\"",
			},
			&cli.IntFlag{
				Name:  "size",
				Value: 12,
				Usage: "size is the word number of mnemonic, support: 12, 15, 18, 21, 24",
			},
		},
	}

	genSeed = cli.Command{
		Name:  "genseed",
		Usage: "generate seed from mnemonic",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				fmt.Println("This command requires an argument.")
				return nil
			}

			protocol := c.String("protocol")

			switch protocol {
			case "bip39":
				seedBts, err := addrtool.Bip39MnemonicToSeed(c.Args().First(), "")
				if err != nil {
					return err
				}
				fmt.Println(hex.EncodeToString(seedBts))
			case "decred":
				fmt.Println("not implement yet")
			default:
				fmt.Println("protocol only support: bip39, decred")
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "protocol",
				Value: "bip39",
				Usage: "protocol for generate seed, support: \"bip39\",\"decred\"",
			},
		},
	}

	genAddr=cli.Command{
		Name: "genaddr",
		Usage: "generate address from seed",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				fmt.Println("This command requires an argument.")
				return nil
			}

			seed:=c.Args().First()
			coinType := c.String("cointype")
			index:=uint32(c.Uint("index"))

			switch coinType {
			case "btc":
				seedBts,err:=hex.DecodeString(seed)
				if err!=nil{
					return err
				}
				pubKeyBts,err:=addrtool.SeedToPubKey(seedBts,44,0,0,0,index)
				if err != nil {
					return err
				}
				addr:=addrtool.PubkeyToAddress(pubKeyBts,0)
				fmt.Println(addr)
			case "dcr":
				fmt.Println("not implement yet")
			default:
				fmt.Println("coin type now only support: btc, dcr")
			}


			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "cointype",
				Value: "btc",
				Usage: "the coin type for generate address, support: \"btc\",\"dcr\"",
			},
			&cli.UintFlag{
				Name: "index",
				Value: 0,
				Usage: "the index of the derive path",
			},
		},
	}
)

func init() {
	app.Name = "addrtool"
	app.Usage = "a tool for transfer easily between [mnemonic, seed, pubkey, address]"
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
