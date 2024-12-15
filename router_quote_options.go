package sqs

import (
	"fmt"
	"net/url"
	"strconv"
)

// RouterQuoteOptions is the options for the /router/quote endpoint.
type RouterQuoteOptions struct {
	// Out given in.
	// TokenIn is the token in and denom to swap from.
	// E.g. 10uosmo
	TokenIn string
	// TokenOutDenom is the denom to swap to.
	TokenOutDenom string

	// In given out.
	// TokenOut is the token out and denom to swap to.
	// E.g. 10uatom
	TokenOut string
	// TokenInDenom is the denom to swap from.
	TokenInDenom string

	// HumanDenoms is whether the input tokens are human readable denoms.
	HumanDenoms bool
	// IsSingleRoute is whether the quote is for a single route.
	// If true, split routes are not returned.
	// If false, split routes are attempted to be computed.
	IsSingleRoute bool
}

// RouterQuoteOption is the type for the options for the /router/quote endpoint.
type RouterQuoteOption func(opts *RouterQuoteOptions)

// Validate validates the RouterQuoteOptions.
// It returns an error if the options are invalid.
func (o *RouterQuoteOptions) Validate() error {
	if o.TokenIn == "" && o.TokenOut == "" {
		return fmt.Errorf("token in or token out must be set")
	}
	if o.TokenInDenom == "" && o.TokenOutDenom == "" {
		return fmt.Errorf("token in denom or token out denom must be set")
	}

	if o.TokenIn != "" && o.TokenOut != "" {
		return fmt.Errorf("token in and token out cannot be set at the same time")
	}
	if o.TokenInDenom != "" && o.TokenOutDenom != "" {
		return fmt.Errorf("token in denom and token out denom cannot be set at the same time")
	}

	return nil
}

func (o *RouterQuoteOptions) IsOutGivenIn() bool {
	return o.TokenIn != "" && o.TokenOutDenom != ""
}

func (o *RouterQuoteOptions) CreateQueryParams() url.Values {
	queryParams := url.Values{}
	queryParams.Add("humanDenoms", strconv.FormatBool(o.HumanDenoms))
	queryParams.Add("singleRoute", strconv.FormatBool(o.IsSingleRoute))

	if o.IsOutGivenIn() {
		queryParams.Add("tokenIn", o.TokenIn)
		queryParams.Add("tokenOutDenom", o.TokenOutDenom)
	} else {
		queryParams.Add("tokenInDenom", o.TokenInDenom)
		queryParams.Add("tokenOut", o.TokenOut)
	}

	if o.HumanDenoms {
		queryParams.Add("humanDenoms", "true")
	}

	if o.IsSingleRoute {
		queryParams.Add("singleRoute", "true")
	}
	return queryParams
}

var _ Options = &RouterQuoteOptions{}
