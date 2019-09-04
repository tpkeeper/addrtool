package addrtool

import (
	"github.com/tyler-smith/go-bip32"
	"encoding/hex"
	"testing"
)

func TestSeedToPubkey(t *testing.T) {
	seed := "04ef53d66b17fdfb6538c5d183f0b0569fc1c79d07f044f7670c3038aff411e5abcbe8c457b584d0c1e3504ab94fb311f9097a793c20dfc746a87087ed5dc119"
	hexByte, _ := hex.DecodeString(seed)
	//m
	masterExtKey, _ := bip32.NewMasterKey(hexByte)
	//m/purpose'
	purposeExtKey,_:=masterExtKey.NewChildKey(bip32.FirstHardenedChild+44)
	//m/purpose'/cointype'
	coinTypeExtKey,_:=purposeExtKey.NewChildKey(bip32.FirstHardenedChild+0)
	//m/purpose'/cointype'/account'
	accountExtKey,_:=coinTypeExtKey.NewChildKey(bip32.FirstHardenedChild+0)
	//m/purpose'/cointype'/account'/change
	changeExtKey,_:=accountExtKey.NewChildKey(0)
	//m/purpose'/cointype'/account'/change/addrIndex
	addrIndex0ExtKey,_:=changeExtKey.NewChildKey(0)
	//pubkey
	t.Log(hex.EncodeToString(addrIndex0ExtKey.PublicKey().Key))
}



func TestSeedToAddr(t *testing.T) {
	debug = true
	//seed:="c07f1b752c7af8abc94740bf9467cde0057165895467fa83062fa78caef60aca"
	//seed:="6b81b8ca17c98397ae843899854e89a3958e492db740469e6759bdc3845d7289"
	seed := "a5957d6e848f94dd1da8806e432aac9ebe836eb36d689eab828593359e93a6ab"
	//seed:="00a84c51041d49acca66e6160c1fa999"
	hexByte, _ := hex.DecodeString(seed)

	addr, err := SeedToAddr(hexByte, hcMainNetParams, 45, 0, 0, 0)
	if err != nil {
		t.Error(err)
	}
	t.Log(addr)

	pubNode1 := "dpubZB3aSKzsQCrtQNiEZdprCTskujNACqXJXAcLRVUCTU4Qwpi8EqgW6kVKygjrTt27wMkuK37v1UsQYzKE1xhrbgqCZdMGw6gBYZVTkSogTDC"
	pubNode2 := "dpubZCRagbXPbBG17kmPaxVjBSody6zu4ChE2ywd7xGfmh2WLbCZJZxbqCH3HL16ytiz91UfDmRbTVBzcgxFAxscqHyh9Hi3AUStcXcA7HV2sBX"
	pubkeyHex1, err := PubkeyFromNode(pubNode1, hcMainNetParams, true, 0, 0, 0)
	pubkeyHex2, err := PubkeyFromNode(pubNode2, hcMainNetParams, true, 1, 0, 0)
	t.Log("pubkeyhex1", pubkeyHex1)
	t.Log("pubkeyhex2", pubkeyHex2)

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
