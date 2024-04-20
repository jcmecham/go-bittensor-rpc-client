package delegateinfo

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// GetDelegatesLatest returns a specific delegate at the latest block
func (d *delegateInfo) GetDelegateLatest(ValidatorHotKey types.AccountID) (types.DelegateInfo, error) {
	return d.getDelegate(&ValidatorHotKey, nil)
}

// GetDelegate returns a specific delegate at a specific block
func (d *delegateInfo) GetDelegate(ValidatorHotKey types.AccountID, blockHash types.Hash) (types.DelegateInfo, error) {
	return d.getDelegate(&ValidatorHotKey, &blockHash)
}

func (d *delegateInfo) getDelegate(ValidatorHotKey *types.AccountID, blockHash *types.Hash) (types.DelegateInfo, error) {
	var res []byte
	var err error

	addressByteString, err := ValidatorHotKey.ToByteString()
	if err != nil {
		return types.DelegateInfo{}, err
	}

	if blockHash == nil {
		err = d.client.Call(&res, "delegateInfo_getDelegate", addressByteString)
	} else {
		err = client.CallWithBlockHash(d.client, &res, "delegateInfo_getDelegate", blockHash, addressByteString)
	}
	if err != nil {
		return types.DelegateInfo{}, err
	}

	return types.DecodeDelegateFromBytes(res)
}
