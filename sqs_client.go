package sqs

import (
	"context"
	"fmt"

	"github.com/osmosis-labs/osmoutil-go/httputil"
)

// SQS is the interface for the Osmosis Sidecar Query Server (SQS) client.
type SQS interface {
	GetPrices(ctx context.Context, denoms []string) (map[string]map[string]string, error)
	GetTokensMetadata(ctx context.Context) (map[string]OsmosisTokenMetadata, error)
	GetRoute(ctx context.Context, options ...RouterQuoteOption) (SQSQuoteResponse, error)
}

type sqs struct {
	url          string
	apiKeyHeader map[string]string
}

// NewOsmosisSQS creates a new OsmosisSQS client.
func NewOsmosisSQS(url string) *sqs {

	return &sqs{
		url:          url,
		apiKeyHeader: nil,
	}
}

// WithAPIKey is a helper function to set the API key header for the sqs client.
func WithAPIKey(apiKey string, sqs *sqs) *sqs {
	sqs.apiKeyHeader = map[string]string{
		"x-api-key-header": apiKey,
	}
	return sqs
}

// GetPrices implements SQS
func (s *sqs) GetPrices(ctx context.Context, denoms []string) (map[string]map[string]string, error) {

	priceURL := fmt.Sprintf("%s/tokens/prices?base=%s,%s", s.url, denoms[0], denoms[1])
	var response map[string]map[string]string
	if err := httputil.RunGet(ctx, priceURL, s.apiKeyHeader, &response); err != nil {
		return nil, fmt.Errorf("error getting base/USDC price: %v", err)
	}

	return response, nil
}

// GetTokensMetadata implements SQS
func (o *sqs) GetTokensMetadata(ctx context.Context) (map[string]OsmosisTokenMetadata, error) {

	tokenMetadataURL := fmt.Sprintf("%s/tokens/metadata", o.url)
	var response map[string]OsmosisTokenMetadata
	if err := httputil.RunGet(ctx, tokenMetadataURL, o.apiKeyHeader, &response); err != nil {
		return nil, fmt.Errorf("error getting token metadata: %v", err)
	}

	return response, nil
}

// GetRoute implements SQS
func (o *sqs) GetRoute(ctx context.Context, options ...RouterQuoteOption) (SQSQuoteResponse, error) {
	opts := RouterQuoteOptions{}
	for _, option := range options {
		option(&opts)
	}

	// Validate the options
	if err := opts.Validate(); err != nil {
		return SQSQuoteResponse{}, err
	}

	// Create the query params
	queryParams := opts.CreateQueryParams()

	url := fmt.Sprintf("%s/router/quote?%s", o.url, queryParams.Encode())

	var quoteResponse SQSQuoteResponse
	if err := httputil.RunGet(ctx, url, o.apiKeyHeader, &quoteResponse); err != nil {
		return SQSQuoteResponse{}, err
	}

	return quoteResponse, nil
}
