# addrtool

![GitHub](https://img.shields.io/github/license/tpkeeper/addrtool?style=plastic)

addrtool is a powerful tool that can generate [mnemonic, seed, address] and supports multi coins

## install
require go 1.11+

```bash
git clone https://github.com/tpkeeper/addrtool.git
cd addrtool
go install ./cmd...
```


## usage

```bash
addrtool.exe

NAME:
   addrtool - a powerful tool that can generate [mnemonic, seed, address] and supports multi coins

USAGE:
   addrtool.exe [global options] command [command options] [arguments...]

COMMANDS:
   genmnemonic, gm  generate mnemonic, protocol support: bip39, decred
   genseed, gs      generate seed from mnemonic, protocol support: bip39, decred
   genaddr, ga      generate address from seed, coin type support: btc, dcr, hc
   help, h          Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)

COPYRIGHT:
   Copyright 2019-2020 tpkeeper

```

**genmnemonic**

```bash
addrtool.exe genmnemonic --size 15
elite marine addict act suffer tuition actor call journey venue kitten width print select dynamic
```
**genseed**

```bash
addrtool.exe genseed "elite marine addict act suffer tuition actor call journey venue kitten width print select dynamic"
dd8d2e80441a1a4440b3edf8febf1b89b78b62e9247924fbc3d11653393c8ebfa124b453f5a25573067ee895e56261b8599d4535649cec580464c77e9d9d7201
```
**genaddr**

```bash
addrtool.exe genaddr --cointype btc --index 1 dd8d2e80441a1a4440b3edf8febf1b89b78b62e9247924fbc3d11653393c8ebfa124b453f5a25573067
ee895e56261b8599d4535649cec580464c77e9d9d7201
17gcwctJfyxnAtprzJadFewHSXs51uV6WW
```