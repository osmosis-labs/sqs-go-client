package main

import (
	"context"
	"fmt"

	sqsclient "github.com/osmosis-labs/sqs-go-client"
)

const (
	uosmoDenom = "uosmo"
	uionDenom  = "uion"
)

// Run with 'go run "github.com/osmosis-labs/sqs-go-client/examples"
func main() {
	ctx := context.Background()

	// Initialize the SQS client
	// You can set the API key with sqsclient.WithAPIKey("your-api-key")
	// See other options in sqs_factory.go
	sqs, err := sqsclient.Initialize()
	if err != nil {
		panic(err)
	}

	// Get prices for uosmo and uion
	// See other options in tokens_prices_options.go
	prices, err := sqs.GetPrices(ctx, sqsclient.WithBaseDenoms([]string{uosmoDenom, uionDenom}))
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\nprices: %+v\n", prices)

	// Get route from uosmo to uion
	// See other options in router_quote_options.go
	route, err := sqs.GetQuote(ctx, sqsclient.WithIsSingleRoute(), sqsclient.WithOutGivenIn(1000000, uosmoDenom, uionDenom))
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\nroute: %+v\n", route)

	// Get tokens metadata
	// See other options in tokens_metadata_options.go
	tokensMetadata, err := sqs.GetTokensMetadata(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\ntokensMetadata:\n")

	for _, token := range tokensMetadata {
		fmt.Println(token)
	}
}
