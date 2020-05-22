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
		Usage: "generate mnemonic, protocol support: bip39, decred",
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
		Usage: "generate seed from mnemonic, protocol support: bip39, decred",
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

	genAddr = cli.Command{
		Name:  "genaddr",
		Usage: "generate address from seed, coin type support: btc, dcr, hc",
		Action: func(c *cli.Context) error {
			if c.NArg() != 1 {
				fmt.Println("This command requires an argument.")
				return nil
			}

			seed := c.Args().First()
			coinType := c.String("cointype")
			index := uint32(c.Uint("index"))

			seedBts, err := hex.DecodeString(seed)
			if err != nil {
				return err
			}
			pubKeyBts, err := addrtool.SeedToPubKey(seedBts, 44, 0, 0, 0, index)
			if err != nil {
				return err
			}

			switch coinType {
			case "btc":
				addr := addrtool.PubkeyToAddress(pubKeyBts, 0)
				fmt.Println(addr)
			case "dcr":
				addr := addrtool.DcrPubkeyToAddress(pubKeyBts, dcrNetParams.PubKeyHashAddrID)
				fmt.Println(addr)
			case "hc":
				addr := addrtool.DcrPubkeyToAddress(pubKeyBts, hcNetParams.PubKeyHashAddrID)
				fmt.Println(addr)
			default:
				fmt.Println("coin type now only support: btc, dcr, hc")
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "cointype",
				Value: "btc",
				Usage: "the coin type for generate address, support: \"btc\",\"dcr\",\"hc\"",
			},
			&cli.UintFlag{
				Name:  "index",
				Value: 0,
				Usage: "the index of the derive path",
			},
		},
	}
	dcrNetParams = addrtool.DcrNetWorkParams{
		HDCoinType: 20,
		// Address encoding magics
		PubKeyAddrID:     [2]byte{0x13, 0x86}, // starts with Dk
		PubKeyHashAddrID: [2]byte{0x07, 0x3f}, // starts with Ds
		PKHEdwardsAddrID: [2]byte{0x07, 0x1f}, // starts with De
		PKHSchnorrAddrID: [2]byte{0x07, 0x01}, // starts with DS
		ScriptHashAddrID: [2]byte{0x07, 0x1a}, // starts with Dc
		PrivateKeyID:     [2]byte{0x22, 0xde}, // starts with Pm

		// BIP32 hierarchical deterministic extended key magics
		HDPrivateKeyID: [4]byte{0x02, 0xfd, 0xa4, 0xe8}, // starts with dprv
		HDPublicKeyID:  [4]byte{0x02, 0xfd, 0xa9, 0x26}, // starts with dpub
	}
	hcNetParams = addrtool.DcrNetWorkParams{
		HDCoinType:       171,
		PubKeyAddrID:     [2]byte{0x19, 0xa4}, // starts with Hk
		PubKeyHashAddrID: [2]byte{0x09, 0x7f}, // starts with Hs
		PKHEdwardsAddrID: [2]byte{0x09, 0x60}, // starts with He
		PKHSchnorrAddrID: [2]byte{0x09, 0x41}, // starts with HS
		ScriptHashAddrID: [2]byte{0x09, 0x5a}, // starts with Hc
		PrivateKeyID:     [2]byte{0x19, 0xab}, // starts with Hm

		// BIP32 hierarchical deterministic extended key magics
		HDPrivateKeyID: [4]byte{0x02, 0xfd, 0xa4, 0xe8}, // starts with dprv
		HDPublicKeyID:  [4]byte{0x02, 0xfd, 0xa9, 0x26}, // starts with dpub
	}
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
