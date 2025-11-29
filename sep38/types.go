package sep38

import "time"

type Info struct {
	Assets []Asset `json:"assets"`
}

type Asset struct {
	Asset               string           `json:"asset"`
	SellDeliveryMethods []DeliveryMethod `json:"sell_delivery_methods,omitempty"`
	BuyDeliveryMethods  []DeliveryMethod `json:"buy_delivery_methods,omitempty"`
	CountryCodes        []string         `json:"country_codes,omitempty"`
}

type DeliveryMethod struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetPricesRequest struct {
	SellAsset          string `url:"sell_asset"`
	BuyAsset           string `url:"buy_asset"`
	SellAmount         string `url:"sell_amount"`
	BuyAmount          string `url:"buy_amount"`
	SellDeliveryMethod string `url:"sell_delivery_method"`
	BuyDeliveryMethod  string `url:"buy_delivery_method"`
	CountryCode        string `url:"country_code"`
}

type GetPricesResponse struct {
	BuyAssets  []AssetPrice `json:"buy_assets,omitempty"`
	SellAssets []AssetPrice `json:"sell_assets,omitempty"`
}

type AssetPrice struct {
	Asset    string `json:"asset"`
	Price    string `json:"price"`
	Decimals int    `json:"decimals"`
}

type GetPriceRequest struct {
	SellAsset          string `url:"sell_asset"`
	BuyAsset           string `url:"buy_asset"`
	SellAmount         string `url:"sell_amount,omitempty"`
	BuyAmount          string `url:"buy_amount,omitempty"`
	SellDeliveryMethod string `url:"sell_delivery_method"`
	BuyDeliveryMethod  string `url:"buy_delivery_method"`
	CountryCode        string `url:"country_code"`
	Context            string `url:"context"`
}

type GetPriceResponse struct {
	TotalPrice string `json:"total_price"`
	Price      string `json:"price"`
	SellAmount string `json:"sell_amount"`
	BuyAmount  string `json:"buy_amount"`
	Fee        Fee    `json:"fee"`
}

type Fee struct {
	Total   string      `json:"total"`
	Asset   string      `json:"asset"`
	Details []FeeDetail `json:"details,omitempty"`
}

type FeeDetail struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Amount      string `json:"amount"`
}

type PostQuoteRequest struct {
	SellAsset          string    `json:"sell_asset"`
	BuyAsset           string    `json:"buy_asset"`
	SellAmount         string    `json:"sell_amount"`
	BuyAmount          string    `json:"buy_amount"`
	ExpireAfter        time.Time `json:"expire_after,omitzero"`
	SellDeliveryMethod string    `json:"sell_delivery_method,omitempty"`
	BuyDeliveryMethod  string    `json:"buy_delivery_method,omitempty"`
	CountryCode        string    `json:"country_code,omitempty"`
	Context            string    `json:"context"`
}

type QuoteResponse struct {
	ID                 string    `json:"id"`
	ExpiresAt          time.Time `json:"expires_at"`
	TotalPrice         string    `json:"total_price"`
	Price              string    `json:"price"`
	SellAsset          string    `json:"sell_asset"`
	SellAmount         string    `json:"sell_amount"`
	SellDeliveryMethod string    `json:"sell_delivery_method,omitempty"`
	BuyAsset           string    `json:"buy_asset"`
	BuyAmount          string    `json:"buy_amount"`
	BuyDeliveryMethod  string    `json:"buy_delivery_method,omitempty"`
	Fee                Fee       `json:"fee"`
}
