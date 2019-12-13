package addrtool

import (
	"encoding/hex"
	"testing"
)

func TestPubkeyToAddress(t *testing.T) {

	tests := []struct {
		name   string
		netId  byte
		pubkey string
		result string
	}{
		{
			name: "pubkeytoaddress",
			netId:  0,
			pubkey: "02be17e801d577970bdbe875ab3c7705e049095e377e3273311b7bc6324f4be349",
			result: "1MCbqnDheNHKtS2eCwgLS9JyTS25HkVZoH"},
	}

	for _, test := range tests {

		hexByte, err := hex.DecodeString(test.pubkey)
		if err != nil {
			t.Errorf("%v: decoding pubkey hex err: %v", test.name, err)
			return
		}
		addr := PubkeyToAddress(hexByte, test.netId)
		if addr != test.result {
			t.Errorf("%v: pubkey to address does not match expected value: want %v got %v", test.name, test.result, addr)
			return
		}

	}
}


func TestDcrPubkeyToAddress(t *testing.T) {
	tests := []struct {
		name   string
		netId  [2]byte
		pubkey string
		result string
	}{
		{
			name: "dcrpubkeytoaddress",
			netId:  hcTestNetParams.PubKeyHashAddrID,
			pubkey: "02b1ad2bc0a9d189c4c644ac2668d62b2b6147cea7858b894987b4689489186d6b",
			result: "TsSmoC9HdBhDhq4ut4TqJY7SBjPqJFAPkGK"},
	}

	for _, test := range tests {

		hexByte, err := hex.DecodeString(test.pubkey)
		if err != nil {
			t.Errorf("%v: decoding pubkey hex err: %v", test.name, err)
			return
		}
		addr := DcrPubkeyToAddress(hexByte, test.netId)
		if addr != test.result {
			t.Errorf("%v: pubkey to address does not match expected value: want %v got %v", test.name, test.result, addr)
			return
		}

	}

}

var hcMainNetParams = &DcrNetWorkParams{
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

var hcTestNetParams = &DcrNetWorkParams{
	// BIP32 hierarchical deterministic extended key magics
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x97}, // starts with tprv
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xd1}, // starts with tpub

	PubKeyAddrID:     [2]byte{0x28, 0xf7}, // starts with Tk
	PubKeyHashAddrID: [2]byte{0x0f, 0x21}, // starts with Ts
	PKHEdwardsAddrID: [2]byte{0x0f, 0x01}, // starts with Te
	PKHSchnorrAddrID: [2]byte{0x0e, 0xe3}, // starts with Ts
	ScriptHashAddrID: [2]byte{0x0e, 0xfc}, // starts with Tm
	PrivateKeyID:     [2]byte{0x23, 0x0e}, // starts with Pt

	HDCoinType: uint32(171),
}
