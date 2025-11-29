package sep6

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	anchorutils "github.com/tryoutbounder/anchor-go/common"
)

type Sep6Client struct {
	baseURL    string
	httpClient *http.Client
	AnchorInfo Info
}

func NewSep6Client(baseURL string, httpClient *http.Client) (*Sep6Client, error) {
	client := &Sep6Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}

	info, err := client.Info("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get anchor info: %w", err)
	}
	client.AnchorInfo = *info

	return client, nil
}

func (c *Sep6Client) Info(lang string, bearer string) (*Info, error) {
	url := fmt.Sprintf("%s/info", c.baseURL)
	params := make(map[string][]string)
	if lang != "" {
		params["lang"] = []string{lang}
	}
	return anchorutils.AnchorRequest[any, Info](
		c.httpClient,
		"GET",
		url,
		nil,
		params,
		bearer,
	)
}

func (c *Sep6Client) Deposit(request *DepositRequest, bearer string) (*DepositResponse, error) {
	url := fmt.Sprintf("%s/deposit", c.baseURL)
	params, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	return anchorutils.AnchorRequest[any, DepositResponse](
		c.httpClient,
		"GET",
		url,
		nil,
		params,
		bearer,
	)
}

func (c *Sep6Client) DepositExchange(request *DepositExchangeRequest, bearer string) (*DepositExchangeResponse, error) {
	url := fmt.Sprintf("%s/deposit-exchange", c.baseURL)
	params, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	return anchorutils.AnchorRequest[any, DepositExchangeResponse](
		c.httpClient,
		"GET",
		url,
		nil,
		params,
		bearer,
	)
}

func (c *Sep6Client) Withdraw(request *WithdrawRequest, bearer string) (*WithdrawResponse, error) {
	url := fmt.Sprintf("%s/withdraw", c.baseURL)
	params, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	return anchorutils.AnchorRequest[any, WithdrawResponse](
		c.httpClient,
		"GET",
		url,
		nil,
		params,
		bearer,
	)
}

func (c *Sep6Client) WithdrawExchange(request *WithdrawExchangeRequest, bearer string) (*WithdrawExchangeResponse, error) {
	url := fmt.Sprintf("%s/withdraw-exchange", c.baseURL)
	params, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	return anchorutils.AnchorRequest[any, WithdrawExchangeResponse](
		c.httpClient,
		"GET",
		url,
		nil,
		params,
		bearer,
	)
}

func (c *Sep6Client) Fee(request *FeeRequest, bearer string) (*Fee, error) {
	url := fmt.Sprintf("%s/fee", c.baseURL)
	params, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	return anchorutils.AnchorRequest[any, Fee](
		c.httpClient,
		"GET",
		url,
		nil,
		params,
		bearer,
	)
}

func (c *Sep6Client) Transactions(request *TransactionListRequest, bearer string) (*TransactionListResponse, error) {
	url := fmt.Sprintf("%s/transactions", c.baseURL)
	params, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	return anchorutils.AnchorRequest[any, TransactionListResponse](
		c.httpClient,
		"GET",
		url,
		nil,
		params,
		bearer,
	)
}

func (c *Sep6Client) Transaction(request *TransactionRequest, bearer string) (*Sep6Transaction, error) {
	url := fmt.Sprintf("%s/transaction", c.baseURL)
	params, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	resp, err := anchorutils.AnchorRequest[any, TransactionResponse](
		c.httpClient,
		"GET",
		url,
		nil,
		params,
		bearer,
	)
	if err != nil {
		return nil, err
	}

	return &resp.Transaction, nil

}

func (c *Sep6Client) UpdateTransaction(id string, request *TransactionUpdateRequest, bearer string) error {
	url := fmt.Sprintf("%s/transactions/%s", c.baseURL, id)
	_, err := anchorutils.AnchorRequest[TransactionUpdateRequest, any](
		c.httpClient,
		"PATCH",
		url,
		request,
		nil,
		bearer,
	)
	if err != nil {
		return err
	}
	return nil
}
