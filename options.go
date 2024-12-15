package sqs

import "net/url"

// Options is the interface for the options for the SQS client.
type Options interface {
	// Validate validates the options.
	Validate() error
	// CreateQueryParams creates the query params for the options.
	CreateQueryParams() url.Values
}
