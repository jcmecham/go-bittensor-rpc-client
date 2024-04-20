package subnetinfo

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type SubnetInfo interface {
	GetSubnetInfo(subnetUid uint16, blockHash types.Hash) (types.SubnetInfo, error)
	GetSubnetInfoLatest(subnetUid uint16) (types.SubnetInfo, error)

	GetSubnetsInfo(blockHash types.Hash) ([]types.SubnetInfo, error)
	GetSubnetsInfoLatest() ([]types.SubnetInfo, error)
}

type subnetInfo struct {
	client client.Client
}

func NewSubnetInfo(client client.Client) SubnetInfo {
	return &subnetInfo{
		client: client,
	}
}
