package addrtool

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func TestDesEncrypt(t *testing.T) {
	key:=[]byte{0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01}
	cipherBlock,err:=des.NewCipher(key)
	if err!=nil{
		t.Error(err)
	}
	src:=[]byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08}
	encrptDst :=make([]byte,len(src))
	cipherBlock.Encrypt(encrptDst,src)
	t.Log(encrptDst)
	plainDst:=make([]byte,len(encrptDst))
	cipherBlock.Decrypt(plainDst, encrptDst)
	t.Log(plainDst)
}

func TestTripleDesEncrypt(t *testing.T) {
	key:=[]byte{0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,
	0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01}
	cipherBlock,err:=des.NewTripleDESCipher(key)
	if err!=nil{
		t.Error(err)
	}
	src:=[]byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08}
	encrptDst :=make([]byte,len(src))
	cipherBlock.Encrypt(encrptDst,src)
	t.Log(encrptDst)
	plainDst:=make([]byte,len(encrptDst))
	cipherBlock.Decrypt(plainDst, encrptDst)
	t.Log(plainDst)
}

func TestAesEncrypt(t *testing.T){
	key:=[]byte{0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01}
	cipherBlock,err:=aes.NewCipher(key)
	if err!=nil{
		t.Error(err)
	}
	src:=[]byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08}
	encrptDst :=make([]byte,len(src))
	cipherBlock.Encrypt(encrptDst,src)
	t.Log(encrptDst)
	plainDst:=make([]byte,len(encrptDst))
	cipherBlock.Decrypt(plainDst, encrptDst)
	t.Log(plainDst)
}

func TestCBCMode(t *testing.T) {
	key:=[]byte{0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01,0x01}
	cipherBlock,err:=aes.NewCipher(key)
	if err!=nil{
		t.Error(err)
	}
	src:=[]byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x01,0x02,0x03,
	0x04,0x05,0x06,0x07,0x08,0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08}
	inv:=[]byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08,0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08}
	cbcEncrypter:=cipher.NewCBCEncrypter(cipherBlock,inv)
	encrptDst :=make([]byte,len(src))
	cbcEncrypter.CryptBlocks(encrptDst,src)
	t.Log(encrptDst)

	plainDst:=make([]byte,len(encrptDst))
	cbcDecrypter:=cipher.NewCBCDecrypter(cipherBlock,inv)
	cbcDecrypter.CryptBlocks(plainDst,encrptDst)
	t.Log(plainDst)
}

func TestRsa(t *testing.T) {
	privKey,err:=rsa.GenerateKey(rand.Reader,2048)
	if err!=nil{
		t.Error(err)
	}
	msg:=[]byte("hello rsa")
	cryptMsg,err:=rsa.EncryptPKCS1v15(rand.Reader,&privKey.PublicKey,msg)
	if err!=nil{
		t.Error(err)
	}
	t.Log(cryptMsg)
	t.Log(string(cryptMsg))

	plainText,err:=rsa.DecryptPKCS1v15(rand.Reader,privKey,cryptMsg)
	if err!=nil{
		t.Error(plainText)
	}
	t.Log(string(plainText))

}

