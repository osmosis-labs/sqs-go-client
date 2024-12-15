package sqs

import (
	"net/url"
)

// TokenPricesOptions is the type for the options for the /tokens/prices endpoint.
type TokenPricesOptions struct {
	HumanDenoms bool
}

// TokenPricesOption is the type for the options for the /tokens/prices endpoint.
type TokenPricesOption func(opts *TokenPricesOptions)

// WithHumanDenoms is an option to set the human denoms for the /tokens/prices endpoint.
func WithHumanDenoms(humanDenoms bool) TokenPricesOption {
	return func(opts *TokenPricesOptions) {
		opts.HumanDenoms = humanDenoms
	}
}

// Validate validates the options for the /tokens/prices endpoint.
func (opts *TokenPricesOptions) Validate() error {
	return nil
}

func (opts *TokenPricesOptions) CreateQueryParams() url.Values {
	queryParams := url.Values{}

	if opts.HumanDenoms {
		queryParams.Add("humanDenoms", "true")
	} else {
		queryParams.Add("humanDenoms", "false")
	}

	return queryParams
}

var _ Options = &TokenPricesOptions{}
