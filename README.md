
# anchor-go

This library provides clients for various Stellar Ecosystem Proposals (SEPs).

## SEP-1: stellar.toml

SEP-1 defines a standard way for Stellar anchors to provide information about their services in a `stellar.toml` file. This file is the first point of contact for any Stellar client interacting with an anchor.

### Supported Functionality

- **Fetch `stellar.toml`:** Retrieve and parse a `stellar.toml` file from a given domain.

### Usage

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/your-username/anchor-go/sep1"
)

func main() {
	// You can use any http.Client
	httpClient := http.DefaultClient

	// The URL of the stellar.toml file
	url := "https://example.com/.well-known/stellar.toml"

	stellarToml, err := sep1.GetToml(url, httpClient)
	if err != nil {
		log.Fatalf("Error fetching stellar.toml: %v", err)
	}

	fmt.Printf("Successfully fetched stellar.toml from %s\n", stellarToml.General.NetworkPassphrase)
}
```

## SEP-6: Deposit and Withdrawal API

SEP-6 provides a standardized API for anchors to handle deposits and withdrawals of assets on the Stellar network.

### Supported Functionality

- **Get Anchor Info:** Fetch information about the anchor's SEP-6 support.
- **Deposit:** Request a deposit of an asset.
- **Withdraw:** Request a withdrawal of an asset.
- **Transactions:** Fetch the status of transactions.

### Usage

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/your-username/anchor-go/sep6"
)

func main() {
	// You can use any http.Client
	httpClient := http.DefaultClient

	// The base URL of the anchor's SEP-6 API
	baseURL := "https://anchor.example.com/sep6"

	sep6Client, err := sep6.NewSep6Client(baseURL, httpClient)
	if err != nil {
		log.Fatalf("Error creating SEP-6 client: %v", err)
	}

	// Example: Get anchor info
	info, err := sep6Client.Info("", "")
	if err != nil {
		log.Fatalf("Error getting anchor info: %v", err)
	}
	fmt.Printf("Anchor supports withdrawal of asset: %s\n", info.Withdraw["USD"].Asset)


	// Example: Deposit request
	depositRequest := &sep6.DepositRequest{
		AssetCode: "USD",
		Account:   "G...", // Your Stellar account ID
	}
	depositResponse, err := sep6Client.Deposit(depositRequest, "your_jwt_token")
	if err != nil {
		log.Fatalf("Error making deposit request: %v", err)
	}
	fmt.Printf("Deposit instructions: %s\n", depositResponse.How)


	// Example: Withdrawal request
	withdrawalRequest := &sep6.WithdrawRequest{
		AssetCode: "USD",
		Type:      "bank_account",
		Dest:      "...", // Your bank account number
	}
	withdrawalResponse, err := sep6Client.Withdraw(withdrawalRequest, "your_jwt_token")
	if err != nil {
		log.Fatalf("Error making withdrawal request: %v", err)
	}
	fmt.Printf("Withdrawal initiated with ID: %s\n", withdrawalResponse.ID)
}
```

## SEP-10: Stellar Authentication

SEP-10 defines a protocol for authenticating a Stellar account with an anchor. This is a necessary step before using other SEPs that require authentication, such as SEP-6 and SEP-24.

### Supported Functionality

- **Get Challenge:** Request a challenge transaction from the anchor.
- **Sign Challenge:** Sign the challenge transaction with the client's private key.
- **Get Token:** Submit the signed transaction to the anchor to get a JWT token.
- **Authenticate:** A convenience function that handles the entire authentication flow.
- **Validate Token:** Check if a JWT token is still valid.

### Usage

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stellar/go/keypair"
	"github.com/your-username/anchor-go/sep10"
)

func main() {
	// You can use any http.Client
	httpClient := http.DefaultClient

	// The URL of the anchor's SEP-10 authentication endpoint
	authURL := "https://anchor.example.com/auth"

	// The Stellar network passphrase
	networkPassphrase := "Test SDF Network ; September 2015"

	// The keypair of the client's Stellar account
	clientKeypair, err := keypair.Random()
	if err != nil {
		log.Fatalf("Error creating keypair: %v", err)
	}


	sep10Client, err := sep10.NewSep10Client(authURL, networkPassphrase, httpClient)
	if err != nil {
		log.Fatalf("Error creating SEP-10 client: %v", err)
	}

	// Authenticate with the anchor
	token, err := sep10Client.Authenticate(clientKeypair, clientKeypair.Address(), "example.com", "anchor.example.com")
	if err != nil {
		log.Fatalf("Error authenticating: %v", err)
	}

	fmt.Printf("Successfully authenticated. JWT Token: %s\n", token)

	// Check if the token is valid
	isValid := sep10Client.TokenValid(token)
	fmt.Printf("Token is valid: %v\n", isValid)
}
```

## SEP-12: KYC API

SEP-12 provides a standard way for anchors to request Know Your Customer (KYC) information from users.

### Supported Functionality

- **Get Customer:** Get the status of a customer's KYC information.
- **Put Customer:** Create or update a customer's KYC information.
- **Delete Customer:** Delete a customer's information.

### Usage

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/your-username/anchor-go/sep12"
)

func main() {
	// You can use any http.Client
	httpClient := http.DefaultClient

	// The base URL of the anchor's SEP-12 API
	baseURL := "https://anchor.example.com/sep12"

	sep12Client := sep12.NewSep12Client(baseURL, httpClient)

	// Example: Get customer info
	getCustomerRequest := &sep12.CustomerGetRequest{
		Account: "G...", // The user's Stellar account ID
	}
	customer, err := sep12Client.GetCustomer(getCustomerRequest, "your_jwt_token")
	if err != nil {
		log.Fatalf("Error getting customer info: %v", err)
	}
	fmt.Printf("Customer status: %s\n", customer.Status)


	// Example: Put customer info
	putCustomerRequest := &sep12.CustomerPutRequest{
		Account: "G...", // The user's Stellar account ID
		Sep9Fields: anchorutils.Sep9Fields {
			FirstName: "John",
			LastName:  "Doe",
		}
	}
	putResponse, err := sep12Client.PutCustomer(putCustomerRequest, "your_jwt_token")
	if err != nil {
		log.Fatalf("Error putting customer info: %v", err)
	}
	fmt.Printf("Customer created with ID: %s\n", putResponse.ID)
}
```

## SEP-24: Interactive Deposit and Withdrawal

SEP-24 provides a standard for anchors to offer interactive, off-chain deposit and withdrawal flows. This is typically used for assets that require user interaction, such as bank transfers.

### Supported Functionality

- **Get Anchor Info:** Fetch information about the anchor's SEP-24 support.
- **Deposit:** Initiate an interactive deposit.
- **Withdraw:** Initiate an interactive withdrawal.
- **Transactions:** Fetch the status of transactions.

### Usage

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/your-username/anchor-go/sep24"
)

func main() {
	// You can use any http.Client
	httpClient := http.DefaultClient

	// The base URL of the anchor's SEP-24 API
	baseURL := "https://anchor.example.com/sep24"

	sep24Client, err := sep24.NewSep24Client(baseURL, httpClient)
	if err != nil {
		log.Fatalf("Error creating SEP-24 client: %v", err)
	}

	// Example: Get anchor info
	info, err := sep24Client.Info("your_jwt_token")
	if err != nil {
		log.Fatalf("Error getting anchor info: %v", err)
	}
	fmt.Printf("Anchor supports deposit of asset: %s\n", info.Deposit["USD"].Asset)


	// Example: Deposit request
	depositRequest := &sep24.DepositRequest{
		AssetCode: "USD",
		Account:   "G...", // Your Stellar account ID
	}
	depositResponse, err := sep24Client.Deposit(depositRequest, "your_jwt_token")
	if err != nil {
		log.Fatalf("Error making deposit request: %v", err)
	}
	fmt.Printf("Interactive deposit URL: %s\n", depositResponse.Url)


	// Example: Withdrawal request
	withdrawalRequest := &sep24.WithdrawRequest{
		AssetCode: "USD",
		Account:   "G...", // Your Stellar account ID
	}
	withdrawalResponse, err := sep24Client.Withdraw(withdrawalRequest, "your_jwt_token")
	if err != nil {
		log.Fatalf("Error making withdrawal request: %v", err)
	}
	fmt.Printf("Interactive withdrawal URL: %s\n", withdrawalResponse.Url)
}
```

## SEP-38: Anchor RFQ API

SEP-38 provides a way for anchors to offer quotes for exchanging assets. This is useful for cross-asset payments where the user wants to know the exchange rate before initiating a transaction.

### Supported Functionality

- **Get Info:** Get information about the assets the anchor supports for quoting.
- **Get Prices:** Get indicative prices for a given asset pair.
- **Get Price:** Get a firm price for a given asset pair.
- **Post Quote:** Request a quote for a given asset pair.
- **Get Quote:** Get a previously requested quote.

### Usage

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/your-username/anchor-go/sep38"
)

func main() {
	// You can use any http.Client
	httpClient := http.DefaultClient

	// The base URL of the anchor's SEP-38 API
	baseURL := "https://anchor.example.com/sep38"

	sep38Client, err := sep38.NewSep38Client(baseURL, httpClient)
	if err != nil {
		log.Fatalf("Error creating SEP-38 client: %v", err)
	}

	// Example: Get anchor info
	info, err := sep38Client.Info("your_jwt_token")
	if err != nil {
		log.Fatalf("Error getting anchor info: %v", err)
	}
	fmt.Printf("Anchor supports quoting for asset: %s\n", info.Assets[0].Asset)


	// Example: Get price
	priceRequest := sep38.GetPriceRequest{
		SellAsset: "iso4217:USD",
		BuyAsset:  "stellar:USDC:G...",
		SellAmount: "100",
	}
	price, err := sep38Client.Price(priceRequest, "your_jwt_token")
	if err != nil {
		log.Fatalf("Error getting price: %v", err)
	}
	fmt.Printf("The price for 100 USD is %s USDC\n", price.BuyAmount)


	// Example: Post quote
	quoteRequest := sep38.PostQuoteRequest{
		SellAsset: "iso4217:USD",
		BuyAsset:  "stellar:USDC:G...",
		SellAmount: "100",
	}
	quote, err := sep38Client.PostQuote(quoteRequest, "your_jwt_token")
	if err != nil {
		log.Fatalf("Error posting quote: %v", err)
	}
	fmt.Printf("Quote ID: %s\n", quote.ID)
}
```
