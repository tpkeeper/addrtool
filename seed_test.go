package addrtool

import (
	"encoding/hex"
	"testing"
)

func TestSeedToPubkey(t *testing.T) {
	seed := "2878654d0de7a6fe43250a3895d14bccfbf47e2b9c6ceac6d1ab6341b5ae5033"
	hexByte, _ := hex.DecodeString(seed)
	pubkey, err := SeedToPubKey(hexByte, 44, 171, 0, 0, 0)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("pubkey:", hex.EncodeToString(pubkey))

}

func TestPubkeyFromExtendKey(t *testing.T) {

	pubNode1 := "xpub6FXTeJGXhaHe9jn3K7bK4YdGBjiPcyRhJxWZ2t9MzxCDBerUXVUF1qADBb2eWGJmWUXj3PtMHx9xPoM9idMduN5UwXRfZUVvQUvLhAasF2c"
	pubkeyHex1, err := PubkeyFromExtendKey(pubNode1, 0)
	pubkeyHex2, err := PubkeyFromExtendKey(pubNode1, 0, 0)
	pubkeyHex3, err := PubkeyFromExtendKey(pubNode1, 1, 0, 0)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("pubkeyhex1", pubkeyHex1)
	t.Log("pubkeyhex2", pubkeyHex2)
	t.Log("pubkeyhex3", pubkeyHex3)

}
