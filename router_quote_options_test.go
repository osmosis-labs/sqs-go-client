package sqsclient_test

import (
	"testing"

	sqsclient "github.com/osmosis-labs/sqs-go-client"
	"github.com/stretchr/testify/require"
)

func TestWithOutGivenIn(t *testing.T) {

	opts := sqsclient.WithOutGivenIn(5000000, "uatom", "uosmo")

	options := &sqsclient.RouterQuoteOptions{}
	opts(options)

	require.Equal(t, options.TokenIn, "5000000uatom")
	require.Equal(t, options.TokenOutDenom, []string{"uosmo"})
}
