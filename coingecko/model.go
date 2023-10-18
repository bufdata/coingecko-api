package coingecko

// PingResponse returned by Ping API.
type PingResponse struct {
	GeckoSays string `json:"gecko_says"`
}

// SimpleSupportedVSCurrenciesResponse returned by SimpleSupportedVSCurrencies API.
type SimpleSupportedVSCurrenciesResponse []string

// ErrorResponse defines error response
type ErrorResponse struct {
	Error string `json:"error"`
}
