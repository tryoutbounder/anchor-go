package sep6

import (
	"time"

	anchorutils "github.com/tryoutbounder/anchor-go/common"
)

type TransactionKind string

const (
	Deposit            TransactionKind = "deposit"
	DepositExchange    TransactionKind = "deposit-exchange"
	Withdrawal         TransactionKind = "withdrawal"
	WithdrawalExchange TransactionKind = "withdrawal-exchange"
)

type TransactionStatus string

const (
	Completed                    TransactionStatus = "completed"
	PendingExternal              TransactionStatus = "pending_external"
	PendingAnchor                TransactionStatus = "pending_anchor"
	OnHold                       TransactionStatus = "on_hold"
	PendingStellar               TransactionStatus = "pending_stellar"
	PendingTrust                 TransactionStatus = "pending_trust"
	PendingUser                  TransactionStatus = "pending_user"
	PendingUserTransferStart     TransactionStatus = "pending_user_transfer_start"
	PendingUserTransferComplete  TransactionStatus = "pending_user_transfer_complete"
	PendingCustomerInfoUpdate    TransactionStatus = "pending_customer_info_update"
	PendingTransactionInfoUpdate TransactionStatus = "pending_transaction_info_update"
	Incomplete                   TransactionStatus = "incomplete"
	NoMarket                     TransactionStatus = "no_market"
	TooSmall                     TransactionStatus = "too_small"
	TooLarge                     TransactionStatus = "too_large"
	Error                        TransactionStatus = "error"
	Refunded                     TransactionStatus = "refunded"
)

type RefundPayment struct {
	ID     string `json:"id"`
	IDType string `json:"id_type"`
	Amount string `json:"amount"`
	Fee    string `json:"fee"`
}

type Refund struct {
	AmountRefunded string          `json:"amount_refunded"`
	AmountFee      string          `json:"amount_fee"`
	Payments       []RefundPayment `json:"payments"`
}

type FeeDetailsDetail struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Amount      string `json:"amount"`
}

type FeeDetails struct {
	Total   string             `json:"total"`
	Asset   string             `json:"asset"`
	Details []FeeDetailsDetail `json:"details,omitempty"`
}

type RequiredInfoUpdates struct {
	Transaction map[string]any `json:"transaction"`
}

type Sep6Transaction struct {
	anchorutils.BaseTransaction

	// SEP-6 specific timestamp fields (time.Time)
	StartedAt            time.Time `json:"started_at,omitzero"`
	UpdatedAt            time.Time `json:"updated_at,omitzero"`
	CompletedAt          time.Time `json:"completed_at,omitzero"`
	UserActionRequiredBy time.Time `json:"user_action_required_by,omitzero"`

	// SEP-6 specific refunds type
	Refunds *Refund `json:"refunds,omitempty"`

	// SEP-6 only fields
	ExternalExtra       string               `json:"external_extra,omitempty"`
	ExternalExtraText   string               `json:"external_extra_text,omitempty"`
	RequiredInfoMessage string               `json:"required_info_message,omitempty"`
	RequiredInfoUpdates *RequiredInfoUpdates `json:"required_info_updates,omitempty"`
}

type TransactionListRequest struct {
	AssetCode   string    `url:"asset_code"`
	Account     string    `url:"account"`
	NoOlderThan time.Time `url:"no_older_than,omitempty"`
	Limit       int       `url:"limit,omitempty"`
	Kind        []string  `url:"kind,omitempty"`
	PagingID    string    `url:"paging_id,omitempty"`
	Lang        string    `url:"lang,omitempty"`
}

type TransactionListResponse struct {
	Transactions []Sep6Transaction `json:"transactions"`
}

type TransactionRequest struct {
	ID           string `url:"id,omitempty"`
	StellarTxID  string `url:"stellar_transaction_id,omitempty"`
	ExternalTxID string `url:"external_transaction_id,omitempty"`
	Lang         string `url:"lang,omitempty"`
}

type TransactionResponse struct {
	Transaction Sep6Transaction `json:"transaction"`
}

type TransactionUpdateRequest struct {
	Transaction map[string]any `json:"transaction"`
}
