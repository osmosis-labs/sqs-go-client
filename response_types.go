package sqsclient

type OsmosisTokenMetadata struct {
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	CoinMinimalDenom string `json:"coinMinimalDenom"`
	Decimals         int    `json:"decimals"`
	Preview          bool   `json:"preview"`
	CoingeckoId      string `json:"coingeckoId"`
}

type sqsPool struct {
	PoolId        uint64 `json:"id"`
	TokenOutDenom string `json:"token_out_denom"`
}

type sqsRoute struct {
	Pools    []sqsPool `json:"pools"`
	InAmount string    `json:"in_amount"`
}

type SQSQuoteResponse struct {
	AmountOut string     `json:"amount_out"`
	Route     []sqsRoute `json:"route"`
}
