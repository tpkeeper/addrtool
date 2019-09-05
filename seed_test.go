package addrtool

import (
	"encoding/hex"
	"testing"
)

func TestSeedToPubkey(t *testing.T) {
	seed := "04ef53d66b17fdfb6538c5d183f0b0569fc1c79d07f044f7670c3038aff411e5abcbe8c457b584d0c1e3504ab94fb311f9097a793c20dfc746a87087ed5dc119"
	hexByte, _ := hex.DecodeString(seed)
	pubkey,err:=SeedToPubKey(hexByte,44,0,0,0,0)
	if err!=nil{
		t.Error(err)
		return
	}

	if hex.EncodeToString(pubkey)!="02a57dc3d8b577f4bdf8dbb53e0083d98298342631fcc24033da0f4b8ebcfdf9f1"{
		t.Error("seedtopubkey failed")
	}

	t.Log("pubkey:",hex.EncodeToString(pubkey))

}

func TestPubkeyFromExtendKey(t *testing.T) {

	pubNode1 := "xpub6FXTeJGXhaHe9jn3K7bK4YdGBjiPcyRhJxWZ2t9MzxCDBerUXVUF1qADBb2eWGJmWUXj3PtMHx9xPoM9idMduN5UwXRfZUVvQUvLhAasF2c"
	pubkeyHex1, err := PubkeyFromExtendKey(pubNode1,0)
	pubkeyHex2, err := PubkeyFromExtendKey(pubNode1, 0, 0)
	pubkeyHex3, err := PubkeyFromExtendKey(pubNode1, 1, 0, 0)

	if err!=nil{
		t.Error(err)
		return
	}

	t.Log("pubkeyhex1", pubkeyHex1)
	t.Log("pubkeyhex2", pubkeyHex2)
	t.Log("pubkeyhex3", pubkeyHex3)

}