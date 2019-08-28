package addrtool

import (
	"encoding/hex"
	"testing"
)

func TestSeedToAddr(t *testing.T) {
	seed:="c07f1b752c7af8abc94740bf9467cde0057165895467fa83062fa78caef60aca"
	//seed:="00a84c51041d49acca66e6160c1fa999"
	hexByte, _ := hex.DecodeString(seed)
	t.Log(SeedToAddr(hexByte,hcTestNetParams))

}
var hcTestNetParams=&NetWorkParams{
	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x97}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xd1}, // starts with tpub

	PubKeyAddrID:         [2]byte{0x28, 0xf7}, // starts with Tk
	PubKeyHashAddrID:     [2]byte{0x0e, 0xc0}, // starts with TC
	PKHEdwardsAddrID:     [2]byte{0x0f, 0x01}, // starts with Te
	PKHSchnorrAddrID:     [2]byte{0x0f, 0x20}, // starts with Ts
	ScriptHashAddrID:     [2]byte{0x0f, 0x12}, // starts with Tm
	PrivateKeyID:         [2]byte{0x23, 0x0e}, // starts with Pt

	HDCoinType:uint32(1),
}
