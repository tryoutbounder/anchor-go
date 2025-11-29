package sep24

// Fee request/response
type FeeRequest struct {
	Operation string  `url:"operation"`
	Type      string  `url:"type,omitempty"`
	AssetCode string  `url:"asset_code"`
	Amount    float64 `url:"amount"`
}

type FeeResponse struct {
	Fee float64 `json:"fee"`
}
