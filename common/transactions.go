package anchorutils

// type TransactionKind string

// type TransactionStatus string

// const (
// 	TransactionStatusCompleted                    TransactionStatus = "completed"
// 	TransactionStatusPendingExternal              TransactionStatus = "pending_external"
// 	TransactionStatusPendingAnchor                TransactionStatus = "pending_anchor"
// 	TransactionStatusOnHold                       TransactionStatus = "on_hold"
// 	TransactionStatusPendingStellar               TransactionStatus = "pending_stellar"
// 	TransactionStatusPendingTrust                 TransactionStatus = "pending_trust"
// 	TransactionStatusPendingUser                  TransactionStatus = "pending_user"
// 	TransactionStatusPendingUserTransferStart     TransactionStatus = "pending_user_transfer_start"
// 	TransactionStatusPendingUserTransferComplete  TransactionStatus = "pending_user_transfer_complete"
// 	TransactionStatusPendingCustomerInfoUpdate    TransactionStatus = "pending_customer_info_update"
// 	TransactionStatusPendingTransactionInfoUpdate TransactionStatus = "pending_transaction_info_update"
// 	TransactionStatusIncomplete                   TransactionStatus = "incomplete"
// 	TransactionStatusNoMarket                     TransactionStatus = "no_market"
// 	TransactionStatusTooSmall                     TransactionStatus = "too_small"
// 	TransactionStatusTooLarge                     TransactionStatus = "too_large"
// 	TransactionStatusError                        TransactionStatus = "error"
// 	TransactionStatusRefunded                     TransactionStatus = "refunded"
// 	TransactionStatusExpired                      TransactionStatus = "expired"
// )

type BaseTransaction struct {
	ID                    string `json:"id"`
	Kind                  string `json:"kind"`
	Status                string `json:"status"`
	StatusEta             int64  `json:"status_eta,omitempty"`
	MoreInfoURL           string `json:"more_info_url,omitempty"`
	AmountIn              string `json:"amount_in,omitempty"`
	AmountInAsset         string `json:"amount_in_asset,omitempty"`
	AmountOut             string `json:"amount_out,omitempty"`
	AmountOutAsset        string `json:"amount_out_asset,omitempty"`
	AmountFee             string `json:"amount_fee,omitempty"`
	AmountFeeAsset        string `json:"amount_fee_asset,omitempty"`
	QuoteID               string `json:"quote_id,omitempty"`
	From                  string `json:"from,omitempty"`
	To                    string `json:"to,omitempty"`
	DepositMemo           string `json:"deposit_memo,omitempty"`
	DepositMemoType       string `json:"deposit_memo_type,omitempty"`
	WithdrawAnchorAccount string `json:"withdraw_anchor_account,omitempty"`
	WithdrawMemo          string `json:"withdraw_memo,omitempty"`
	WithdrawMemoType      string `json:"withdraw_memo_type,omitempty"`
	StellarTxID           string `json:"stellar_transaction_id,omitempty"`
	ExternalTxID          string `json:"external_transaction_id,omitempty"`
	Message               string `json:"message,omitempty"`
	Refunded              bool   `json:"refunded,omitempty"`
	ClaimableBalanceID    string `json:"claimable_balance_id,omitempty"`
}
