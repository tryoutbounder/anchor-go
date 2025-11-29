package sep6

// WithdrawRequest represents a request to withdraw assets from the anchor
type WithdrawRequest struct {
	AssetCode        string `url:"asset_code"`
	FundingMethod    string `url:"funding_method"`
	Account          string `url:"account,omitempty"`
	Memo             string `url:"memo,omitempty"`
	Lang             string `url:"lang,omitempty"`
	OnChangeCallback string `url:"on_change_callback,omitempty"`
	Amount           string `url:"amount,omitempty"`
	CountryCode      string `url:"country_code,omitempty"`
	RefundMemo       string `url:"refund_memo,omitempty"`
	RefundMemoType   string `url:"refund_memo_type,omitempty"`
	CustomerID       string `url:"customer_id,omitempty"`
	LocationID       string `url:"location_id,omitempty"`
	Type             string `url:"type,omitempty"`
	Dest             string `url:"dest,omitempty"`
	DestExtra        string `url:"dest_extra,omitempty"`
	MemoType         string `url:"memo_type,omitempty"`
	WalletName       string `url:"wallet_name,omitempty"`
	WalletURL        string `url:"wallet_url,omitempty"`
}

// WithdrawResponse represents a successful response from the withdraw endpoint
type WithdrawResponse struct {
	AccountID  string         `json:"account_id,omitempty"`
	MemoType   string         `json:"memo_type,omitempty"`
	Memo       string         `json:"memo,omitempty"`
	ID         string         `json:"id,omitempty"`
	ETA        int            `json:"eta,omitempty"`
	MinAmount  float64        `json:"min_amount,omitempty"`
	MaxAmount  float64        `json:"max_amount,omitempty"`
	FeeFixed   float64        `json:"fee_fixed,omitempty"`
	FeePercent float64        `json:"fee_percent,omitempty"`
	ExtraInfo  map[string]any `json:"extra_info,omitempty"`
}

// WithdrawExchangeRequest represents a request to withdraw and exchange assets from the anchor
type WithdrawExchangeRequest struct {
	SourceAsset      string `url:"source_asset"`
	DestinationAsset string `url:"destination_asset"`
	Amount           string `url:"amount,omitempty"`
	FundingMethod    string `url:"funding_method"`
	QuoteID          string `url:"quote_id,omitempty"`
	Account          string `url:"account,omitempty"`
	Memo             string `url:"memo,omitempty"`
	Lang             string `url:"lang,omitempty"`
	OnChangeCallback string `url:"on_change_callback,omitempty"`
	CountryCode      string `url:"country_code,omitempty"`
	RefundMemo       string `url:"refund_memo,omitempty"`
	RefundMemoType   string `url:"refund_memo_type,omitempty"`
	CustomerID       string `url:"customer_id,omitempty"`
	LocationID       string `url:"location_id,omitempty"`
	Type             string `url:"type,omitempty"`
	Dest             string `url:"dest,omitempty"`
	DestExtra        string `url:"dest_extra,omitempty"`
	MemoType         string `url:"memo_type,omitempty"`
	WalletName       string `url:"wallet_name,omitempty"`
	WalletURL        string `url:"wallet_url,omitempty"`
}

// WithdrawExchangeResponse represents a successful response from the withdraw-exchange endpoint
type WithdrawExchangeResponse struct {
	AccountID  string         `json:"account_id,omitempty"`
	MemoType   string         `json:"memo_type,omitempty"`
	Memo       string         `json:"memo,omitempty"`
	ID         string         `json:"id,omitempty"`
	ETA        int            `json:"eta,omitempty"`
	MinAmount  float64        `json:"min_amount,omitempty"`
	MaxAmount  float64        `json:"max_amount,omitempty"`
	FeeFixed   float64        `json:"fee_fixed,omitempty"`
	FeePercent float64        `json:"fee_percent,omitempty"`
	ExtraInfo  map[string]any `json:"extra_info,omitempty"`
}
