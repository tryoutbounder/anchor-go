package sep24

// Deposit request/response
type DepositRequest struct {
	AssetCode                 string  `json:"asset_code"`
	AssetIssuer               string  `json:"asset_issuer,omitempty"`
	SourceAsset               string  `json:"source_asset,omitempty"`
	Amount                    float64 `json:"amount,omitempty"`
	QuoteId                   string  `json:"quote_id,omitempty"`
	Account                   string  `json:"account,omitempty"`
	MemoType                  string  `json:"memo_type,omitempty"`
	Memo                      string  `json:"memo,omitempty"`
	WalletName                string  `json:"wallet_name,omitempty"`
	WalletUrl                 string  `json:"wallet_url,omitempty"`
	Lang                      string  `json:"lang,omitempty"`
	ClaimableBalanceSupported bool    `json:"claimable_balance_supported,omitempty"`
	CustomerId                string  `json:"customer_id,omitempty"`
}

type DepositResponse struct {
	Type string `json:"type"`
	Url  string `json:"url,omitempty"`
	Id   string `json:"id,omitempty"`
}
