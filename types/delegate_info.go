package types

import (
	"bytes"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
)

type DelegateInfo struct {
	DelegateSS58     AccountID
	Take             UCompact
	Nominators       []Nominator
	OwnerSS58        AccountID
	Registrations    []UCompact
	ValidatorPermits []UCompact
	ReturnPer1000    UCompact
	TotalDailyReturn UCompact
}

func DecodeDelegateFromBytes(b []byte) (DelegateInfo, error) {
	var delegate DelegateInfo

	d := scale.NewDecoder(bytes.NewReader(b))
	err := d.Decode(&delegate)
	if err != nil {
		return delegate, err
	}

	return delegate, nil
}

func DecodeDelegatesFromBytes(b []byte) ([]DelegateInfo, error) {
	var delegates []DelegateInfo

	d := scale.NewDecoder(bytes.NewReader(b))
	err := d.Decode(&delegates)
	if err != nil {
		return nil, err
	}

	return delegates, nil
}
