package delegateinfo

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type DelegateInfo interface {

	// GetDelegate returns a delegate (Validator) at a specific block.
	GetDelegate(ValidatorHotKey types.AccountID, blockHash types.Hash) (types.DelegateInfo, error)

	// GetDelegateLatest returns a delegate (Validator) at the latest block.
	GetDelegateLatest(ValidatorHotKey types.AccountID) (types.DelegateInfo, error)

	// GetDelegates returns all delegates (Validators) at a specific block.
	GetDelegates(blockHash types.Hash) ([]types.DelegateInfo, error)

	// GetDelegatesLatest returns all delegates (Validators) at the latest block.
	GetDelegatesLatest() ([]types.DelegateInfo, error)
}

type delegateInfo struct {
	client client.Client
}

func NewDelegateInfo(client client.Client) DelegateInfo {
	return &delegateInfo{
		client: client,
	}
}
