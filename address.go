package addrtool

import (
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
)

func PubkeyToAddress(key []byte,netId byte)(string){
	hash160Bytes:=btcutil.Hash160(key)
	return base58.CheckEncode(hash160Bytes,netId)
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
