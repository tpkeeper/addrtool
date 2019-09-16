package addrtool

import (
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

func PubkeyToAddress(key []byte,netId byte)(string){
	hash160Bytes:=btcutil.Hash160(key)
	return base58.CheckEncode(hash160Bytes[:ripemd160.Size],netId)
}


func MultiPubkeyToAddress(netId byte,nRequired int,keys ... []byte) (string ,error){
	builder:=txscript.NewScriptBuilder().AddInt64(int64(nRequired))
	for _,key:=range keys{
		builder.AddData(key)
	}
	builder.AddInt64(int64(len(keys)))
	builder.AddOp(txscript.OP_CHECKMULTISIG)
	script,err:=builder.Script()
	if err!=nil{
		return "",err
	}
	scriptHash := btcutil.Hash160(script)
	return base58.CheckEncode(scriptHash[:ripemd160.Size], netId),nil
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
