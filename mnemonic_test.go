package addrtool

import (
	"encoding/hex"
	"testing"
)

func TestGenMnemonicBip39(t *testing.T) {
	mnemonic,_:=Bip39GenMnemonic(128)
	t.Log("mnemonic：", mnemonic)
}

func TestMnemonicToSeedBip39(t *testing.T) {
	mnemonic := "chef fiction deputy stage pudding pink skirt often decade drift music loop"
	//助记词生成种子 password 为空
	seed, _ := Bip39MnemonicToSeed(mnemonic, "")
	if hex.EncodeToString(seed) != "04ef53d66b17fdfb6538c5d183f0b0569fc1c79d07f044f7670c3038aff411e5abcbe8c457b584d0c1e3504ab94fb311f9097a793c20dfc746a87087ed5dc119" {
		t.Error("Bip39MnemonicToSeed failed")
	}
}
