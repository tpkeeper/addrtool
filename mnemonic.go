package addrtool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/tyler-smith/go-bip39"
)

func mnemonicTest() {
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
}
//026fc0084f3af4509b1b744be9a0b912b2889601e7e858d2305dfeae2dae585cbb

func Md5(source string) string {
	md5h := md5.New()
	md5h.Write([]byte(source))
	return hex.EncodeToString(md5h.Sum(nil))
}
func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}


func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	//origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	//origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}


