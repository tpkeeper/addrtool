package addrtool

import (
	"encoding/hex"
	"fmt"
	"github.com/decred/base58"
	"github.com/decred/dcrd/dcrutil/v2"
	"github.com/decred/dcrd/hdkeychain/v2"
	"github.com/tyler-smith/go-bip39"
)

var debug = false

// The hierarchy described by BIP0043 is:
//  m/<purpose>'/*
// This is further extended by BIP0044 to:
//  m/44'/<coin type>'/<account>'/<branch>/<address index>
//
// The branch is 0 for external addresses and 1 for internal addresses.
//  m/purpose(44)'/coinType(171)'/account(0)'/change(0)/index(0)
func SeedToAddr(seed []byte, nwp *NetWorkParams, purpose uint32, account uint32, change uint32, index uint32) (string, error) {
	masterNode, err := hdkeychain.NewMaster(seed, nwp)
	purposeNode, err := masterNode.Child(purpose + hdkeychain.HardenedKeyStart)
	if err != nil {
		return "", err
	}

	PrintPubNode(purposeNode,"purpose")

	// Derive the coin type key as a child of the purpose key.
	coinTypeNode, err := purposeNode.Child(nwp.HDCoinType + hdkeychain.HardenedKeyStart)
	if err != nil {
		return "", err
	}

	PrintPubNode(coinTypeNode,"coinType")

	accountNode, err := coinTypeNode.Child(account + hdkeychain.HardenedKeyStart)

	PrintPubNode(accountNode,"account")

	changeNode, err := accountNode.Child(change)

	PrintPubNode(changeNode,"change")

	indexNode, err := changeNode.Child(index)

	PrintPubNode(indexNode,"index")

	pubKey, err := indexNode.ECPubKey()
	hash160Byte := dcrutil.Hash160(pubKey.SerializeCompressed())
	address := base58.CheckEncode(hash160Byte, nwp.PubKeyHashAddrID)

	return address, nil
}

//get serialized ecc pubkey or children serialized ecc pubkey from a base58-encoded extended key
func PubkeyFromNode(str string,nwp *NetWorkParams,isCompress bool, children ... uint32)(string,error)  {
	node,err:=hdkeychain.NewKeyFromString(str,nwp)
	if err!=nil{
		return "",err
	}

	for _,child:=range children {
		node,err = node.Child(child)
		if err!=nil{
			return "",err
		}
	}

	pubkey,err:=node.ECPubKey()
	if err!=nil{
		return "",err
	}

	var serializedBytes []byte
	if isCompress{
		serializedBytes =pubkey.SerializeCompressed()
	}else {
		serializedBytes = pubkey.SerializeUncompressed()
	}

	return hex.EncodeToString(serializedBytes),nil
}



func PrintPubNode(key *hdkeychain.ExtendedKey,layer string) {
	if debug {
		pub, _ := key.Neuter()
		fmt.Printf("%s   \t: %v\n",layer,pub)
	}
}

func MnemonicToAddr(words string, nwp *NetWorkParams) (string, error) {
	seed, err := bip39.MnemonicToByteArray(words, true)
	if err != nil {
		return "", err
	}
	return SeedToAddr(seed, nwp, 44, 0, 0, 0)
}

type NetWorkParams struct {
	HDPrivateKeyID [4]byte
	HDPublicKeyID  [4]byte
	HDCoinType     uint32

	PubKeyAddrID     [2]byte
	PubKeyHashAddrID [2]byte
	PKHEdwardsAddrID [2]byte
	PKHSchnorrAddrID [2]byte
	ScriptHashAddrID [2]byte
	PrivateKeyID     [2]byte

	Address func() string
}

// HDPrivKeyVersion returns the hierarchical deterministic extended private key
// magic version bytes for the network the parameters define.
func (p *NetWorkParams) HDPrivKeyVersion() [4]byte {
	return p.HDPrivateKeyID
}

// HDPubKeyVersion returns the hierarchical deterministic extended public key
// magic version bytes for the network the parameters define.
func (p *NetWorkParams) HDPubKeyVersion() [4]byte {
	return p.HDPublicKeyID
}
