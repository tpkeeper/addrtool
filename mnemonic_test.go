package addrtool

import (
	"bytes"
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
	dcrbase58 "github.com/decred/base58"
	"strings"
	"testing"
)

func TestGenMnemonicBip39(t *testing.T) {
	tests := []struct {
		name     string
		randSize int
		wordsNum int
	}{
		{
			"rand 128bit ",
			128,
			12,
		},
		{
			"rand 160bit",
			160,
			15,
		},
		{
			"rand 192bit",
			192,
			18,
		},
		{
			"rand 224bit",
			224,
			21,
		},
		{
			"rand 256bit",
			256,
			24,
		},
	}

	for _, test := range tests {
		mnemonic, err := Bip39GenMnemonic(test.randSize)
		if err != nil {
			t.Errorf("%v: GenMnemonicBip39 err: %v", test.name, err)
			return
		}
		mLen := len(strings.Split(mnemonic, " "))
		if mLen != test.wordsNum {
			t.Errorf("%v: GenMnemonicBip39 fail : want %v got %v", test.name, test.wordsNum, mLen)
			return
		}
	}
}

func TestMnemonicToSeedBip39(t *testing.T) {
	tests := []struct {
		name     string
		mnemonic string
		seed     []byte
	}{
		{
			name:     "12 words",
			mnemonic: "close same tongue random ice cave aim input whale salute squirrel vivid",
			seed: []byte{0x12, 0x6b, 0x7f, 0x86, 0x53, 0xce, 0x2b, 0x1f, 0x05, 0xdd, 0x78, 0xd3, 0x3c, 0x57, 0x73, 0x7d, 0xf4, 0xed, 0xf8, 0x89, 0xee, 0x27, 0x29, 0x33, 0x82, 0x02, 0xd1, 0x64,
				0x83, 0x1e, 0x2a, 0xb4, 0x3d, 0x40, 0xd2, 0xa2, 0x6d, 0x73, 0x73, 0x95, 0x70, 0xcf, 0x81, 0x6c, 0xb9, 0x6d, 0x76, 0x6b, 0x8d, 0x38, 0x50, 0x25, 0x8d, 0x58, 0xc8, 0x9f,
				0x7e, 0x99, 0x01, 0xed, 0xf1, 0x3e, 0x80, 0xa8},
		},
		{
			name:     "15 words",
			mnemonic: "wedding physical genre ill segment junior engage mixture casino cricket license vast screen pond jacket",
			seed: func() []byte {
				bts, _ := hex.DecodeString("27345116e47b47bae2e54192808c6b8d1ee6fcc4b783b0983d3876c5d0b514d864566d097e120529d0fafdd747fbfeccaff6d9c0998186980e86d19f441108c8")
				return bts
			}(),
		},
		{
			name: "18 words", mnemonic: "obscure radio true scale minimum treat arrange rebel outside charge crowd bargain shock family bus tray short aim",
			seed: func() []byte {
				bts, _ := hex.DecodeString("ce1adc78da44f2bfe4084fce26aa2b351c5f687bce1a57fddbd8977a4dca9b5a6b1b948bdd60841e8188736202165c297b7e9f81a6c584732183d6c547db28b5")
				return bts
			}(),
		},
		{
			name:     "21 words",
			mnemonic: "ranch same oak comfort recipe you pen suffer slide trap lift floor silent job tiny tennis approve lyrics rain garment lucky",
			seed: func() []byte {
				bts, _ := hex.DecodeString("b227b5460a2ac7f5834742428709fb0c83ced9c01716fa0a5e803b2b288ed9e3face38bfc3a3f34e1e6621a6ad5900d2c4f3f1a4fc390d9bd7511d7f01d4a944")
				return bts
			}(),
		},
		{
			name:     "24 words",
			mnemonic: "armed fantasy witness similar prosper poet throw video cannon original video zone talk swear economy bachelor urban crunch mouse trial joy little smart marble",
			seed: func() []byte {
				bts, _ := hex.DecodeString("8d86bf3b05fc16c80a10b2f608d30e869d12ded99f1453013e834260abed23508422cdf451ba533a935947a05d63e56ce6432ac8c1a5748b04003b83cbbc541a")
				return bts
			}(),
		},
	}

	for _, test := range tests {
		seed, err := Bip39MnemonicToSeed(test.mnemonic, "")
		if err != nil {
			t.Errorf("%v: Bip39MnemonicToSeed err: %v", test.name, err)
			return
		}

		if !bytes.Equal(seed, test.seed) {
			t.Errorf("%v: Bip39MnemonicToSeed failed : want %v got %v", test.name, test.seed, seed)
			return
		}

	}
}

func TestBase58(t *testing.T) {
	s,_:=hex.DecodeString("02b1ad2bc0a9d189c4c644ac2668d62b2b6147ce")

	t.Log(base58.Encode(s))
	t.Log(dcrbase58.Encode(s))

}