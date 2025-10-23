package sqsclient

import (
	"context"
	"fmt"

	"github.com/osmosis-labs/osmoutil-go/httputil"
)

type sqsExactInQuoteResponse struct {
	AmountIn                Coin      `json:"amount_in"`
	AmountOut               string    `json:"amount_out"`
	Route                   []Route   `json:"route"`
	EffectiveFee            string    `json:"effective_fee"`
	PriceImpact             string    `json:"price_impact"`
	InBaseOutQuoteSpotPrice string    `json:"in_base_out_quote_spot_price"`
	PriceInfo               PriceInfo `json:"price_info"`
}

type sqsExactOutQuoteResponse struct {
	AmountIn                string    `json:"amount_in"`
	AmountOut               Coin      `json:"amount_out"`
	Route                   []Route   `json:"route"`
	EffectiveFee            string    `json:"effective_fee"`
	PriceImpact             string    `json:"price_impact"`
	InBaseOutQuoteSpotPrice string    `json:"in_base_out_quote_spot_price"`
	PriceInfo               PriceInfo `json:"price_info"`
}

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
		"x-api-key": apiKey,
	}
	return sqs
}

// GetPrices implements SQSClient
func (o *sqs) GetPrices(ctx context.Context, options ...TokenPricesOption) (map[string]map[string]string, error) {
	// Apply the options
	opts := TokenPricesOptions{}
	for _, option := range options {
		option(&opts)
	}

	var response map[string]map[string]string
	if err := o.httpGetWithOptions(ctx, "tokens/prices", &response, &opts); err != nil {
		return nil, fmt.Errorf("error getting base/USDC price: %v", err)
	}

	return response, nil
}

// GetTokensMetadata implements SQSClient
func (o *sqs) GetTokensMetadata(ctx context.Context) (map[string]OsmosisTokenMetadata, error) {

	tokenMetadataURL := fmt.Sprintf("%s/tokens/metadata", o.url)
	var response map[string]OsmosisTokenMetadata
	if _, err := httputil.Get(ctx, tokenMetadataURL, o.apiKeyHeader, &response); err != nil {
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

	var urlExtension string
	if len(opts.PoolIDs) == 0 {
		urlExtension = "router/quote"
	} else {
		urlExtension = "router/custom-direct-quote"
	}

	if opts.IsOutGivenIn() {
		var exactInResponse sqsExactInQuoteResponse
		if err := o.httpGetWithOptions(ctx, urlExtension, &exactInResponse, &opts); err != nil {
			return SQSQuoteResponse{}, err
		}
		return convertExactInResponseToQuoteResponse(exactInResponse, opts), nil
	} else {
		var exactOutResponse sqsExactOutQuoteResponse
		if err := o.httpGetWithOptions(ctx, urlExtension, &exactOutResponse, &opts); err != nil {
			return SQSQuoteResponse{}, err
		}
		return convertExactOutResponseToQuoteResponse(exactOutResponse, opts), nil
	}

}

// convertExactInResponseToQuoteResponse converts an OutGivenIn response to the standard SQSQuoteResponse format
func convertExactInResponseToQuoteResponse(response sqsExactInQuoteResponse, opts RouterQuoteOptions) SQSQuoteResponse {
	var outputDenom string
	if len(opts.TokenOutDenom) > 0 {
		outputDenom = opts.TokenOutDenom[0]
	}

	return SQSQuoteResponse{
		AmountIn:                response.AmountIn,
		AmountOut:               Coin{Denom: outputDenom, Amount: response.AmountOut},
		Route:                   response.Route,
		EffectiveFee:            response.EffectiveFee,
		PriceImpact:             response.PriceImpact,
		InBaseOutQuoteSpotPrice: response.InBaseOutQuoteSpotPrice,
		PriceInfo:               response.PriceInfo,
	}
}

// convertExactOutResponseToQuoteResponse converts an InGivenOut response to the standard SQSQuoteResponse format
func convertExactOutResponseToQuoteResponse(response sqsExactOutQuoteResponse, opts RouterQuoteOptions) SQSQuoteResponse {
	var inputDenom string
	if len(opts.TokenInDenom) > 0 {
		inputDenom = opts.TokenInDenom[0]
	}

	return SQSQuoteResponse{
		AmountIn:                Coin{Denom: inputDenom, Amount: response.AmountIn},
		AmountOut:               response.AmountOut,
		Route:                   response.Route,
		EffectiveFee:            response.EffectiveFee,
		PriceImpact:             response.PriceImpact,
		InBaseOutQuoteSpotPrice: response.InBaseOutQuoteSpotPrice,
		PriceInfo:               response.PriceInfo,
	}
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

	if _, err := httputil.Get(ctx, url, o.apiKeyHeader, response); err != nil {
		return err
	}

	return nil
}
