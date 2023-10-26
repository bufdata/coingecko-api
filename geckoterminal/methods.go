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
// "included" key. Available resources: base_token, quote_token, dex. Example: base_token,quote_token
func (c *Client) GetSpecificPool(ctx context.Context, network, address, include string) (*GetSpecificPoolResponse, error) {
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

	var data GetSpecificPoolResponse
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
// "included" key. Available resources: base_token, quote_token, dex. Example: base_token,quote_token
func (c *Client) GetMultiPools(ctx context.Context, network, include string, addresses []string) (*GetMultiPoolsResponse, error) {
	if network == "" {
		return nil, fmt.Errorf("network should not be empty")
	}
	if len(addresses) == 0 {
		return nil, fmt.Errorf("addresses should not be empty")
	}

	params := url.Values{}
	if include != "" {
		params.Add("include", include)
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

	var data GetMultiPoolsResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal get multi pools response", "error", err)
		return nil, err
	}
	return &data, nil
}
