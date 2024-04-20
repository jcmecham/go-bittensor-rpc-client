package delegateinfo

import types "github.com/centrifuge/go-substrate-rpc-client/v4/types"

// GetDelegatesLatest returns all delegates (Validators) at the latest block.
func (d *delegateInfo) GetDelegatesLatest() ([]types.DelegateInfo, error) {
	return d.getDelegates(nil)
}

// GetDelegates returns all delegates (Validators) at a specific block.
func (d *delegateInfo) GetDelegates(blockHash types.Hash) ([]types.DelegateInfo, error) {
	return d.getDelegates(&blockHash)
}

// getDelegates retrieves the delegates (Validators) at a specific block.
// If blockHash is nil, it retrieves the delegates at the latest block.
func (d *delegateInfo) getDelegates(blockHash *types.Hash) ([]types.DelegateInfo, error) {
	var res []byte
	var err error

	if blockHash == nil {
		err = d.client.Call(&res, "delegateInfo_getDelegates")
	} else {
		err = d.client.Call(&res, "delegateInfo_getDelegates", *blockHash)
	}
	if err != nil {
		return nil, err
	}

	return types.DecodeDelegatesFromBytes(res)
}
