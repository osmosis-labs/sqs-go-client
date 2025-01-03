package sqsclient

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// RouterQuoteOptions is the options for the /router/quote endpoint.
type RouterQuoteOptions struct {
	// Out given in.
	// TokenIn is the token in and denom to swap from.
	// E.g. 10uosmo
	TokenIn string
	// TokenOutDenom is the denom to swap to.
	TokenOutDenom []string

	// In given out.
	// TokenOut is the token out and denom to swap to.
	// E.g. 10uatom
	TokenOut string
	// TokenInDenom is the denom to swap from.
	TokenInDenom []string

	// HumanDenoms is whether the input tokens are human readable denoms.
	HumanDenoms bool
	// IsSingleRoute is whether the quote is for a single route.
	// If true, split routes are not returned.
	// If false, split routes are attempted to be computed.
	IsSingleRoute bool

	// PoolIDs is the pool IDs to use for the quote.
	PoolIDs []string

	// AppendBaseFee is whether the base fee is appended to the quote.
	AppendBaseFee bool
}

// RouterQuoteOption is the type for the options for the /router/quote endpoint.
type RouterQuoteOption func(opts *RouterQuoteOptions)

// Validate validates the RouterQuoteOptions.
// It returns an error if the options are invalid.
func (o *RouterQuoteOptions) Validate() error {
	if o.TokenIn == "" && o.TokenOut == "" {
		return fmt.Errorf("token in or token out must be set")
	}
	if len(o.TokenInDenom) == 0 && len(o.TokenOutDenom) == 0 {
		return fmt.Errorf("token in denom or token out denom must be set")
	}

	if o.TokenIn != "" && o.TokenOut != "" {
		return fmt.Errorf("token in and token out cannot be set at the same time")
	}
	if len(o.TokenInDenom) > 0 && len(o.TokenOutDenom) > 0 {
		return fmt.Errorf("token in denom and token out denom cannot be set at the same time")
	}

	return nil
}

// IsOutGivenIn returns true if the quote is for an out given in swap.
func (o *RouterQuoteOptions) IsOutGivenIn() bool {
	return o.TokenIn != "" && len(o.TokenOutDenom) > 0
}

// CreateQueryParams creates the query parameters for the /router/quote endpoint.
func (o *RouterQuoteOptions) CreateQueryParams() url.Values {
	queryParams := url.Values{}
	queryParams.Add("humanDenoms", strconv.FormatBool(o.HumanDenoms))
	queryParams.Add("singleRoute", strconv.FormatBool(o.IsSingleRoute))

	if o.IsOutGivenIn() {
		queryParams.Add("tokenIn", o.TokenIn)
		queryParams.Add("tokenOutDenom", strings.Join(o.TokenOutDenom, ","))
	} else {
		queryParams.Add("tokenInDenom", strings.Join(o.TokenInDenom, ","))
		queryParams.Add("tokenOut", o.TokenOut)
	}

	if o.HumanDenoms {
		queryParams.Add("humanDenoms", "true")
	}

	if o.IsSingleRoute {
		queryParams.Add("singleRoute", "true")
	}

	if o.AppendBaseFee {
		queryParams.Add("appendBaseFee", "true")
	}

	if len(o.PoolIDs) > 0 {
		queryParams.Add("poolID", strings.Join(o.PoolIDs, ","))
	}

	return queryParams
}

// WithOutGivenIn sets the options for an out given in swap for the /router/quote endpoint.
func WithOutGivenIn[T any](inAmount T, tokenInDenom string, tokenOutDenom string) RouterQuoteOption {
	return func(opts *RouterQuoteOptions) {
		opts.TokenIn = fmt.Sprintf("%v%s", inAmount, tokenInDenom)
		opts.TokenOutDenom = []string{tokenOutDenom}
	}
}

// WithOutGivenInCustom sets the options for an out given in swap for the /router/custom-direct-quote endpoint.
func WithOutGivenInCustom[T any](inAmount T, tokenInDenom string, tokenOutDenom []string, poolIDs []uint64) RouterQuoteOption {
	return func(opts *RouterQuoteOptions) {
		opts.TokenIn = fmt.Sprintf("%v%s", inAmount, tokenInDenom)
		opts.TokenOutDenom = tokenOutDenom

		poolIDsStr := make([]string, len(poolIDs))
		for i, id := range poolIDs {
			poolIDsStr[i] = strconv.FormatUint(id, 10)
		}

		opts.PoolIDs = poolIDsStr
	}
}

// WithInGivenOut sets the options for an in given out swap for the /router/quote endpoint.
func WithInGivenOut[T any](outAmount T, tokenOutDenom string, tokenInDenom string) RouterQuoteOption {
	return func(opts *RouterQuoteOptions) {
		opts.TokenInDenom = []string{tokenInDenom}
		opts.TokenOut = fmt.Sprintf("%v%s", outAmount, tokenOutDenom)
	}
}

// WithInGivenOutCustom sets the options for an in given out swap for the /router/custom-direct-quote endpoint.
func WithInGivenOutCustom[T any](outAmount T, tokenOutDenom string, tokenInDenom []string) RouterQuoteOption {
	return func(opts *RouterQuoteOptions) {
		opts.TokenInDenom = tokenInDenom
		opts.TokenOut = fmt.Sprintf("%v%s", outAmount, tokenOutDenom)
	}
}

// WithHumanDenomsQuote is an option to set the human denoms for the /router/quote endpoint.
func WithHumanDenoms() RouterQuoteOption {
	return func(opts *RouterQuoteOptions) {
		opts.HumanDenoms = true
	}
}

// WithAppendBaseFee sets the options to append the base fee to the quote.
func WithAppendBaseFee() RouterQuoteOption {
	return func(opts *RouterQuoteOptions) {
		opts.AppendBaseFee = true
	}
}

// WithIsSingleRoute sets the options for a single route for the /router/quote endpoint.
// If true, split routes are not returned.
// If false, split routes are attempted to be computed.
func WithIsSingleRoute() RouterQuoteOption {
	return func(opts *RouterQuoteOptions) {
		opts.IsSingleRoute = true
	}
}

// WithIsSingleRoute sets the options for a single route.

var _ Options = &RouterQuoteOptions{}
