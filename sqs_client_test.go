package sqsclient_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	sqsclient "github.com/osmosis-labs/sqs-go-client"
)

const (
	uosmoDenom = "uosmo"
	uionDenom  = "uion"
)

func TestGetTokensMetadata(t *testing.T) {
	t.Skip("skipping integration test")

	ctx := context.Background()

	sqs, err := sqsclient.Initialize()
	require.NoError(t, err)

	metadata, err := sqs.GetTokensMetadata(ctx)
	require.NoError(t, err)

	t.Logf("tokens metadata: %+v", metadata)
}

func TestGetRoute(t *testing.T) {
	t.Skip("skipping integration test")

	ctx := context.Background()

	sqs, err := sqsclient.Initialize(sqsclient.WithCustomURL("https://sqs.osmosis.zone"))
	require.NoError(t, err)

	route, err := sqs.GetRoute(ctx, sqsclient.WithOutGivenIn(1000000, uosmoDenom, uionDenom))
	require.NoError(t, err)

	t.Logf("route: %+v", route)
}

func TestGetPrice(t *testing.T) {
	t.Skip("skipping integration test")

	ctx := context.Background()

	sqs, err := sqsclient.Initialize()
	require.NoError(t, err)

	prices, err := sqs.GetPrices(ctx, sqsclient.WithBaseDenom(uosmoDenom))
	require.NoError(t, err)

	t.Logf("prices: %+v", prices)
}
