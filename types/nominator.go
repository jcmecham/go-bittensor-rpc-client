package types

// Nominator is a struct that contains a Delegator Hotkey and the amount staked for a specific Validator
type Nominator struct {
	NominatorSS58 AccountID
	StakeAmount   UCompact
}
