package geckoterminal

// NetworksResponse returned by GetAllNetworks or GetDexes API.
type NetworksResponse struct {
	Data      []struct{ NetworksItem } `json:"data"`
	LinksItem `json:"links"`
}

// PoolsResponse returned by pools APIs.
type PoolsResponse struct {
	Data     PoolDataItem       `json:"data"`
	Included []PoolIncludedItem `json:"included,omitempty"`
}

// TokenResponse returned by token APIs.
type TokenResponse struct {
	Data     TokenDataItem `json:"data"`
	Included []struct {
		basicStruct
		PoolIncludedItem
	} `json:"included,omitempty"`
}

// TokenInfoResponse used in token info APIs.
type TokenInfoResponse struct {
	Data TokenInfoDataItem `json:"data"`
}

// PoolTokensInfoResponse
type PoolTokensInfoResponse struct {
	Data []TokenInfoDataItem `json:"data"`
}

// RecentlyUpdatedTokensResponse
type RecentlyUpdatedTokensResponse struct {
	Data    []TokenInfoDataItem      `json:"data"`
	Include []TokensInfoIncludedItem `json:"included"`
}

// OHLCVResponse
type OHLCVResponse struct {
	basicStruct
	Attributes struct {
		OHLCVList []OHLCVItem `json:"ohlcv_list"`
	} `json:"attributes"`
}

type OHLCVItem [6]float64

// ErrorResponse is returned when failing to call API.
type ErrorResponse struct {
	Errors []struct {
		Status string `json:"status"`
		Title  string `json:"title"`
	} `json:"errors"`
}
