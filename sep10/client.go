package sep10

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/txnbuild"
	"github.com/stellar/go/xdr"
	anchorutils "github.com/tryoutbounder/anchor-go/common"
)

type Sep10Client struct {
	httpClient        *http.Client
	serverURL         string
	networkPassphrase string
}

func NewSep10Client(
	serverURL string,
	networkPassphrase string,
	httpClient *http.Client,
) (*Sep10Client, error) {

	return &Sep10Client{
		httpClient:        httpClient,
		serverURL:         serverURL,
		networkPassphrase: networkPassphrase,
	}, nil
}

func (c *Sep10Client) GetChallenge(req ChallengeRequest) (*ChallengeResponse, error) {
	params, err := query.Values(req)

	if err != nil {
		return nil, err
	}

	return anchorutils.AnchorRequest[ChallengeRequest, ChallengeResponse](
		c.httpClient,
		http.MethodGet,
		c.serverURL,
		nil,
		params,
		"",
	)
}

func (c *Sep10Client) GetToken(transaction string) (*TokenResponse, error) {
	req := &TokenRequest{
		Transaction: transaction,
	}

	return anchorutils.AnchorRequest[TokenRequest, TokenResponse](
		c.httpClient,
		http.MethodPost,
		c.serverURL,
		req,
		nil,
		"",
	)
}

func (c *Sep10Client) SignChallenge(
	challengeTx string,
	networkPassphrase string,
	keypair *keypair.Full,
) (string, error) {
	// Decode the challenge transaction
	var envelope xdr.TransactionEnvelope
	err := xdr.SafeUnmarshalBase64(challengeTx, &envelope)
	if err != nil {
		return "", fmt.Errorf("failed to decode challenge transaction: %w", err)
	}

	// Convert to transaction object
	tx, err := txnbuild.TransactionFromXDR(challengeTx)
	if err != nil {
		return "", fmt.Errorf("failed to parse challenge transaction: %w", err)
	}

	// Sign the transaction
	genTx, ok := tx.Transaction()
	if !ok {
		return "", fmt.Errorf("failed to extract transaction from challenge")
	}

	if genTx.SequenceNumber() != 0 {
		return "", fmt.Errorf("Invalid sequence number. Must be 0")
	}

	signedTx, err := genTx.Sign(networkPassphrase, keypair)
	if err != nil {
		return "", fmt.Errorf("failed to sign challenge transaction: %w", err)
	}

	// Get the signed transaction XDR
	signedTxXDR, err := signedTx.Base64()
	if err != nil {
		return "", fmt.Errorf("failed to encode signed transaction: %w", err)
	}

	return signedTxXDR, nil
}

func (c *Sep10Client) Authenticate(
	keypair *keypair.Full,
	address,
	clientDomain,
	homeDomain string,
) (string, error) {

	resp, err := c.GetChallenge(ChallengeRequest{
		Account:      address,
		ClientDomain: clientDomain,
		HomeDomain:   homeDomain,
	})
	if err != nil {
		return "", fmt.Errorf("failed to get challenge: %w", err)
	}

	// Sign the challenge transaction
	signedTx, err := c.SignChallenge(resp.Transaction, c.networkPassphrase, keypair)
	if err != nil {
		return "", fmt.Errorf("failed to sign challenge: %w", err)
	}

	// Submit the signed transaction and get token
	tokenResp, err := c.GetToken(signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to get token: %w", err)
	}

	return tokenResp.Token, nil
}
func (c *Sep10Client) TokenValid(bearerToken string) bool {
	// Parse the JWT token without verification first to extract claims
	parts := strings.Split(bearerToken, ".")
	if len(parts) != 3 {
		return false
	}

	// Decode the payload (second part)
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}

	// Parse the claims
	var claims TokenClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return false
	}

	// Check if token is expired or expires within 5 minutes
	now := time.Now().Unix()
	expiresSoonOrExpired := claims.ExpiresAt < now+int64(5*time.Minute/time.Second)

	// A token is valid if it's not expired and does not expire within 5 minutes
	return !expiresSoonOrExpired
}
