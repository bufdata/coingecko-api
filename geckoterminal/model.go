package geckoterminal

// NetworksResponse returned by GetAllNetworks or GetDexes API.
type NetworksResponse struct {
	Data      []struct{ NetworksItem } `json:"data"`
	LinksItem `json:"links"`
}

// PoolsResponse returned by pools APIs.
type PoolsResponse struct {
	Data    PoolDataItem       `json:"data"`
	Include []PoolIncludedItem `json:"include,omitempty"`
}

// ErrorResponse is returned when failing to call API.
type ErrorResponse struct {
	Errors []struct {
		Status string `json:"status"`
		Title  string `json:"title"`
	} `json:"errors"`
}
