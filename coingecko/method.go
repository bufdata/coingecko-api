package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"strings"
)

// Ping checks API server status.
func (c *Client) Ping(ctx context.Context) (*PingResponse, error) {
	endpoint := c.checkAPIKey(fmt.Sprintf("%s%s", c.apiURL, pingPath))
	resp, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to ping api", "error", err)
		return nil, err
	}

	var data PingResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal ping response", "error", err)
		return nil, err
	}
	return &data, nil
}

// SimplePrice gets the current price of any cryptocurrencies in any other supported currencies that you need.
//
// Note: to check if a price is stale, please flag include_last_updated_at=true to get the latest updated time.
// You may also flag include_24hr_change=true to check if it returns 'null' value.
//
// Cache/Update Frequency: every 60 seconds(every 30 seconds for Pro API).
//
// Query parameters:
//
// ids: id of coins, comma-separated if querying more than 1 coin;
// refers to coins/list.
//
// vs_currencies: vs_currency of coins, comma-separated if querying more than 1 vs_currency;
// refers to simple/supported_vs_currencies.
//
// include_market_cap: true/false to include market_cap, default: false.
//
// include_24hr_vol: true/false to include 24hr_vol, default: false.
//
// include_24hr_change: true/false to include 24hr_change, default: false.
//
// include_last_updated_at: true/false to include last_updated_at of price, default: false.
//
// precision: full or any value between 0-18 to specify decimal place for currency price value.
func (c *Client) SimplePrice(ctx context.Context, ids, vsCurrencies []string, includeMarketCap, include24hrVol,
	include24hrChange, includeLastUpdatedAt, precision string) (*map[string]map[string]float32, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("the length of ids should be greater than 0")
	}
	if len(vsCurrencies) == 0 {
		return nil, fmt.Errorf("the length of vsCurrencies should be greater than 0")
	}

	idsParams := strings.Join(ids, ",")
	vsCurrenciesParams := strings.Join(vsCurrencies, ",")

	params := url.Values{}
	params.Add("ids", idsParams)
	params.Add("vs_currencies", vsCurrenciesParams)
	// TODO: the following items should be optimized, SimpleTokenPrice, too.
	if includeMarketCap != "" {
		params.Add("include_market_cap", includeMarketCap)
	}
	if include24hrVol != "" {
		params.Add("include_24hr_vol", include24hrVol)
	}
	if include24hrChange != "" {
		params.Add("include_24hr_change", include24hrChange)
	}
	if includeLastUpdatedAt != "" {
		params.Add("include_last_updated_at", includeLastUpdatedAt)
	}
	if precision != "" {
		params.Add("precision", precision)
	}

	endpoint := c.checkAPIKey(fmt.Sprintf("%s%s?%s", c.apiURL, simplePricePath, params.Encode()))
	resp, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to simple price api", "error", err)
		return nil, err
	}

	data := make(map[string]map[string]float32)
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal simple price response", "error", err)
		return nil, err
	}
	return &data, nil
}

// SimpleTokenPrice gets current price of tokens(using contract address) for a given platform in any other currency
// that you need.
//
// This endpoint allows you to query a token price by using contract address. It returns the global average price that
// is aggregated across all active exchanges on CoinGecko. It does not return the price of a specific network, you may
// cross-check the price on CoinGecko.com.
//
// Learn more about our price methodology https://www.coingecko.com/en/methodology.
//
// Cache/Update Frequency: every 60 seconds(every 30 seconds for Pro API).
//
// Query parameters:
//
// id: the id of the platform issuing tokens (See asset_platforms endpoint for list of options).
//
// contract_address: the contract address of tokens, comma separated.
//
// vs_currencies: vs_currency of coins, comma-separated if querying more than 1 vs_currency;
// refers to simple/supported_vs_currencies.
//
// include_market_cap: true/false to include market_cap, default: false.
//
// include_24hr_vol: true/false to include 24hr_vol, default: false.
//
// include_24hr_change: true/false to include 24hr_change, default: false.
//
// include_last_updated_at: true/false to include last_updated_at of price, default: false.
//
// precision: full or any value between 0-18 to specify decimal place for currency price value.
func (c *Client) SimpleTokenPrice(ctx context.Context, id string, contractAddress, vsCurrencies []string, includeMarketCap, include24hrVol,
	include24hrChange, includeLastUpdatedAt, precision string) (*map[string]map[string]float32, error) {
	if id == "" {
		return nil, fmt.Errorf("id shouldn't be empty")
	}
	if len(contractAddress) == 0 {
		return nil, fmt.Errorf("the length of contract address should be greater than 0")
	}
	if len(vsCurrencies) == 0 {
		return nil, fmt.Errorf("the length of vs currencies should be greater than 0")
	}

	contractAddressParams := strings.Join(contractAddress, ",")
	vsCurrenciesParams := strings.Join(vsCurrencies, ",")

	params := url.Values{}
	params.Add("contract_addresses", contractAddressParams)
	params.Add("vs_currencies", vsCurrenciesParams)
	if includeMarketCap != "" {
		params.Add("include_market_cap", includeMarketCap)
	}
	if include24hrVol != "" {
		params.Add("include_24hr_vol", include24hrVol)
	}
	if include24hrChange != "" {
		params.Add("include_24hr_change", include24hrChange)
	}
	if includeLastUpdatedAt != "" {
		params.Add("include_last_updated_at", includeLastUpdatedAt)
	}
	if precision != "" {
		params.Add("precision", precision)
	}

	path := fmt.Sprintf(simpleTokenPricePath, id)
	fmt.Println(path)
	endpoint := c.checkAPIKey(fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode()))
	fmt.Println(endpoint)
	resp, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to simple token price api", "error", err)
		return nil, err
	}

	data := make(map[string]map[string]float32)
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal simple token price response", "error", err)
		return nil, err
	}
	return &data, nil
}

// SimpleSupportedVSCurrencies gets list of supported_vs_currencies.
//
// Cache/Update Frequency: every 60 seconds.
func (c *Client) SimpleSupportedVSCurrencies(ctx context.Context) (*SimpleSupportedVSCurrenciesResponse, error) {
	endpoint := c.checkAPIKey(fmt.Sprintf("%s%s", c.apiURL, supportedVsCurrenciesPath))
	resp, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to simple supported vs currencies api", "error", err)
		return nil, err
	}

	var data SimpleSupportedVSCurrenciesResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal simple supported vs currencies response", "error", err)
		return nil, err
	}
	return &data, nil
}
