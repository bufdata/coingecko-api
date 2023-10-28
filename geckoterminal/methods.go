package geckoterminal

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
	"strings"
)

// GetNetworks gets list of supported networks.
//
// Query parameters:
//
// page(optional): page through results.
//
// Note: rate limit for this API is 30 calls per minute.
func (c *Client) GetNetworks(ctx context.Context, page uint) (*NetworksResponse, error) {
	params := url.Values{}
	if page == 0 {
		page = 1
	}
	params.Add("page", strconv.Itoa(int(page)))

	endpoint := fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, getNetworksPath, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to networks api", "error", err)
		return nil, err
	}

	var data NetworksResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal networks response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetDexes gets list of supported dexes on a network.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Path parameters:
//
// network(required): network id from /networks list. Example: eth.
//
// Query parameters:
//
// page(optional): page through results.
func (c *Client) GetDexes(ctx context.Context, network string, page uint) (*DexesResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}

	params := url.Values{}
	if page == 0 {
		page = 1
	}
	params.Add("page", strconv.Itoa(int(page)))

	path := fmt.Sprintf(getDexesPath, network)
	endpoint := fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get dexes api", "error", err)
		return nil, err
	}

	var data DexesResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get dexes response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetSpecificPool gets specific pool on a network.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Path parameters:
//
// network(required): network id from /networks list. Example: eth.
//
// address(required): pool address. Example: 0x60594a405d53811d3bc4766596efd80fd545a270.
//
// Query parameters:
//
// include(optional): Attributes for related resources to include, which will be returned under the top-level
// "included" key. Available resources: base_token, quote_token, dex. Example: base_token,quote_token.
func (c *Client) GetSpecificPool(ctx context.Context, network, address string, include []string) (*SpecificPoolResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}
	if address == "" {
		return nil, fmt.Errorf("address should not be empty")
	}

	params := url.Values{}
	if len(include) != 0 {
		includeParam := strings.Join(include, ",")
		params.Add("include", includeParam)
	}

	path := fmt.Sprintf(getSpecificPoolPath, network, address)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	}

	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get specific pool api", "error", err)
		return nil, err
	}

	var data SpecificPoolResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get specific pool response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetMultiPools gets multi pools on a network.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Path parameters:
//
// network(required): network id from /networks list. Example: eth.
//
// addresses(required): comma-separated list of pool addresses (up to 30 addresses); addresses not found in the
// GeckoTerminal database will be ignored.
// Example: 0x60594a405d53811d3bc4766596efd80fd545a270,0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640.
//
// Query parameters:
//
// include(optional): Attributes for related resources to include, which will be returned under the top-level
// "included" key. Available resources: base_token, quote_token, dex. Example: base_token,quote_token.
func (c *Client) GetMultiPools(ctx context.Context, network string, include, addresses []string) (*PoolsResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}
	if len(addresses) == 0 {
		return nil, fmt.Errorf("addresses should not be empty")
	}

	params := url.Values{}
	if len(include) != 0 {
		includeParam := strings.Join(include, ",")
		params.Add("include", includeParam)
	}

	address := strings.Join(addresses, ",")
	path := fmt.Sprintf(getMultiPoolsPath, network, address)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	}

	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get multi pools api", "error", err)
		return nil, err
	}

	var data PoolsResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get multi pools response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetTop20PoolsOnOneNetwork gets top 20 pools on a network.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Path parameters:
//
// network(required): network id from /networks list. Example: eth.
//
// Query parameters:
//
// include(optional): Attributes for related resources to include, which will be returned under the top-level
// "included" key. Available resources: base_token, quote_token, dex. Example: base_token,quote_token.
func (c *Client) GetTop20PoolsOnOneNetwork(ctx context.Context, network string, include []string) (*PoolsResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}

	params := url.Values{}
	if len(include) != 0 {
		includeParam := strings.Join(include, ",")
		params.Add("include", includeParam)
	}

	path := fmt.Sprintf(getTop20PoolsPath, network)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	}

	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get top 20 pools api", "error", err)
		return nil, err
	}

	var data PoolsResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get top 20 pools response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetTop20PoolsOnOneDex gets top 20 pools on a network's dex.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Path parameters:
//
// network(required): network id from /networks list. Example: eth.
//
// Query parameters:
//
// dex(required): dex id from /networks/{network}/dexes list. Example: sushiswap.
//
// include(optional): Attributes for related resources to include, which will be returned under the top-level
// "included" key. Available resources: base_token, quote_token, dex. Example: base_token,quote_token.
func (c *Client) GetTop20PoolsOnOneDex(ctx context.Context, network, dex string, include []string) (*PoolsResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}

	params := url.Values{}
	if len(include) != 0 {
		includeParam := strings.Join(include, ",")
		params.Add("include", includeParam)
	}

	path := fmt.Sprintf(getTop20PoolsOnOneDexPath, network, dex)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	}

	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get top 20 pools on one dex api", "error", err)
		return nil, err
	}

	var data PoolsResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get top 20 pools on one dex response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetLatest20PoolsOnOneNetwork gets latest 20 pools on a network.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Path parameters:
//
// network(required): network id from /networks list. Example: eth.
//
// Query parameters:
//
// include(optional): Attributes for related resources to include, which will be returned under the top-level
// "included" key. Available resources: base_token, quote_token, dex. Example: base_token,quote_token.
func (c *Client) GetLatest20PoolsOnOneNetwork(ctx context.Context, network string, include []string) (*PoolsResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}

	params := url.Values{}
	if len(include) != 0 {
		includeParam := strings.Join(include, ",")
		params.Add("include", includeParam)
	}

	path := fmt.Sprintf(getLatest20PoolsOnOneNetworkPath, network)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	}

	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get latest 20 pools on one network api", "error", err)
		return nil, err
	}

	var data PoolsResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get latest 20 pools on one network response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetLatest20PoolsOnAllNetworks gets latest 20 pools across all networks.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Query parameters:
//
// include(optional): Attributes for related resources to include, which will be returned under the top-level
// "included" key. Available resources: base_token, quote_token, dex. Example: base_token,quote_token.
func (c *Client) GetLatest20PoolsOnAllNetworks(ctx context.Context, include []string) (*PoolsResponse, error) {
	params := url.Values{}
	if len(include) != 0 {
		includeParam := strings.Join(include, ",")
		params.Add("include", includeParam)
	}

	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, getLatest20PoolsOnAllNetworkPath, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, getLatest20PoolsOnAllNetworkPath)
	}

	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get latest 20 pools on all networks api", "error", err)
		return nil, err
	}

	var data PoolsResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get latest 20 pools on all networks response", "error", err)
		return nil, err
	}
	return &data, nil
}

// SearchPools search for pools on a network.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Query parameters:
//
// query(optional): Search query: can be pool address, token address, or token symbol. Returns top 5 matching pools.
// Example: ETH.
//
// network(optional): network id from /networks list. Example: eth.
//
// include(optional): Attributes for related resources to include, which will be returned under the top-level
// "included" key. Available resources: base_token, quote_token, dex. Example: base_token,quote_token.
func (c *Client) SearchPools(ctx context.Context, query, network string, include []string) (*PoolsResponse, error) {
	params := url.Values{}
	if query != "" {
		params.Add("query", query)
	}
	if network != "" {
		params.Add("network", network)
	}
	if len(include) != 0 {
		includeParam := strings.Join(include, ",")
		params.Add("include", includeParam)
	}

	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, searchPoolsPath, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, searchPoolsPath)
	}

	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to search pools api", "error", err)
		return nil, err
	}

	var data PoolsResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal search pools response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetTop20PoolsForOneToken gets top 20 pools for a token.
// Contains special field token_price_usd representing price of requested token.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Path parameters:
//
// network(required): network id from /networks list. Example: eth.
//
// token_address(required): address of token. Example: 0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48.
//
// Query parameters:
//
// include(optional): Attributes for related resources to include, which will be returned under the top-level
// "included" key. Available resources: base_token, quote_token, dex. Example: base_token,quote_token.
func (c *Client) GetTop20PoolsForOneToken(ctx context.Context, network, tokenAddress string, include []string) (*PoolsResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}
	if tokenAddress == "" {
		return nil, fmt.Errorf("token_address should not be empty")
	}

	params := url.Values{}
	if len(include) != 0 {
		includeParam := strings.Join(include, ",")
		params.Add("include", includeParam)
	}

	path := fmt.Sprintf(getTop20PoolsForOneTokenPath, network, tokenAddress)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	}

	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get top 20 pools for one token api", "error", err)
		return nil, err
	}

	var data PoolsResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get top 20 pools for one token response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetSpecificTokenOnOneNetwork gets specific token on a network.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Path parameters:
//
// network(required): network id from /networks list. Example: eth.
//
// address(required): token address. Example: 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2.
//
// Query parameters:
//
// include(optional): Attributes for related resources to include, which will be returned under the top-level
// "included" key. Available resources: top_pools. Example: top_pools.
func (c *Client) GetSpecificTokenOnOneNetwork(ctx context.Context, network, address string, include []string) (*TokenResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}
	if address == "" {
		return nil, fmt.Errorf("address should not be empty")
	}

	params := url.Values{}
	if len(include) != 0 {
		includeParam := strings.Join(include, ",")
		params.Add("include", includeParam)
	}

	path := fmt.Sprintf(getSpecificTokenOnOneNetworkPath, network, address)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	}

	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get specific token on one network api", "error", err)
		return nil, err
	}

	var data TokenResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get specific token on one network response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetMultiTokensOnOneNetwork gets multi tokens on a network.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Path parameters:
//
// network(required): network id from /networks list. Example: eth.
//
// addresses(required): comma-separated list of token addresses (up to 30 addresses).
// addresses not found in the GeckoTerminal database will be ignored.
// Note: top_pools for this endpoint returns only the first top pool for each token.
// Example: 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2,0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48.
//
// Query parameters:
//
// include(optional): Attributes for related resources to include, which will be returned under the top-level
// "included" key. Available resources: top_pools. Example: top_pools.
func (c *Client) GetMultiTokensOnOneNetwork(ctx context.Context, network, address string, include []string) (*TokenDataItem, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}
	if address == "" {
		return nil, fmt.Errorf("address should not be empty")
	}

	params := url.Values{}
	if len(include) != 0 {
		includeParam := strings.Join(include, ",")
		params.Add("include", includeParam)
	}

	path := fmt.Sprintf(getMultiTokensOnOneNetworkPath, network, address)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	}

	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get multi tokens on one network api", "error", err)
		return nil, err
	}

	var data TokenDataItem
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get multi tokens on one network response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetSpecificTokenInfoOnOneNetwork gets specific token info on a network.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Path parameters:
//
// network(required): network id from /networks list. Example: eth.
//
// addresses(required): token address. Example: 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
func (c *Client) GetSpecificTokenInfoOnOneNetwork(ctx context.Context, network, address string) (*TokenInfoResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}
	if address == "" {
		return nil, fmt.Errorf("address should not be empty")
	}

	path := fmt.Sprintf(getSpecificTokenInfoOnOneNetworkPath, network, address)
	endpoint := fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get specific token info on one network api", "error", err)
		return nil, err
	}

	var data TokenInfoResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get specific token info on one network response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetPoolTokensInfoOnOneNetwork gets pool tokens info on a network.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Path parameters:
//
// network(required): network id from /networks list. Example: eth.
//
// pool_address(required): pool address. Example: 0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852
func (c *Client) GetPoolTokensInfoOnOneNetwork(ctx context.Context, network, poolAddress string) (*PoolTokensInfoResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}
	if poolAddress == "" {
		return nil, fmt.Errorf("pool_address should not be empty")
	}

	path := fmt.Sprintf(getPoolTokensInfoOnOneNetworkPath, network, poolAddress)
	endpoint := fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get pool tokens info on one network api", "error", err)
		return nil, err
	}

	var data PoolTokensInfoResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get pool tokens info on one network response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetRecentlyUpdated100TokensInfo gets most recently 100 tokens info across all networks.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Query parameters:
//
// include(optional): attributes for related resources to include, which will be returned under the top-level "included"
// key. Available resources: network. Example: network.
func (c *Client) GetRecentlyUpdated100TokensInfo(ctx context.Context, include string) (*RecentlyUpdatedTokensResponse, error) {
	params := url.Values{}
	if include == "" {
		params.Add("include", include)
	}

	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, getRecentlyUpdated100TokensInfoPath, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, getRecentlyUpdated100TokensInfoPath)
	}
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get recently updated 100 tokens info api", "error", err)
		return nil, err
	}

	var data RecentlyUpdatedTokensResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal recently updated tokens response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetOHLCV gets OHLCV data of a pool, up to 6 months ago. Empty response if there is no earlier data available.
//
// Note: rate limit for this API is 30 calls per minute.
//
// Path parameters:
//
// network(required): network id from /networks list. Example: eth.
//
// pool_address(required): pool address. Example: 0x60594a405d53811d3bc4766596efd80fd545a270.
// Note: Pools with more than 2 tokens are not yet supported for this endpoint.
//
// timeframe(required): timeframe. Available values: day, hour, minute. Example: day.
//
// Query parameters:
//
// aggregate(optional): time period to aggregate for each ohlcv (e.g. /minute?aggregate=15 for 15m ohlcv).
// Available values (day): 1. Available values (hour): 1, 4, 12. Available values (minute): 1, 5, 15. Default: 1.
//
// before_timestamp(optional): return ohlcv data before this timestamp (integer seconds since epoch).
// Example: 1679414400.
//
// limit(optional): limit number of ohlcv results to return (default: 100, max: 1000). Example: 100.
//
// currency(optional): return ohlcv in USD or quote token (default: usd). Available values: usd, token.
//
// token(optional): return ohlcv for base or quote token; use this to invert the chart. (default: base).
// Available values: base, quote.
func (c *Client) GetOHLCV(ctx context.Context, network, poolAddress, timeframe string, aggregate uint, beforeTimestamp int64,
	limit uint, currency, token string) (*OHLCVResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}
	if poolAddress == "" {
		return nil, fmt.Errorf("pool_address should not be empty")
	}
	if timeframe == "" {
		return nil, fmt.Errorf("timeframe should not be empty")
	}

	params := url.Values{}
	if aggregate == 0 {
		aggregate = 1
	}
	params.Add("aggregate", strconv.Itoa(int(aggregate)))
	if beforeTimestamp != 0 {
		params.Add("before_timestamp", strconv.Itoa(int(beforeTimestamp)))
	}
	if limit == 0 {
		limit = 100
	}
	params.Add("limit", strconv.Itoa(int(limit)))
	if currency == "" {
		currency = "usd"
	}
	params.Add("currency", currency)
	if token == "" {
		params.Add("token", token)
	}

	path := fmt.Sprintf(getOHLCVPath, network, poolAddress, timeframe)
	endpoint := fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get OHLCV api", "error", err)
		return nil, err
	}

	var data OHLCVResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal OHLCV response", "error", err)
		return nil, err
	}
	return &data, nil
}
