package addrtool

import (
	"encoding/hex"
	"github.com/tyler-smith/go-bip39"
	"testing"
)

func TestGenMnemonic(t *testing.T) {
	//生成熵
	entropyBytes,_:=bip39.NewEntropy(128)
	t.Log("entropyBytes：",entropyBytes)

	//生成助记词
	mnemonic,_:=bip39.NewMnemonic(entropyBytes)
	t.Log("mnemonic：",mnemonic)
}
func TestMnemonicToSeed(t *testing.T) {
	mnemonic :="chef fiction deputy stage pudding pink skirt often decade drift music loop"
	//助记词生成种子 password 为空
	seed:=bip39.NewSeed(mnemonic,"")
	t.Log("seed：",hex.EncodeToString(seed))
}



