package addrtool

import (
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	dcrbase58 "github.com/decred/base58"
	"github.com/decred/dcrd/dcrutil"
)

func PubkeyToAddress(key []byte, netId byte) (string) {
	hash160Bytes := btcutil.Hash160(key)
	return base58.CheckEncode(hash160Bytes[:ripemd160.Size], netId)
}
func DcrPubkeyToAddress(key []byte, netId [2]byte) (string) {
	//dcr的hash160与btc的不同
	hash160Bytes := dcrutil.Hash160(key)
	return dcrbase58.CheckEncode(hash160Bytes[:ripemd160.Size], netId)
}

func MultiPubkeyToAddress(netId byte, nRequired int, keys ... []byte) (string, error) {
	builder := txscript.NewScriptBuilder().AddInt64(int64(nRequired))
	for _, key := range keys {
		builder.AddData(key)
	}
	builder.AddInt64(int64(len(keys)))
	builder.AddOp(txscript.OP_CHECKMULTISIG)
	script, err := builder.Script()
	if err != nil {
		return "", err
	}
	scriptHash := btcutil.Hash160(script)
	return base58.CheckEncode(scriptHash[:ripemd160.Size], netId), nil
}

func DcrMultiPubkeyToAddress(netId [2]byte, nRequired int, keys [][]byte) (string, error) {
	builder := txscript.NewScriptBuilder().AddInt64(int64(nRequired))
	for _, key := range keys {
		builder.AddData(key)
	}
	builder.AddInt64(int64(len(keys)))
	builder.AddOp(txscript.OP_CHECKMULTISIG)
	script, err := builder.Script()
	if err != nil {
		return "", err
	}
	scriptHash := btcutil.Hash160(script)
	return dcrbase58.CheckEncode(scriptHash[:ripemd160.Size], netId), nil
}

type DcrNetWorkParams struct {
	HDPrivateKeyID [4]byte
	HDPublicKeyID  [4]byte
	HDCoinType     uint32

	PubKeyAddrID     [2]byte
	PubKeyHashAddrID [2]byte
	PKHEdwardsAddrID [2]byte
	PKHSchnorrAddrID [2]byte
	ScriptHashAddrID [2]byte
	PrivateKeyID     [2]byte

}
