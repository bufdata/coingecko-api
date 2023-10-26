package geckoterminal

// NetworksResponse returned by GetAllNetworks or GetDexes API.
type NetworksResponse struct {
	Data      []struct{ NetworksItem } `json:"data"`
	LinksItem `json:"links"`
}

// GetSpecificPoolResponse returned by GetSpecificPool API.
type GetSpecificPoolResponse struct {
	Data    PoolDataItem       `json:"data"`
	Include []PoolIncludedItem `json:"include,omitempty"`
}

// GetMultiPoolsResponse returned by GetMultiPools API.
type GetMultiPoolsResponse struct {
	Data    []PoolDataItem     `json:"data"`
	Include []PoolIncludedItem `json:"include,omitempty"`
}
