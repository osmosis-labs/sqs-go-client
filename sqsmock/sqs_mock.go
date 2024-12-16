package sqsmock

import (
	"context"

	sqsclient "github.com/osmosis-labs/sqs-go-client"
)

// SQSMock is a mock implementation of the sqsclient.SQSClient interface.
type SQSMock struct {
	GetPricesFunc         func(ctx context.Context, options ...sqsclient.TokenPricesOption) (map[string]map[string]string, error)
	GetRouteFunc          func(ctx context.Context, options ...sqsclient.RouterQuoteOption) (sqsclient.SQSQuoteResponse, error)
	GetTokensMetadataFunc func(ctx context.Context) (map[string]sqsclient.OsmosisTokenMetadata, error)
}

// GetPrices implements sqsclient.SQSClient.
func (s *SQSMock) GetPrices(ctx context.Context, options ...sqsclient.TokenPricesOption) (map[string]map[string]string, error) {
	if s.GetPricesFunc != nil {
		return s.GetPricesFunc(ctx, options...)
	}

	return nil, nil
}

// GetQuote implements sqsclient.SQSClient.
func (s *SQSMock) GetQuote(ctx context.Context, options ...sqsclient.RouterQuoteOption) (sqsclient.SQSQuoteResponse, error) {
	if s.GetRouteFunc != nil {
		return s.GetRouteFunc(ctx, options...)
	}

	return sqsclient.SQSQuoteResponse{}, nil
}

// GetTokensMetadata implements sqsclient.SQSClient.
func (s *SQSMock) GetTokensMetadata(ctx context.Context) (map[string]sqsclient.OsmosisTokenMetadata, error) {
	if s.GetTokensMetadataFunc != nil {
		return s.GetTokensMetadataFunc(ctx)
	}

	return nil, nil
}

var _ sqsclient.SQSClient = (*SQSMock)(nil)
