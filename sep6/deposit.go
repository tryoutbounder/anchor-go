package sep6

import anchorutils "github.com/tryoutbounder/anchor-go/common"

type DepositRequest struct {
	AssetCode                 string `url:"asset_code"`
	Account                   string `url:"account"`
	FundingMethod             string `url:"funding_method"`
	MemoType                  string `url:"memo_type,omitempty"`
	Memo                      string `url:"memo,omitempty"`
	EmailAddress              string `url:"email_address,omitempty"`
	Lang                      string `url:"lang,omitempty"`
	OnChangeCallback          string `url:"on_change_callback,omitempty"`
	Amount                    string `url:"amount,omitempty"`
	CountryCode               string `url:"country_code,omitempty"`
	ClaimableBalanceSupported string `url:"claimable_balance_supported,omitempty"`
	CustomerId                string `url:"customer_id,omitempty"`
	LocationId                string `url:"location_id,omitempty"`
	Type                      string `url:"type,omitempty"`
	WalletName                string `url:"wallet_name,omitempty"`
	WalletUrl                 string `url:"wallet_url,omitempty"`
}

type DepositInstructionField struct {
	Value       string `json:"value"`
	Description string `json:"description"`
}

type DepositInstructions struct {
	anchorutils.Sep9OrganizationalFields
}

type DepositExtraInfo struct {
	Message string `json:"message,omitempty"`
}

type DepositResponse struct {
	How          string              `json:"how,omitempty"`
	Instructions DepositInstructions `json:"instructions"`
	ID           string              `json:"id,omitempty"`
	Eta          int                 `json:"eta,omitempty"`
	MinAmount    float64             `json:"min_amount,omitempty"`
	MaxAmount    float64             `json:"max_amount,omitempty"`
	FeeFixed     float64             `json:"fee_fixed,omitempty"`
	FeePercent   float64             `json:"fee_percent,omitempty"`
	ExtraInfo    DepositExtraInfo    `json:"extra_info"`
}

type DepositExchangeRequest struct {
	DestinationAsset          string `url:"destination_asset"`
	SourceAsset               string `url:"source_asset"`
	Amount                    string `url:"amount,omitempty"`
	FundingMethod             string `url:"funding_method"`
	Account                   string `url:"account"`
	QuoteId                   string `url:"quote_id,omitempty"`
	MemoType                  string `url:"memo_type,omitempty"`
	Memo                      string `url:"memo,omitempty"`
	EmailAddress              string `url:"email_address,omitempty"`
	Lang                      string `url:"lang,omitempty"`
	OnChangeCallback          string `url:"on_change_callback,omitempty"`
	CountryCode               string `url:"country_code,omitempty"`
	ClaimableBalanceSupported string `url:"claimable_balance_supported,omitempty"`
	CustomerId                string `url:"customer_id,omitempty"`
	LocationId                string `url:"location_id,omitempty"`
	Type                      string `url:"type,omitempty"`
	WalletName                string `url:"wallet_name,omitempty"`
	WalletUrl                 string `url:"wallet_url,omitempty"`
}

type DepositExchangeResponse struct {
	How          string              `json:"how,omitempty"`
	Instructions DepositInstructions `json:"instructions"`
	ID           string              `json:"id,omitempty"`
	Eta          int                 `json:"eta,omitempty"`
	MinAmount    float64             `json:"min_amount,omitempty"`
	MaxAmount    float64             `json:"max_amount,omitempty"`
	FeeFixed     float64             `json:"fee_fixed,omitempty"`
	FeePercent   float64             `json:"fee_percent,omitempty"`
	ExtraInfo    DepositExtraInfo    `json:"extra_info"`
}
