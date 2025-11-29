package sep6

// Info represents basic info about the TRANSFER_SERVER capabilities.
type Info struct {
	Deposit          map[string]DepositAsset          `json:"deposit,omitempty"`
	DepositExchange  map[string]DepositExchangeAsset  `json:"deposit-exchange,omitempty"`
	Withdraw         map[string]WithdrawAsset         `json:"withdraw,omitempty"`
	WithdrawExchange map[string]WithdrawExchangeAsset `json:"withdraw-exchange,omitempty"`
	Fee              FeeInfo                          `json:"fee"`
	Transactions     EndpointInfo                     `json:"transactions"`
	Transaction      EndpointInfo                     `json:"transaction"`
	Features         Features                         `json:"features"`
}

// DepositAsset represents info about a single deposit asset.
type DepositAsset struct {
	Enabled                bool     `json:"enabled"`
	AuthenticationRequired bool     `json:"authentication_required,omitempty"`
	MinAmount              float64  `json:"min_amount,omitempty"`
	MaxAmount              float64  `json:"max_amount,omitempty"`
	FeeFixed               float64  `json:"fee_fixed,omitempty"`
	FeePercent             float64  `json:"fee_percent,omitempty"`
	FundingMethods         []string `json:"funding_methods,omitempty"`
}

// DepositExchangeAsset represents info about a single deposit-exchange asset.
type DepositExchangeAsset struct {
	AuthenticationRequired bool     `json:"authentication_required"`
	FundingMethods         []string `json:"funding_methods,omitempty"`
}

// WithdrawAsset represents info about a single withdraw asset.
type WithdrawAsset struct {
	Enabled                bool     `json:"enabled"`
	AuthenticationRequired bool     `json:"authentication_required,omitempty"`
	MinAmount              float64  `json:"min_amount,omitempty"`
	MaxAmount              float64  `json:"max_amount,omitempty"`
	FeeFixed               float64  `json:"fee_fixed,omitempty"`
	FeePercent             float64  `json:"fee_percent,omitempty"`
	FundingMethods         []string `json:"funding_methods,omitempty"`
}

// WithdrawExchangeAsset represents info about a single withdraw-exchange asset.
type WithdrawExchangeAsset struct {
	AuthenticationRequired bool     `json:"authentication_required"`
	MinAmount              float64  `json:"min_amount,omitempty"`
	MaxAmount              float64  `json:"max_amount,omitempty"`
	FundingMethods         []string `json:"funding_methods,omitempty"`
}

// FeeInfo represents info about fees.
type FeeInfo struct {
	Enabled     bool   `json:"enabled"`
	Description string `json:"description,omitempty"`
}

// EndpointInfo represents info about an endpoint.
type EndpointInfo struct {
	Enabled                bool `json:"enabled"`
	AuthenticationRequired bool `json:"authentication_required"`
}

// Features represents supported feature flags.
type Features struct {
	AccountCreation   bool `json:"account_creation"`
	ClaimableBalances bool `json:"claimable_balances"`
}

// Fee represents the fee response fields.
type Fee struct {
	Fee float64 `json:"fee"`
}

// FeeRequest represents fee request parameters.
type FeeRequest struct {
	Operation string  `schema:"operation"`
	Type      string  `schema:"type,omitempty"`
	AssetCode string  `schema:"asset_code"`
	Amount    float64 `schema:"amount"`
}
