package sep24

import anchorutils "github.com/tryoutbounder/anchor-go/common"

type TransactionRequest struct {
	Id                    string `url:"id,omitempty"`
	StellarTransactionId  string `url:"stellar_transaction_id,omitempty"`
	ExternalTransactionId string `url:"external_transaction_id,omitempty"`
	Lang                  string `url:"lang,omitempty"`
}

type TransactionResponse struct {
	Transaction Sep24Transaction `json:"transaction"`
}

// Shared transaction object used in responses
type Sep24Transaction struct {
	anchorutils.BaseTransaction

	// SEP-24 specific timestamp fields (string)
	StartedAt            string `json:"started_at"`
	UpdatedAt            string `json:"updated_at,omitempty"`
	CompletedAt          string `json:"completed_at,omitempty"`
	UserActionRequiredBy string `json:"user_action_required_by,omitempty"`

	// SEP-24 specific refunds type
	Refunds *RefundsObject `json:"refunds,omitempty"`

	// SEP-24 only fields
	KycVerified bool `json:"kyc_verified,omitempty"`
}

type FeeDetails struct {
	Total     string         `json:"total"`
	Asset     string         `json:"asset"`
	Breakdown []FeeBreakdown `json:"breakdown,omitempty"`
}

type FeeBreakdown struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Amount      string `json:"amount"`
}

type RefundsObject struct {
	AmountRefunded string          `json:"amount_refunded"`
	AmountFee      string          `json:"amount_fee"`
	Payments       []RefundPayment `json:"payments"`
}

type RefundPayment struct {
	Id     string `json:"id"`
	IdType string `json:"id_type"`
	Amount string `json:"amount"`
	Fee    string `json:"fee"`
}

// Transaction history request/response
type TransactionsRequest struct {
	AssetCode   string `url:"asset_code"`
	NoOlderThan string `url:"no_older_than,omitempty"`
	Limit       int    `url:"limit,omitempty"`
	Kind        string `url:"kind,omitempty"`
	PagingId    string `url:"paging_id,omitempty"`
	Lang        string `url:"lang,omitempty"`
}

type TransactionsResponse struct {
	Transactions []Sep24Transaction `json:"transactions"` // Refers to the Transaction struct in transaction.go
}
