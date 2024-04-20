package subnetinfo

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// GetSubnetInfoLatest returns the latest SubnetInfo for the given Subnet UID.
func (s *subnetInfo) GetSubnetInfoLatest(subnetUid uint16) (types.SubnetInfo, error) {
	return s.getSubnet(subnetUid, nil)
}

// GetSubnetInfo returns the SubnetInfo for the given Subnet UID at the given block number.
func (s *subnetInfo) GetSubnetInfo(subnetUid uint16, blockHash types.Hash) (types.SubnetInfo, error) {
	return s.getSubnet(subnetUid, &blockHash)
}

func (s *subnetInfo) getSubnet(subnetUid uint16, blockHash *types.Hash) (types.SubnetInfo, error) {
	var subnetInfo []byte
	var err error

	if blockHash == nil {
		err = s.client.Call(&subnetInfo, "subnetInfo_getSubnetInfo", subnetUid)
	} else {
		err = client.CallWithBlockHash(s.client, &subnetInfo, "subnetInfo_getSubnetInfo", blockHash, subnetUid)
	}
	if err != nil {
		return types.SubnetInfo{}, err
	}

	return types.DecodeSubnetInfoFromBytes(subnetInfo)
}
