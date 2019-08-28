package addrtool

import (
	"github.com/decred/base58"
	"github.com/decred/dcrd/dcrutil/v2"
	"github.com/decred/dcrd/hdkeychain/v2"
)

func SeedToAddr(seed []byte,nwp *NetWorkParams)(string,error) {
	masterNode, err := hdkeychain.NewMaster(seed, nwp)
	purpose, err := masterNode.Child(44 + hdkeychain.HardenedKeyStart)
	if err != nil {
		return "",err
	}
	// Derive the coin type key as a child of the purpose key.
	coinTypeKey, err := purpose.Child(nwp.HDCoinType + hdkeychain.HardenedKeyStart)
	if err != nil {
		return "",err
	}
	account0 := uint32(0)
	acct0Key, err := coinTypeKey.Child(account0 + hdkeychain.HardenedKeyStart)
	change0key, err := acct0Key.Child(0)
	// The hierarchy described by BIP0043 is:
	//  m/<purpose>'/*
	// This is further extended by BIP0044 to:
	//  m/44'/<coin type>'/<account>'/<branch>/<address index>
	//
	// The branch is 0 for external addresses and 1 for internal addresses.
	//  m/purpose(44)'/coinType(171)'/account(0)'/change(0)/index(0)
	index0key, err := change0key.Child(0)
	pubKey, err := index0key.ECPubKey()
	hash160Byte := dcrutil.Hash160(pubKey.SerializeCompressed())
	address := base58.CheckEncode(hash160Byte, nwp.PubKeyHashAddrID)
	return address,nil
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

