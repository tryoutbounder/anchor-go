package sep24

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	anchorutils "github.com/tryoutbounder/anchor-go/common"
)

type Sep24Client struct {
	baseUrl    string
	httpClient *http.Client
	AnchorInfo Info
}

// NewSep24Client creates a new Sep24Client.
// It also fetches the anchor's info for future reference.
func NewSep24Client(baseUrl string, httpClient *http.Client) (*Sep24Client, error) {
	client := &Sep24Client{
		baseUrl:    baseUrl,
		httpClient: httpClient,
	}

	// Fetch anchor info for future reference
	info, err := client.Info("") // Auth is optional for /info, so an empty bearer is fine
	if err != nil {
		return nil, fmt.Errorf("failed to fetch SEP-24 info during client initialization: %w", err)
	}
	client.AnchorInfo = *info

	return client, nil
}

// GetInfo fetches information about the anchor's SEP-24 support.
// It corresponds to the GET /info endpoint.
// Auth is optional for this endpoint, but a bearer token can be provided.
func (c *Sep24Client) Info(bearer string) (*Info, error) {
	endpoint := fmt.Sprintf("%s/info", c.baseUrl)
	return anchorutils.AnchorRequest[any, Info](c.httpClient, "GET", endpoint, nil, nil, bearer)
}

// GetFee fetches the fee for a specific operation.
// It corresponds to the GET /fee endpoint.
func (c *Sep24Client) Fee(request *FeeRequest, bearer string) (*FeeResponse, error) {
	endpoint := fmt.Sprintf("%s/fee", c.baseUrl)

	queryParams, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	return anchorutils.AnchorRequest[any, FeeResponse](c.httpClient, "GET", endpoint, nil, queryParams, bearer)
}

// PostDeposit initiates a deposit transaction.
// It corresponds to the POST /transactions/deposit/interactive endpoint.
func (c *Sep24Client) Deposit(request *DepositRequest, bearer string) (*DepositResponse, error) {
	endpoint := fmt.Sprintf("%s/transactions/deposit/interactive", c.baseUrl)
	return anchorutils.AnchorRequest[DepositRequest, DepositResponse](c.httpClient, "POST", endpoint, request, nil, bearer)
}

// PostWithdraw initiates a withdrawal transaction.
// It corresponds to the POST /transactions/withdraw/interactive endpoint.
func (c *Sep24Client) Withdraw(request *WithdrawRequest, bearer string) (*WithdrawResponse, error) {
	endpoint := fmt.Sprintf("%s/transactions/withdraw/interactive", c.baseUrl)
	return anchorutils.AnchorRequest[WithdrawRequest, WithdrawResponse](c.httpClient, "POST", endpoint, request, nil, bearer)
}

// GetTransactions fetches a list of historical transactions.
// It corresponds to the GET /transactions endpoint.
func (c *Sep24Client) Transactions(request *TransactionsRequest, bearer string) (*TransactionsResponse, error) {
	endpoint := fmt.Sprintf("%s/transactions", c.baseUrl)

	queryParams, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	return anchorutils.AnchorRequest[any, TransactionsResponse](c.httpClient, "GET", endpoint, nil, queryParams, bearer)
}

// GetTransaction fetches a single transaction by its properties.
// It corresponds to the GET /transaction endpoint.
func (c *Sep24Client) Transaction(request *TransactionRequest, bearer string) (*Sep24Transaction, error) {
	endpoint := fmt.Sprintf("%s/transaction", c.baseUrl)

	queryParams, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	resp, err := anchorutils.AnchorRequest[any, TransactionResponse](c.httpClient, "GET", endpoint, nil, queryParams, bearer)
	if err != nil {
		return nil, err
	}

	return &resp.Transaction, err
}
