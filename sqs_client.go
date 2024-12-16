package sqsclient

import (
	"context"
	"fmt"

	"github.com/osmosis-labs/osmoutil-go/httputil"
)

// SQSClient is the interface for the Osmosis Sidecar Query Server (SQSClient) client.
type SQSClient interface {
	GetPrices(ctx context.Context, options ...TokenPricesOption) (map[string]map[string]string, error)
	GetTokensMetadata(ctx context.Context) (map[string]OsmosisTokenMetadata, error)
	GetQuote(ctx context.Context, options ...RouterQuoteOption) (SQSQuoteResponse, error)
}

type sqs struct {
	url          string
	apiKeyHeader map[string]string
}

// NewClient creates a new OsmosisSQS client.
func NewClient(url string) *sqs {
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

// GetPrices implements SQSClient
func (s *sqs) GetPrices(ctx context.Context, options ...TokenPricesOption) (map[string]map[string]string, error) {
	// Apply the options
	opts := TokenPricesOptions{}
	for _, option := range options {
		option(&opts)
	}

	var response map[string]map[string]string
	if err := s.httpGetWithOptions(ctx, "tokens/prices", &response, &opts); err != nil {
		return nil, fmt.Errorf("error getting base/USDC price: %v", err)
	}

	return response, nil
}

// GetTokensMetadata implements SQSClient
func (o *sqs) GetTokensMetadata(ctx context.Context) (map[string]OsmosisTokenMetadata, error) {

	tokenMetadataURL := fmt.Sprintf("%s/tokens/metadata", o.url)
	var response map[string]OsmosisTokenMetadata
	if err := httputil.RunGet(ctx, tokenMetadataURL, o.apiKeyHeader, &response); err != nil {
		return nil, fmt.Errorf("error getting token metadata: %v", err)
	}

	return response, nil
}

// GetQuote implements SQS
func (o *sqs) GetQuote(ctx context.Context, options ...RouterQuoteOption) (SQSQuoteResponse, error) {
	opts := RouterQuoteOptions{}
	for _, option := range options {
		option(&opts)
	}

	var quoteResponse SQSQuoteResponse
	if err := o.httpGetWithOptions(ctx, "router/quote", &quoteResponse, &opts); err != nil {
		return SQSQuoteResponse{}, err
	}

	return quoteResponse, nil
}

// httpGetWithOptions is a helper function to make an HTTP GET request with options.
// It validates the options, retrieves the query params, and makes the request, parsing the response
// into the given response paramter.
func (o *sqs) httpGetWithOptions(ctx context.Context, endpoint string, response interface{}, options Options) error {
	// Validate the options
	if err := options.Validate(); err != nil {
		return err
	}

	// Create the query params
	queryParams := options.CreateQueryParams()

	url := fmt.Sprintf("%s/%s?%s", o.url, endpoint, queryParams.Encode())

	if err := httputil.RunGet(ctx, url, o.apiKeyHeader, response); err != nil {
		return err
	}

	return nil
}
