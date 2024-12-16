package sqsclient

import "errors"

// InitializeOptions are the options for the Initialize function.
type InitializeOptions struct {
	Environment SQSEnvironment
	CustomURL   string
	APIKey      string
}

// Validate validates the InitializeOptions.
func (opts *InitializeOptions) Validate() error {
	if opts.Environment == "" && opts.CustomURL == "" {
		return errors.New("one of environment or custom url is required")
	}

	if opts.Environment != "" && opts.CustomURL != "" {
		return errors.New("only one of environment or custom url is allowed")
	}

	return nil
}

// GetURL returns the URL to use for the SQS client.
// If a custom URL is provided, it returns the custom URL.
// Otherwise, it returns the URL for the environment.
// CONTRACT: Validate() must pass for this to work.
func (opts *InitializeOptions) GetURL() string {
	if opts.CustomURL != "" {
		return opts.CustomURL
	}

	return EnvironmentURLMap[opts.Environment]
}

// InitializeOption is a function that modifies the InitializeOptions.
type InitializeOption func(opts *InitializeOptions)

// WithEnvironmentOpt is an option to set the environment for the SQS client.
func WithEnvironmentOpt(environment SQSEnvironment) InitializeOption {
	return func(opts *InitializeOptions) {
		opts.Environment = environment
	}
}

// WithAPIKeyOpt is an option to set the API key for the SQS client.
func WithAPIKeyOpt(apiKey string) InitializeOption {
	return func(opts *InitializeOptions) {
		opts.APIKey = apiKey
	}
}

// WithCustomURL is an option to set the custom URL for the SQS client.
func WithCustomURL(customURL string) InitializeOption {
	return func(opts *InitializeOptions) {
		opts.CustomURL = customURL
	}
}

// Initialize initializes a new SQS client.
// It validates the options and returns a new SQS client.
// If no environment is provided, it defaults to Prod.
// Only one of environment or custom url is allowed.
func Initialize(options ...InitializeOption) (SQSClient, error) {
	opts := InitializeOptions{}

	// Apply the options
	for _, option := range options {
		option(&opts)
	}

	// If no environment or custom URL is provided, it defaults to Prod.
	if opts.Environment == "" && opts.CustomURL == "" {
		opts.Environment = Prod
	}

	// Validate the options
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	// Get the URL
	url := opts.GetURL()

	// Create the SQS client
	sqs := NewClient(url)

	// Add the API key if applicable.
	if opts.APIKey != "" {
		sqs = WithAPIKey(opts.APIKey, sqs)
	}

	return sqs, nil
}
