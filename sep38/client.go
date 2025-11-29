package sep38

import (
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	anchorutils "github.com/tryoutbounder/anchor-go/common"
)

type Sep38Client struct {
	baseURL      string
	http         *http.Client
	ExchangeInfo Info
}

func NewSep38Client(baseURL string, httpClient *http.Client) (*Sep38Client, error) {
	client := &Sep38Client{
		baseURL: baseURL,
		http:    httpClient,
	}

	info, err := client.Info("")
	if err != nil {
		return nil, err
	}
	client.ExchangeInfo = *info

	return client, nil

}

func (c *Sep38Client) Info(bearer string) (*Info, error) {
	return anchorutils.AnchorRequest[any, Info](
		c.http,
		http.MethodGet,
		c.baseURL+"/info",
		nil,
		nil,
		bearer,
	)
}

func (c *Sep38Client) Prices(request GetPricesRequest, bearer string) (*GetPricesResponse, error) {

	params, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	return anchorutils.AnchorRequest[any, GetPricesResponse](
		c.http,
		http.MethodGet,
		c.baseURL+"/prices",
		nil,
		params,
		bearer,
	)
}

func (c *Sep38Client) Price(request GetPriceRequest, bearer string) (*GetPriceResponse, error) {
	params, err := query.Values(request)
	if err != nil {
		return nil, err
	}
	return anchorutils.AnchorRequest[any, GetPriceResponse](
		c.http,
		http.MethodGet,
		c.baseURL+"/price",
		nil,
		params,
		bearer,
	)
}

func (c *Sep38Client) PostQuote(request PostQuoteRequest, bearer string) (*QuoteResponse, error) {
	return anchorutils.AnchorRequest[PostQuoteRequest, QuoteResponse](
		c.http,
		http.MethodPost,
		c.baseURL+"/quote",
		&request,
		nil,
		bearer,
	)
}

func (c *Sep38Client) GetQuote(id string, bearer string) (*QuoteResponse, error) {
	params := url.Values{}
	if id != "" {
		params.Set("id", id)
	}

	return anchorutils.AnchorRequest[any, QuoteResponse](
		c.http,
		http.MethodGet,
		c.baseURL+"/quote/"+id,
		nil,
		params,
		bearer,
	)
}
