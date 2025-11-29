package sep10

type ChallengeRequest struct {
	Account      string `url:"account"`                 // The Client Account to authenticate
	Memo         string `url:"memo,omitempty"`          // Optional memo for shared accounts
	HomeDomain   string `url:"home_domain,omitempty"`   // Optional home domain for servers supporting multiple
	ClientDomain string `url:"client_domain,omitempty"` // Optional client domain for additional verification
}

type ChallengeResponse struct {
	Transaction       string `url:"transaction"`        // XDR-encoded Stellar challenge transaction
	NetworkPassphrase string `url:"network_passphrase"` // Stellar network passphrase
}

type TokenRequest struct {
	Transaction string `json:"transaction"` // The signed challenge transaction XDR
}

type TokenResponse struct {
	Token string `json:"token"` // JWT session token
}

type TokenClaims struct {
	Issuer       string `json:"iss"`                     // The principal that issued the token
	Subject      string `json:"sub"`                     // The principal that is the subject of the JWT
	IssuedAt     int64  `json:"iat"`                     // The time at which the JWT was issued
	ExpiresAt    int64  `json:"exp"`                     // The expiration time
	ClientDomain string `json:"client_domain,omitempty"` // Optional client domain if included in challenge
}
