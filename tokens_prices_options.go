package sqsclient

import (
	"fmt"
	"net/url"
	"strings"
)

// TokenPricesOptions is the type for the options for the /tokens/prices endpoint.
type TokenPricesOptions struct {
	// HumanDenoms is a flag to set the human denoms for the /tokens/prices endpoint.
	HumanDenoms bool

	// BaseDenoms is a list of base denoms to get the prices for.
	BaseDenoms []string
}

// TokenPricesOption is the type for the options for the /tokens/prices endpoint.
type TokenPricesOption func(opts *TokenPricesOptions)

// WithHumanDenomsPrices is an option to set the human denoms for the /tokens/prices endpoint.
func WithHumanDenomsPrices() TokenPricesOption {
	return func(opts *TokenPricesOptions) {
		opts.HumanDenoms = true
	}
}

// WithBaseDenoms is an option to set the base denoms for the /tokens/prices endpoint.
func WithBaseDenoms(denoms []string) TokenPricesOption {
	return func(opts *TokenPricesOptions) {
		opts.BaseDenoms = denoms
	}
}

// WithBaseDenom is an option to set the base denom for the /tokens/prices endpoint.
func WithBaseDenom(denom string) TokenPricesOption {
	return WithBaseDenoms([]string{denom})
}

// Validate validates the options for the /tokens/prices endpoint.
func (opts *TokenPricesOptions) Validate() error {

	if len(opts.BaseDenoms) == 0 {
		return fmt.Errorf("base denoms is required")
	}

	return nil
}

func (opts *TokenPricesOptions) CreateQueryParams() url.Values {
	queryParams := url.Values{}

	if opts.HumanDenoms {
		queryParams.Add("humanDenoms", "true")
	} else {
		queryParams.Add("humanDenoms", "false")
	}

	queryParams.Add("base", strings.Join(opts.BaseDenoms, ","))

	return queryParams
}

var _ Options = &TokenPricesOptions{}
