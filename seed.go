package addrtool

import "github.com/tyler-smith/go-bip32"

// The hierarchy described by BIP0043 is:
//  m/<purpose>'/*
// This is further extended by BIP0044 to:
//  m/44'/<coin type>'/<account>'/<branch>/<address index>
//
// The branch is 0 for external addresses and 1 for internal addresses.
//  m/purpose(44)'/coinType(171)'/account(0)'/change(0)/index(0)


//return compressed pubkey
func SeedToPubKey(seed []byte,purpose uint32, coinType uint32,account uint32, change uint32, index uint32) ([]byte, error) {
	masterExtKey, err := bip32.NewMasterKey(seed)
	if err!=nil{
		return nil,err
	}
	//m/purpose'
	purposeExtKey,err:=masterExtKey.NewChildKey(bip32.FirstHardenedChild+purpose)
	if err!=nil{
		return nil,err
	}
	//m/purpose'/cointype'
	coinTypeExtKey,err:=purposeExtKey.NewChildKey(bip32.FirstHardenedChild+coinType)
	if err!=nil{
		return nil,err
	}
	//m/purpose'/cointype'/account'
	accountExtKey,err:=coinTypeExtKey.NewChildKey(bip32.FirstHardenedChild+account)
	if err!=nil{
		return nil,err
	}
	//m/purpose'/cointype'/account'/change
	changeExtKey,err:=accountExtKey.NewChildKey(change)
	if err!=nil{
		return nil,err
	}
	//m/purpose'/cointype'/account'/change/addrIndex
	addrIndex0ExtKey,err:=changeExtKey.NewChildKey(index)
	if err!=nil{
		return nil,err
	}
	return addrIndex0ExtKey.PublicKey().Key,nil
}


//return compressed ecc pubkey or children compressed ecc pubkey from an extended key base58-encoded
func PubkeyFromExtendKey(extendKeyB58Str string, children ... uint32)([]byte,error)  {

	extKey,err:=bip32.B58Deserialize(extendKeyB58Str)

	//node,err:=hdkeychain.NewKeyFromString(extendKeyStr,nwp)
	if err!=nil{
		return nil,err
	}

	for _,child:=range children {
		extKey,err = extKey.NewChildKey(child)
		if err!=nil{
			return nil,err
		}
	}

	pubkey:=extKey.PublicKey()

	return pubkey.Key,nil
}