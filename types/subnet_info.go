package types

import (
	"bytes"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
)

type SubnetInfo struct {
	Netuid               UCompact
	Rho                  UCompact
	Kappa                UCompact
	Difficulty           UCompact
	ImmunityPeriod       UCompact // Number of blocks before a validator can be removed from the subnet
	MaxAllowedValidators UCompact // Maximum number of validators allowed on the subnet
	MinAllowedWeights    UCompact // Minimum weight allowed for a validator
	MaxWeightsLimit      UCompact // Maximum weight allowed for a validator
	ScalingLawPower      UCompact // Power of the scaling law
	SubnetworkN          UCompact
	MaxAllowedUids       UCompact
	BlocksSinceLastStep  UCompact
	Tempo                UCompact
	NetworkModality      UCompact
	NetworkConnect       [][]UCompact
	EmissionValues       UCompact
	Burn                 UCompact
	Owner                AccountID
}

func DecodeSubnetInfoFromBytes(b []byte) (SubnetInfo, error) {
	var subnetInfo SubnetInfo

	d := scale.NewDecoder(bytes.NewReader(b))
	err := d.Decode(&subnetInfo)
	if err != nil {
		return subnetInfo, err
	}

	return subnetInfo, nil
}

func DecodeSubnetsInfoFromBytes(b []byte) ([]SubnetInfo, error) {
	var subnetInfos []SubnetInfo

	d := scale.NewDecoder(bytes.NewReader(b))
	err := d.Decode(&subnetInfos)
	if err != nil {
		return nil, err
	}

	for i, subnet := range subnetInfos {
		subnetInfos[i].Netuid = NewUCompactFromUInt(uint64(subnet.Netuid.Int64() >> 8))
	}

	return subnetInfos, nil
}
