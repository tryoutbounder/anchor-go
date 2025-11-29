package sep12

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
	anchorutils "github.com/tryoutbounder/anchor-go/common"
)

type Sep12Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewSep12Client creates a new SEP-12 client
func NewSep12Client(baseURL string, httpClient *http.Client) *Sep12Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Sep12Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

// GetCustomer implements GET /customer endpoint
func (c *Sep12Client) GetCustomer(request *CustomerGetRequest, bearer string) (*CustomerGetResponse, error) {
	url := fmt.Sprintf("%s/customer", c.baseURL)

	// // Convert request to URL parameters
	urlParams, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	return anchorutils.AnchorRequest[CustomerGetRequest, CustomerGetResponse](
		c.httpClient,
		"GET",
		url,
		nil,
		urlParams,
		bearer,
	)
}

// PutCustomer implements PUT /customer endpoint
func (c *Sep12Client) PutCustomer(request *CustomerPutRequest, bearer string) (*CustomerPutResponse, error) {
	url := fmt.Sprintf("%s/customer", c.baseURL)

	return anchorutils.AnchorRequest[CustomerPutRequest, CustomerPutResponse](
		c.httpClient,
		"PUT",
		url,
		request,
		nil,
		bearer,
	)
}

// PutCustomerCallback implements PUT /customer/callback endpoint
func (c *Sep12Client) PutCustomerCallback(request *CustomerCallbackPutRequest, bearer string) error {
	url := fmt.Sprintf("%s/customer/callback", c.baseURL)

	_, err := anchorutils.AnchorRequest[CustomerCallbackPutRequest, any](
		c.httpClient,
		"PUT",
		url,
		request,
		nil,
		bearer,
	)
	return err
}

// DeleteCustomer implements DELETE /customer/[account] endpoint
func (c *Sep12Client) DeleteCustomer(account string, request *CustomerDeleteRequest, bearer string) error {
	url := fmt.Sprintf("%s/customer/%s", c.baseURL, url.PathEscape(account))

	urlParams, validateErr := query.Values(request)

	if validateErr != nil {
		return validateErr
	}

	_, err := anchorutils.AnchorRequest[any, any](
		c.httpClient,
		"DELETE",
		url,
		nil,
		urlParams,
		bearer,
	)
	return err
}

// PostCustomerFiles implements POST /customer/files endpoint
func (c *Sep12Client) PostCustomerFiles(file io.Reader, filename string, bearer string) (*CustomerFilesPostResponse, error) {
	url := fmt.Sprintf("%s/customer/files", c.baseURL)

	return anchorutils.AnchorFileUpload[CustomerFilesPostResponse](
		c.httpClient,
		url,
		file,
		filename,
		bearer,
	)
}

// GetCustomerFiles implements GET /customer/files endpoint
func (c *Sep12Client) GetCustomerFiles(request *CustomerFilesGetRequest, bearer string) (*CustomerFilesGetResponse, error) {
	url := fmt.Sprintf("%s/customer/files", c.baseURL)

	urlParams, err := query.Values(request)

	if err != nil {
		return nil, err
	}

	return anchorutils.AnchorRequest[CustomerFilesGetRequest, CustomerFilesGetResponse](
		c.httpClient,
		"GET",
		url,
		nil,
		urlParams,
		bearer,
	)
}
