package addrtool

import (
	"encoding/hex"
	"testing"
)

func TestPubkeyToAddress(t *testing.T) {
	hexByte,_:=hex.DecodeString("02a57dc3d8b577f4bdf8dbb53e0083d98298342631fcc24033da0f4b8ebcfdf9f1")

	t.Log(PubkeyToAddress(hexByte,0))
}

var hcMainNetParams = &NetWorkParams{
	PubKeyAddrID:     [2]byte{0x19, 0xa4}, // starts with Hk
	PubKeyHashAddrID: [2]byte{0x09, 0x7f}, // starts with Hs
	PKHEdwardsAddrID: [2]byte{0x09, 0x60}, // starts with He
	PKHSchnorrAddrID: [2]byte{0x09, 0x41}, // starts with HS
	ScriptHashAddrID: [2]byte{0x09, 0x5a}, // starts with Hc
	PrivateKeyID:     [2]byte{0x19, 0xab}, // starts with Hm

	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x02, 0xfd, 0xa4, 0xe8}, // starts with dprv
	HDPublicKeyID:  [4]byte{0x02, 0xfd, 0xa9, 0x26}, // starts with dpub
	HDCoinType:     uint32(171),
}

var hcTestNetParams = &NetWorkParams{
	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x97}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xd1}, // starts with tpub

	PubKeyAddrID:     [2]byte{0x28, 0xf7}, // starts with Tk
	PubKeyHashAddrID: [2]byte{0x0e, 0xc0}, // starts with TC
	PKHEdwardsAddrID: [2]byte{0x0f, 0x01}, // starts with Te
	PKHSchnorrAddrID: [2]byte{0x0f, 0x20}, // starts with Ts
	ScriptHashAddrID: [2]byte{0x0f, 0x12}, // starts with Tm
	PrivateKeyID:     [2]byte{0x23, 0x0e}, // starts with Pt

	HDCoinType: uint32(171),
}
