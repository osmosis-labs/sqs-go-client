package sqs

func WithIsSingleRoute(opts *RouterQuoteOptions) {
	opts.IsSingleRoute = true
}

type TokenPricesOptions struct {
	HumanDenoms bool
}
