package sep24

// Withdrawal request/response
type WithdrawRequest struct {
	AssetCode        string  `json:"asset_code"`
	AssetIssuer      string  `json:"asset_issuer,omitempty"`
	DestinationAsset string  `json:"destination_asset,omitempty"`
	Amount           float64 `json:"amount,omitempty"`
	QuoteId          string  `json:"quote_id,omitempty"`
	Account          string  `json:"account,omitempty"`
	Memo             string  `json:"memo,omitempty"`
	MemoType         string  `json:"memo_type,omitempty"`
	WalletName       string  `json:"wallet_name,omitempty"`
	WalletUrl        string  `json:"wallet_url,omitempty"`
	Lang             string  `json:"lang,omitempty"`
	RefundMemo       string  `json:"refund_memo,omitempty"`
	RefundMemoType   string  `json:"refund_memo_type,omitempty"`
	CustomerId       string  `json:"customer_id,omitempty"`
}

type WithdrawResponse struct {
	Type string `json:"type"`
	Url  string `json:"url,omitempty"`
	ID   string `json:"id,omitempty"`
}
