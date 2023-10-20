package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
	"strings"
)

// Ping checks API server status.
func (c *Client) Ping(ctx context.Context) (*PingResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.apiURL, pingPath)
	resp, _, err := c.sendReq(ctx, endpoint)
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
// ids(required): id of coins, comma-separated if querying more than 1 coin;
// refers to coins/list.
//
// vs_currencies(required): vs_currency of coins, comma-separated if querying more than 1 vs_currency;
// refers to simple/supported_vs_currencies.
//
// include_market_cap(optional): true/false to include market_cap, default: false.
//
// include_24hr_vol(optional): true/false to include 24hr_vol, default: false.
//
// include_24hr_change(optional): true/false to include 24hr_change, default: false.
//
// include_last_updated_at(optional): true/false to include last_updated_at of price, default: false.
//
// precision(optional): full or any value between 0-18 to specify decimal place for currency price value.
func (c *Client) SimplePrice(ctx context.Context, ids, vsCurrencies []string, includeMarketCap, include24hrVol,
	include24hrChange, includeLastUpdatedAt, precision string) (*map[string]map[string]float64, error) {
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

	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, simplePricePath, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to simple price api", "error", err)
		return nil, err
	}

	data := make(map[string]map[string]float64)
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
// Path parameters:
//
// id(required): the id of the platform issuing tokens (See asset_platforms endpoint for list of options).
//
// Query parameters:
//
// contract_addresses(required): the contract address of tokens, comma separated.
//
// vs_currencies(required): vs_currency of coins, comma-separated if querying more than 1 vs_currency;
// refers to simple/supported_vs_currencies.
//
// include_market_cap(optional): true/false to include market_cap, default: false.
//
// include_24hr_vol(optional): true/false to include 24hr_vol, default: false.
//
// include_24hr_change(optional): true/false to include 24hr_change, default: false.
//
// include_last_updated_at(optional): true/false to include last_updated_at of price, default: false.
//
// precision(optional): full or any value between 0-18 to specify decimal place for currency price value.
func (c *Client) SimpleTokenPrice(ctx context.Context, id string, contractAddresses, vsCurrencies []string, includeMarketCap, include24hrVol,
	include24hrChange, includeLastUpdatedAt, precision string) (*map[string]map[string]float64, error) {
	if id == "" {
		return nil, fmt.Errorf("id shouldn't be empty")
	}
	if len(contractAddresses) == 0 {
		return nil, fmt.Errorf("the length of contract addresses should be greater than 0")
	}
	if len(vsCurrencies) == 0 {
		return nil, fmt.Errorf("the length of vs currencies should be greater than 0")
	}

	contractAddressParams := strings.Join(contractAddresses, ",")
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
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to simple token price api", "error", err)
		return nil, err
	}

	data := make(map[string]map[string]float64)
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
	endpoint := fmt.Sprintf("%s%s", c.apiURL, supportedVsCurrenciesPath)
	resp, _, err := c.sendReq(ctx, endpoint)
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

// CoinsList lists all supported coins id, name and symbol(no pagination required).
//
// All the coins that show up on this /coins/list endpoint are Active coins that listed by CoinGecko team on CoinGecko.com
//
// If a coin is inactive or deactivated, it will be removed from /coins/list.
//
// Cache/Update Frequency: every 5 minutes.
//
// Query parameters:
//
// include_platform(optional): flag to include platform contract addresses (eg. 0x.... for Ethereum based tokens).
// valid values: true, false
func (c *Client) CoinsList(ctx context.Context, includePlatform bool) (*[]CoinsListResponse, error) {
	params := url.Values{}
	params.Add("include_platform", strconv.FormatBool(includePlatform))
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, coinsListPath, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to coins list api", "error", err)
		return nil, err
	}

	var data []CoinsListResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coins list response", "error", err)
		return nil, err
	}
	return &data, nil
}

// CoinsMarkets lists all supported coins price, market cap,volume and market related data.
//
// Use this to obtain all the coins market data (price, market cap, volume), per page.
//
// Note: when both 'category' and 'ids' parameters are supplied, the 'category' parameter takes precedence over the
// 'ids' parameter.
//
// Cache/Update Frequency: every 45 seconds.
//
// Query parameters:
//
// vs_currency(required): the target currency of market data (usd, eur, jpy, etc.).
//
// ids(optional): the ids of the coin, comma separated cryptocurrency symbols (base). refers to /coins/list.
//
// category(optional): filter by coin category. Refer to /coin/categories/list.
//
// order(optional): valid values: market_cap_asc, market_cap_desc, volume_asc, volume_desc, id_asc, id_desc
// sort results by field. Default value: market_cap_desc.
//
// per_page(optional): valid values: 1..250; total results per page. Default value: 100.
//
// page(optional): page through results. Default value: 1.
//
// sparkline(optional): include sparkline 7 days data (eg. true, false). Default value: false.
//
// price_change_percentage(optional): include price change percentage in 1h, 24h, 7d, 14d, 30d, 200d, 1y
// (eg. '1h,24h,7d' comma-separated, invalid values will be discarded).
//
// locale(optional): valid values: ar, bg, cs, da, de, el, en, es, fi, fr, he, hi, hr, hu, id, it, ja, ko, lt, nl, no,
// pl, pt, ro, ru, sk, sl, sv, th, tr, uk, vi, zh, zh-tw. Default value: en.
//
// precision(optional): full or any value between 0-18 to specify decimal place for currency price value.
func (c *Client) CoinsMarkets(ctx context.Context, vsCurrency string, ids []string, category, order string, perPage, page uint,
	sparkline bool, priceChangePercentage []string, locale, precision string) (*[]CoinsMarketsResponse, error) {
	if vsCurrency == "" {
		return nil, fmt.Errorf("vs currencies should not be empty")
	}

	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	if len(ids) != 0 {
		id := strings.Join(ids, ",")
		params.Add("ids", id)
	}
	if category != "" {
		params.Add("category", category)
	}
	if order != "" {
		params.Add("order", order)
	}
	if perPage != 0 {
		params.Add("per_page", strconv.Itoa(int(perPage)))
	} else {
		params.Add("per_page", "100")
	}
	if page != 0 {
		params.Add("page", strconv.Itoa(int(page)))
	} else {
		params.Add("page", "1")
	}
	params.Add("sparkline", strconv.FormatBool(sparkline))
	if len(priceChangePercentage) != 0 {
		price := strings.Join(priceChangePercentage, ",")
		params.Add("price_change_percentage", price)
	}
	if locale != "" {
		params.Add("locale", locale)
	}
	if precision != "" {
		params.Add("precision", precision)
	}

	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, coinsMarketsPath, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to simple token price api", "error", err)
		return nil, err
	}

	var data []CoinsMarketsResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal simple token price response", "error", err)
		return nil, err
	}
	return &data, nil
}

// CoinsID gets current data(name, price, market, including exchange tickers) for a coins.
//
// IMPORTANT:
// Ticker <object> is limited to 100 items, to get more tickers, use /coins/{id}/tickers.
// Ticker <is_stale> is true when ticker that has not been updated/unchanged from the exchange for more than 8 hours.
// Ticker <is_anomaly> is true if ticker's price is outliered by our system.
// You are responsible for managing how you want to display these information (e.g. footnote, different background, change opacity, hide).
//
// Note: to check if a price is stale, please refer to last_updated of the price.
//
// Dictionary:
//
// last: latest unconverted price in the respective pair target currency.
//
// volume: unconverted 24h trading volume in the respective pair target currency.
//
// converted_last: latest converted price in BTC, ETH, and USD.
//
// converted_volume: converted 24h trading volume in BTC, ETH, and USD.
//
// Cache/Update Frequency: every 60 seconds.
//
// Path parameters:
//
// id(required): pass the coin id (can be obtained from /coins) eg. bitcoin.
//
// Query parameters:
//
// localization(optional): include all localized languages in response (true/false) [default: true].
//
// tickers(optional): include tickers data (true/false) [default: true].
//
// market_data(optional): include market_data (true/false) [default: true].
//
// community_data(optional): include community_data data (true/false) [default: true].
//
// developer_data(optional): include developer_data data (true/false) [default: true].
//
// sparkline(optional): include sparkline 7 days data (eg. true, false) [default: false].
func (c *Client) CoinsID(ctx context.Context, id string, localization, tickers, marketData, communityData,
	developerData, sparkline bool) (*CoinsIDResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id should not be empty")
	}

	params := url.Values{}
	params.Add("localization", strconv.FormatBool(localization))
	params.Add("tickers", strconv.FormatBool(tickers))
	params.Add("market_data", strconv.FormatBool(marketData))
	params.Add("community_data", strconv.FormatBool(communityData))
	params.Add("developer_data", strconv.FormatBool(developerData))
	params.Add("sparkline", strconv.FormatBool(sparkline))

	path := fmt.Sprintf(coinsIDPath, id)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to coins id api", "error", err)
		return nil, err
	}

	var data CoinsIDResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coins id response", "error", err)
		return nil, err
	}
	return &data, nil
}

// CoinsIDTickers gets coin tickers (paginated to 100 items).
//
// IMPORTANT:
// Ticker is_stale is true when ticker that has not been updated/unchanged from the exchange for more than 8 hours.
// Ticker is_anomaly is true if ticker's price is outliered by our system.
// You are responsible for managing how you want to display these information (e.g. footnote, different background, change opacity, hide)
//
// Dictionary:
//
// last: latest unconverted price in the respective pair target currency.
//
// volume: unconverted 24h trading volume in the respective pair target currency.
//
// converted_last: latest converted price in BTC, ETH, and USD.
//
// converted_volume: converted 24h trading volume in BTC, ETH, and USD.
//
// Cache/Update Frequency: every 2 minutes.
//
// Path parameters:
//
// id(required): pass the coin id (can be obtained from /coins) eg. bitcoin.
//
// Query parameters:
//
// exchange_ids(optional): filter results by exchange_ids (ref: v3/exchanges/list).
//
// include_exchange_logo(optional): flag to show exchange_logo. valid values: true, false.
//
// page(optional): page through results.
//
// order(optional): valid values: trust_score_desc (default), trust_score_asc and volume_desc.
//
// depth(optional): flag to show 2% orderbook depth. i.e., cost_to_move_up_usd and cost_to_move_down_usd. valid
// values: true, false.
func (c *Client) CoinsIDTickers(ctx context.Context, id, exchangeIDs string, includeExchangeLogo bool, page uint,
	order string, depth bool) (*CoinsIDTickersResponse, int, error) {
	if id == "" {
		return nil, -1, fmt.Errorf("id should not be empty")
	}

	params := url.Values{}
	if exchangeIDs != "" {
		params.Add("exchange_ids", exchangeIDs)
	}
	params.Add("include_exchange_logo", strconv.FormatBool(includeExchangeLogo))
	if page != 0 {
		params.Add("page", strconv.Itoa(int(page)))
	}
	if order != "" {
		params.Add("order", order)
	}
	params.Add("depth", strconv.FormatBool(depth))

	path := fmt.Sprintf(coinsTickersPath, id)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, header, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to coins tickers api", "error", err)
		return nil, -1, err
	}

	total := header.Get("total")
	totalInt, err := strconv.Atoi(total)
	if err != nil {
		slog.Error("failed to parse total http response header", "total", totalInt)
		return nil, -1, err
	}
	pageCount := totalInt/100 + 1

	var data CoinsIDTickersResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coins tickers response", "error", err)
		return nil, -1, err
	}
	return &data, pageCount, nil
}

// CoinsIDHistory gets historical data (price, market cap, 24hr volume, ..) at a given date for a coin.
//
// The data returned is at 00:00:00 UTC.
// The last completed UTC day (00:00) is available 35 minutes after midnight on the next UTC day (00:35).
//
// Path parameters:
//
// id(required): pass the coin id (can be obtained from /coins) eg. bitcoin.
//
// Query parameters:
//
// date(required): the date of data snapshot in dd-mm-yyyy eg. 30-12-2022.
//
// localization(optional): set false to exclude localized languages in response.
func (c *Client) CoinsIDHistory(ctx context.Context, id, date string, localization bool) (*CoinsIDHistoryResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id should not be empty")
	}
	if date == "" {
		return nil, fmt.Errorf("date should not be empty")
	}

	params := url.Values{}
	params.Add("date", date)
	params.Add("localization", strconv.FormatBool(localization))

	path := fmt.Sprintf(coinsHistoryPath, id)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to coins history api", "error", err)
		return nil, err
	}

	var data CoinsIDHistoryResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coins history response", "error", err)
		return nil, err
	}
	return &data, nil
}

// AssetPlatforms lists all asset platforms(Blockchain networks).
//
// Query parameters:
//
// filter(optional): apply relevant filters to results; valid values: "nft" (asset_platform nft-support).
func (c *Client) AssetPlatforms(ctx context.Context, filter string) (*[]AssetPlatformsResponse, error) {
	var params url.Values
	if filter != "" {
		params.Add("filter", filter)
	}

	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", c.apiURL, assetPlatformsPath, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", c.apiURL, assetPlatformsPath)
	}

	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to asset platforms api", "error", err)
		return nil, err
	}

	var data []AssetPlatformsResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal asset platforms response", "error", err)
		return nil, err
	}
	return &data, nil
}
