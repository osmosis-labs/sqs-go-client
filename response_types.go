package sqsclient

type OsmosisTokenMetadata struct {
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	CoinMinimalDenom string `json:"coinMinimalDenom"`
	Decimals         int    `json:"decimals"`
	Preview          bool   `json:"preview"`
	CoingeckoId      string `json:"coingeckoId"`
}

type SQSQuoteResponse struct {
	AmountIn                coin         "json:\"amount_in\""
	AmountOut               string       "json:\"amount_out\""
	Route                   []sqsRoute   "json:\"route\""
	EffectiveFee            string       "json:\"effective_fee\""
	PriceImpact             string       "json:\"price_impact\""
	InBaseOutQuoteSpotPrice string       "json:\"in_base_out_quote_spot_price\""
	PriceInfo               sqsPriceInfo "json:\"price_info\""
}

// sqsPool is the implementation of the routable pool for sqs
type sqsPool struct {
	ID            uint64 "json:\"id\""
	Type          int32  "json:\"type\""
	Balances      []coin "json:\"balances\""
	SpreadFactor  string "json:\"spread_factor\""
	TokenOutDenom string "json:\"token_out_denom,omitempty\""
	TokenInDenom  string "json:\"token_in_denom,omitempty\""
	TakerFee      string "json:\"taker_fee\""
	CodeID        uint64 "json:\"code_id,omitempty\""
}

// sqsRoute is the implementation of the route for sqs
type sqsRoute struct {
	Pools                      []sqsPool "json:\"pools\""
	HasGeneralizedCosmWasmPool bool      "json:\"has-cw-pool\""
	OutAmount                  string    "json:\"out_amount\""
	InAmount                   string    "json:\"in_amount\""
}

// sqsPriceInfo represents the price info returned by the SQS API.
type sqsPriceInfo struct {
	AdjustedGasUsed uint64 `json:"adjusted_gas_used,omitempty"`
	FeeCoin         coin   `json:"fee_coin,omitempty"`
	BaseFee         string `json:"base_fee"`
	Err             string `json:"error,omitempty"`
}

// coin is the implementation of the coin for sqs
type coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
