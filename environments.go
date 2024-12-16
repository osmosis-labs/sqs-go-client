package sqsclient

// SQSEnvironment represents an environment supported by the Osmosis team.
type SQSEnvironment string

const (
	Stage   SQSEnvironment = "stage"
	Prod    SQSEnvironment = "prod"
	Testnet SQSEnvironment = "testnet"
)

const (
	DefaultStageURL   = "https://sqs.stage.osmosis.zone"
	DefaultProdURL    = "https://sqs.osmosis.zone"
	DefaultTestnetURL = "https://sqs.testnet.osmosis.zone"
)

var (
	EnvironmentURLMap = map[SQSEnvironment]string{
		Stage:   DefaultStageURL,
		Prod:    DefaultProdURL,
		Testnet: DefaultTestnetURL,
	}
)
