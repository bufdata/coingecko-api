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

// GetAllNetworks gets list of supported networks.
//
// Query parameters:
//
// page(optional): page through results.
//
// Note: rate limit for this API is 30 calls per minute.
func (c *Client) GetAllNetworks(ctx context.Context, page uint) (*NetworksResponse, error) {
	params := url.Values{}
	if page == 0 {
		page = 1
	}
	params.Add("page", strconv.Itoa(int(page)))

	endpoint := fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, networksPath, params.Encode())
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
func (c *Client) GetDexes(ctx context.Context, network string, page uint) (*NetworksResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}

	params := url.Values{}
	if page == 0 {
		page = 1
	}
	params.Add("page", strconv.Itoa(int(page)))

	path := fmt.Sprintf(dexesPath, network)
	endpoint := fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get dexes api", "error", err)
		return nil, err
	}

	var data NetworksResponse
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
// address(required): pool address. Example: 0x60594a405d53811d3bc4766596efd80fd545a270
//
// Query parameters:
//
// include(optional): Attributes for related resources to include, which will be returned under the top-level
// "included" key. Available resources: base_token, quote_token, dex. Example: base_token,quote_token.
func (c *Client) GetSpecificPool(ctx context.Context, network, address, include string) (*PoolsResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}
	if address == "" {
		return nil, fmt.Errorf("address should not be empty")
	}

	params := url.Values{}
	if include != "" {
		params.Add("include", include)
	}

	path := fmt.Sprintf(networksIDPoolsAddressPath, network, address)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	} else {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
	}

	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get specific pool api", "error", err)
		return nil, err
	}

	var data PoolsResponse
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
// Example: 0x60594a405d53811d3bc4766596efd80fd545a270,0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640
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
		includeItem := strings.Join(include, ",")
		params.Add("include", includeItem)
	}

	address := strings.Join(addresses, ",")
	path := fmt.Sprintf(networksIDPoolsMultiPath, network, address)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	} else {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
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

// GetTop20Pools gets top 20 pools on a network.
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
func (c *Client) GetTop20Pools(ctx context.Context, network string, include []string) (*PoolsResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}

	params := url.Values{}
	if len(include) != 0 {
		includeItem := strings.Join(include, ",")
		params.Add("include", includeItem)
	}

	path := fmt.Sprintf(networksIDPoolsPath, network)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	} else {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
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
		includeItem := strings.Join(include, ",")
		params.Add("include", includeItem)
	}

	path := fmt.Sprintf(networksIDDexesPoolsPath, network, dex)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	} else {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
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
		includeItem := strings.Join(include, ",")
		params.Add("include", includeItem)
	}

	path := fmt.Sprintf(networksIDNewPoolsPath, network)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, path)
	} else {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, path, params.Encode())
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
		includeItem := strings.Join(include, ",")
		params.Add("include", includeItem)
	}

	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, networksNewPoolsPath)
	} else {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, networksNewPoolsPath, params.Encode())
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
		includeItem := strings.Join(include, ",")
		params.Add("include", includeItem)
	}

	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s", geckoTerminalAPIEndpoint, networksNewPoolsPath)
	} else {
		endpoint = fmt.Sprintf("%s%s?%s", geckoTerminalAPIEndpoint, networksNewPoolsPath, params.Encode())
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
