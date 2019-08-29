package addrtool

import (
	"fmt"
	"github.com/decred/base58"
	"github.com/decred/dcrd/dcrutil/v2"
	"github.com/decred/dcrd/hdkeychain/v2"
	"github.com/tyler-smith/go-bip39"
)

func SeedToAddr(seed []byte,nwp *NetWorkParams)(string,error) {
	masterNode, err := hdkeychain.NewMaster(seed, nwp)
	purpose, err := masterNode.Child(45 + hdkeychain.HardenedKeyStart)

	PrintPubkey(purpose)

	if err != nil {
		return "",err
	}
	// Derive the coin type key as a child of the purpose key.
	coinTypeKey, err := purpose.Child(nwp.HDCoinType + hdkeychain.HardenedKeyStart)
	if err != nil {
		return "",err
	}

	PrintPubkey(coinTypeKey)

	account0 := uint32(0)
	acct0Key, err := coinTypeKey.Child(account0 + hdkeychain.HardenedKeyStart)

	PrintPubkey(acct0Key)

	change0key, err := acct0Key.Child(0)

	PrintPubkey(change0key)
	// The hierarchy described by BIP0043 is:
	//  m/<purpose>'/*
	// This is further extended by BIP0044 to:
	//  m/44'/<coin type>'/<account>'/<branch>/<address index>
	//
	// The branch is 0 for external addresses and 1 for internal addresses.
	//  m/purpose(44)'/coinType(171)'/account(0)'/change(0)/index(0)
	index0key, err := change0key.Child(0)

	PrintPubkey(index0key)

	pubKey, err := index0key.ECPubKey()
	hash160Byte := dcrutil.Hash160(pubKey.SerializeCompressed())
	address := base58.CheckEncode(hash160Byte, nwp.PubKeyHashAddrID)

	fmt.Println(address)
	return address,nil
}

func PrintPubkey(key *hdkeychain.ExtendedKey)  {
	pub,_:=key.Neuter()
	fmt.Println(pub)
}


func MnemonicToAddr(words string,nwp *NetWorkParams)(string,error)  {
	seed, err := bip39.MnemonicToByteArray(words, true)
	if err!=nil{
		return "",err
	}
	return SeedToAddr(seed,nwp)
}


type NetWorkParams struct {
	HDPrivateKeyID [4]byte
	HDPublicKeyID [4]byte
	HDCoinType uint32

	PubKeyAddrID         [2]byte
	PubKeyHashAddrID     [2]byte
	PKHEdwardsAddrID     [2]byte
	PKHSchnorrAddrID     [2]byte
	ScriptHashAddrID     [2]byte
	PrivateKeyID         [2]byte

	Address func()string
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

