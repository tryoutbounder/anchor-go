package sep24

// Fee represents the fee configuration
type Fee struct {
	Enabled bool `json:"enabled"`
}

// Features represents the available platform features
type Features struct {
	AccountCreation   bool `json:"account_creation"`
	ClaimableBalances bool `json:"claimable_balances"`
}

// Info represents the available assets and their info for deposits/withdrawals
type Info struct {
	Deposit  map[string]Asset `json:"deposit"`
	Withdraw map[string]Asset `json:"withdraw"`
	Fee      Fee              `json:"fee"`
	Features Features         `json:"features"`
}

// Asset represents the information for a deposit or withdrawal asset
type Asset struct {
	Enabled    bool    `json:"enabled"`
	FeeFixed   float64 `json:"fee_fixed,omitempty"`
	FeePercent float64 `json:"fee_percent,omitempty"`
	MinAmount  float64 `json:"min_amount,omitempty"`
	MaxAmount  float64 `json:"max_amount,omitempty"`
	FeeMinimum float64 `json:"fee_minimum,omitempty"`
}
