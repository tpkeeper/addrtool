package addrtool

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/tyler-smith/go-bip39"
	"testing"
)

func TestMnemonicToAddr(t *testing.T) {
	inputMnemonic := "tone canoe toward edge voyage reveal large ignore enough draft worry erosion"
	addr,err:=MnemonicToAddr(inputMnemonic,hcTestNetParams)
	if err!=nil{
		t.Error(err)
	}
	t.Log(addr)
}

func TestMnemonicToSeed(t *testing.T) {
	Mnemonic:="hunt toward echo expire local mystery robust success digital advice erase tail"
	//"cdd0493c7fd1bcd72ac9814db5218a7585b40a66968580846f7b983688b050e8"
	bts,_:=bip39.MnemonicToByteArray(Mnemonic,false)
	fmt.Println(hex.EncodeToString(bts))
}



func TestMnemonic(t *testing.T) {
	pwd := "123"
	inputMnemonic := "tone canoe toward edge voyage reveal large ignore enough draft worry erosion"
	fmt.Println("密码", pwd)
	fmt.Println("原始助记词", inputMnemonic)
	decryptBytes0, _ := bip39.MnemonicToByteArray(inputMnemonic, true)
	fmt.Println("原始种子", decryptBytes0)
	//mnemonicHex, _ := bip39.NewMnemonic(decryptBytes0)
	//fmt.Println(mnemonicHex)
	pwdbyte := sha256.Sum256([]byte(pwd))
	fmt.Println(pwdbyte[:])
	encryptByte, _ := AesEncrypt(decryptBytes0, pwdbyte[:16])
	fmt.Println("加密之后的种子", encryptByte)
	mnemonicEccrypt, _ := bip39.NewMnemonic(encryptByte)
	fmt.Println("加密之后的助记词", mnemonicEccrypt)
	decryptBytes, _ := bip39.MnemonicToByteArray(mnemonicEccrypt, true)
	//fmt.Println(decryptBytes)
	//fmt.Println(len(encryptByte),EncodeMnemonic(encryptByte))
	plainByte, _ := AesDecrypt(decryptBytes, pwdbyte[:16])
	//palinByte:=make([]byte,len(encryptByte))
	//for i,byte:=range encryptByte{
	//	palinByte[i]=uint8((uint16(byte)-uint16(i))%uint16(256))
	//}
	fmt.Println("恢复的原始种子", plainByte)
	mnemonicPlain, _ := bip39.NewMnemonic(plainByte)
	fmt.Println("恢复的原始助记词", mnemonicPlain)
	//fmt.Println(len(plainByte),EncodeMnemonic(plainByte))

	for i:=0;i<2047;i++{
		for j:=i+1;j<2048;j++  {
			if len(wordList[i])>=4&&len(wordList[j])>=4{

			if string(wordList[i][:4])==string(wordList[j][:4]){
				print(i,j,wordList[i])
			}
			}
		}
	}
}
