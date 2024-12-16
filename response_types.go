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
	AmountIn                Coin      "json:\"amount_in\""
	AmountOut               string    "json:\"amount_out\""
	Route                   []Route   "json:\"route\""
	EffectiveFee            string    "json:\"effective_fee\""
	PriceImpact             string    "json:\"price_impact\""
	InBaseOutQuoteSpotPrice string    "json:\"in_base_out_quote_spot_price\""
	PriceInfo               PriceInfo "json:\"price_info\""
}

// Pool is the implementation of the routable pool for sqs
type Pool struct {
	ID            uint64 "json:\"id\""
	Type          int32  "json:\"type\""
	Balances      []Coin "json:\"balances\""
	SpreadFactor  string "json:\"spread_factor\""
	TokenOutDenom string "json:\"token_out_denom,omitempty\""
	TokenInDenom  string "json:\"token_in_denom,omitempty\""
	TakerFee      string "json:\"taker_fee\""
	CodeID        uint64 "json:\"code_id,omitempty\""
}

// Route is the implementation of the route for sqs
type Route struct {
	Pools                      []Pool "json:\"pools\""
	HasGeneralizedCosmWasmPool bool   "json:\"has-cw-pool\""
	OutAmount                  string "json:\"out_amount\""
	InAmount                   string "json:\"in_amount\""
}

// PriceInfo represents the price info returned by the SQS API.
type PriceInfo struct {
	AdjustedGasUsed uint64 `json:"adjusted_gas_used,omitempty"`
	FeeCoin         Coin   `json:"fee_coin,omitempty"`
	BaseFee         string `json:"base_fee"`
	Err             string `json:"error,omitempty"`
}

// coin is the implementation of the coin for sqs
type Coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}
