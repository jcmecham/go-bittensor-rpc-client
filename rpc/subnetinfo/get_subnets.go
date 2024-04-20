package subnetinfo

import (
	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// GetSubnetsInfoLatest	returns the latest SubnetInfo for all the subnets.
func (s *subnetInfo) GetSubnetsInfoLatest() ([]types.SubnetInfo, error) {
	return s.getSubnets(nil)
}

// GetSubnetsInfo returns the SubnetInfo for all the subnets at the given block number.
func (s *subnetInfo) GetSubnetsInfo(blockHash types.Hash) ([]types.SubnetInfo, error) {
	return s.getSubnets(&blockHash)
}

func (s *subnetInfo) getSubnets(blockHash *types.Hash) ([]types.SubnetInfo, error) {
	var subnetsInfo []byte
	var err error

	if blockHash == nil {
		err = s.client.Call(&subnetsInfo, "subnetInfo_getSubnetsInfo")
	} else {
		err = client.CallWithBlockHash(s.client, &subnetsInfo, "subnetInfo_getSubnetsInfo", blockHash)
	}
	if err != nil {
		return nil, err
	}

	return types.DecodeSubnetsInfoFromBytes(subnetsInfo)
}
