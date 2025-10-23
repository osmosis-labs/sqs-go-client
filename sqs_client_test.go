package sqsclient_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	sqsclient "github.com/osmosis-labs/sqs-go-client"
)

const (
	uosmoDenom = "uosmo"
	uionDenom  = "uion"
	atomDenom  = "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"
	usdcDenom  = "ibc/498A0751C798A0D9A389AA3691123DADA57DAA4FE165D5C75894505B876BA6E4"
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

func TestGetExactInRoute(t *testing.T) {
	t.Skip("skipping integration test")

	ctx := context.Background()

	sqs, err := sqsclient.Initialize(sqsclient.WithCustomURL("https://sqs.osmosis.zone"))
	require.NoError(t, err)

	route, err := sqs.GetQuote(ctx, sqsclient.WithOutGivenIn(1000000, uosmoDenom, uionDenom))
	require.NoError(t, err)

	require.Equal(t, route.AmountIn.Denom, uosmoDenom)
	require.Equal(t, route.AmountIn.Amount, "1000000")
	require.Equal(t, route.AmountOut.Denom, uionDenom)
	parsedAmount, err := strconv.ParseInt(route.AmountOut.Amount, 10, 64)
	require.NoError(t, err)
	require.Greater(t, parsedAmount, int64(0))

	t.Logf("route: %+v", route)
}

func TestGetExactOutRoute(t *testing.T) {
	t.Skip("skipping integration test")

	ctx := context.Background()

	sqs, err := sqsclient.Initialize(sqsclient.WithCustomURL("https://sqs.osmosis.zone"))
	require.NoError(t, err)

	route, err := sqs.GetQuote(ctx, sqsclient.WithInGivenOut(1000000, uionDenom, uosmoDenom))
	require.NoError(t, err)
	require.Equal(t, route.AmountOut.Denom, uionDenom)
	require.Equal(t, route.AmountOut.Amount, "1000000")
	require.Equal(t, route.AmountIn.Denom, uosmoDenom)
	parsedAmount, err := strconv.ParseInt(route.AmountIn.Amount, 10, 64)
	require.NoError(t, err)
	require.Greater(t, parsedAmount, int64(0))

	t.Logf("route: %+v", route)
}

func TestGetRoute_CustomDirectQuote(t *testing.T) {
	t.Skip("skipping integration test")

	ctx := context.Background()

	sqs, err := sqsclient.Initialize(sqsclient.WithCustomURL("https://sqs.osmosis.zone"))
	require.NoError(t, err)

	route, err := sqs.GetQuote(ctx, sqsclient.WithOutGivenInCustom(1000000, usdcDenom, []string{uosmoDenom, atomDenom}, []uint64{1464, 1135}))
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

func TestGetCustomDirectQuote(t *testing.T) {
	t.Skip("skipping integration test")

	ctx := context.Background()

	sqs, err := sqsclient.Initialize()
	require.NoError(t, err)

	prices, err := sqs.GetPrices(ctx, sqsclient.WithBaseDenom(uosmoDenom))
	require.NoError(t, err)

	t.Logf("prices: %+v", prices)
}
