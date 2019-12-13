package addrtool

import (
	"encoding/hex"
	"testing"
)

func TestSeedToPubkey(t *testing.T) {
	tests := []struct {
		name   string
		seed   string
		pubkey string
	}{
		{
			"64byte seed",
			"bffc96a8a2f3317bb5d5b7b107cb257a42a61ca324ce8b5cd2475e05ca579a3e7620f39e58cd6793f755d980fa097b9039e68b09260d25f2eb783449803e63c0",
			"03ef02bfe36119042c30af59b81e77c4912870592d941679253cc2e788505834ec",
		},
	}

	for _, test := range tests {
		hexByte, _ := hex.DecodeString(test.seed)
		pubkey, err := SeedToPubKey(hexByte, 44, 0, 0, 0, 0)
		if err != nil {
			t.Errorf("%v: SeedToPubKey err: %v", test.name, err)
			return
		}
		pubkeyStr := hex.EncodeToString(pubkey)
		if pubkeyStr != test.pubkey {
			t.Errorf("%v: SeedToPubKey failed: want %v got %v", test.name, test.pubkey, pubkeyStr)
			return
		}
	}

}

func TestPubkeyFromExtendKey(t *testing.T) {
	tests := []struct {
		name         string
		extendKey    string
		child0Pubkey string
	}{
		{
			"base58 encode",
			"xpub6EgSKmmPoUPaopdDYyixYBQvz9t9aczJSj2atGuta9gHPEWxvMJLvVf6dHiMW4Wba8PzJG7k4m7tWNTTvc2YnBw5G9PHMU68uvwfUGiRomP",
			"03ef02bfe36119042c30af59b81e77c4912870592d941679253cc2e788505834ec",
		},
	}

	for _, test := range tests {
		pubkey, err := ExtendKeyB58ToPubkey(test.extendKey, 0)
		if err != nil {
			t.Errorf("%v: ExtendKeyB58ToPubkey err: %v", test.name, err)
			return
		}
		pubkeyStr := hex.EncodeToString(pubkey)
		if pubkeyStr != test.child0Pubkey {
			t.Errorf("%v: PubKeyFromExtendKeyB58 failed: want %v got %v", test.name, test.child0Pubkey, pubkeyStr)
			return
		}

	}

}
