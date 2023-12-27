package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
	"strings"

	"github.com/bufdata/coingecko-api/util"
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
// cross-check the price on CoinGecko.com. Learn more about our price methodology https://www.coingecko.com/en/methodology.
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
		return nil, fmt.Errorf("id should not be empty")
	}
	if len(contractAddresses) == 0 {
		return nil, fmt.Errorf("the length of contractAddresses should be greater than 0")
	}
	if len(vsCurrencies) == 0 {
		return nil, fmt.Errorf("the length of vsCurrencies should be greater than 0")
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

// ListCoinsInfo lists all supported coins id, name and symbol(no pagination required).
//
// All the coins that show up on this /coins/list endpoint are Active coins that listed by CoinGecko team on
// CoinGecko.com. If a coin is inactive or deactivated, it will be removed from /coins/list.
//
// Cache/Update Frequency: every 5 minutes.
//
// Query parameters:
//
// include_platform(optional): flag to include platform contract addresses (eg. 0x.... for Ethereum based tokens).
// valid values: true, false.
func (c *Client) ListCoinsInfo(ctx context.Context, includePlatform bool) (*[]ListCoinsInfoResponse, error) {
	params := url.Values{}
	params.Add("include_platform", strconv.FormatBool(includePlatform))
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, coinsListPath, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to list coins info api", "error", err)
		return nil, err
	}

	var data []ListCoinsInfoResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal list coins info response", "error", err)
		return nil, err
	}
	return &data, nil
}

// ListCoinsMarketsData lists all supported coins price, market cap,volume and market related data.
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
// sparkline(optional): include sparkline 7 days data (e.g. true, false). Default value: false.
//
// price_change_percentage(optional): include price change percentage in 1h, 24h, 7d, 14d, 30d, 200d, 1y
// (eg. '1h,24h,7d' comma-separated, invalid values will be discarded).
//
// locale(optional): valid values: ar, bg, cs, da, de, el, en, es, fi, fr, he, hi, hr, hu, id, it, ja, ko, lt, nl, no,
// pl, pt, ro, ru, sk, sl, sv, th, tr, uk, vi, zh, zh-tw. Default value: en.
//
// precision(optional): full or any value between 0-18 to specify decimal place for currency price value.
func (c *Client) ListCoinsMarketsData(ctx context.Context, vsCurrency string, ids []string, category, order string, perPage, page uint,
	sparkline bool, priceChangePercentage []string, locale, precision string) (*[]ListCoinsMarketsDataResponse, error) {
	if vsCurrency == "" {
		return nil, fmt.Errorf("vsCurrency should not be empty")
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
	}
	if page != 0 {
		params.Add("page", strconv.Itoa(int(page)))
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
		slog.Error("failed to send request to list coins market data api", "error", err)
		return nil, err
	}

	var data []ListCoinsMarketsDataResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal list coins market data response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetCoinDataByCoinID gets current data(name, price, market, including exchange tickers) for a coin.
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
// [last]: latest unconverted price in the respective pair target currency.
// [volume]: unconverted 24h trading volume in the respective pair target currency.
// [converted_last]: latest converted price in BTC, ETH, and USD.
// [converted_volume]: converted 24h trading volume in BTC, ETH, and USD.
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
func (c *Client) GetCoinDataByCoinID(ctx context.Context, id string, localization, tickers, marketData, communityData,
	developerData, sparkline bool) (*CoinDataResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("coin id should not be empty")
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
		slog.Error("failed to send request to get coin data api", "error", err)
		return nil, err
	}

	var data CoinDataResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coin data response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetCoinTickersByCoinID gets coin tickers (paginated to 100 items).
//
// IMPORTANT:
// Ticker <is_stale> is true when ticker that has not been updated/unchanged from the exchange for more than 8 hours.
// Ticker <is_anomaly> is true if ticker's price is outliered by our system.
// You are responsible for managing how you want to display these information (e.g. footnote, different background, change opacity, hide)
//
// Dictionary:
//
// [last]: latest unconverted price in the respective pair target currency.
// [volume]: unconverted 24h trading volume in the respective pair target currency.
// [converted_last]: latest converted price in BTC, ETH, and USD.
// [converted_volume]: converted 24h trading volume in BTC, ETH, and USD.
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
func (c *Client) GetCoinTickersByCoinID(ctx context.Context, id, exchangeIDs string, includeExchangeLogo bool, page uint,
	order string, depth bool) (*CoinTickersResponse, int, error) {
	if id == "" {
		return nil, -1, fmt.Errorf("coin id should not be empty")
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
		slog.Error("failed to send request to get coin tickers api", "error", err)
		return nil, -1, err
	}

	total := header.Get(totalHeader)
	totalInt, err := strconv.Atoi(total)
	if err != nil {
		slog.Error("failed to parse total http response header", "total", totalInt)
		return nil, -1, err
	}
	pageCount := util.CalculateTotalPages(totalInt, 100)

	var data CoinTickersResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coin tickers response", "error", err)
		return nil, -1, err
	}
	return &data, pageCount, nil
}

// GetCoinHistoryDataByCoinID gets historical data (price, market cap, 24hr volume, ..) at a given date for a coin.
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
func (c *Client) GetCoinHistoryDataByCoinID(ctx context.Context, id, date string, localization bool) (*CoinHistoryDataResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("coin id should not be empty")
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
		slog.Error("failed to send request to get history data api", "error", err)
		return nil, err
	}

	var data CoinHistoryDataResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coin history data response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetCoinMarketChartByCoinID gets historical market data include price, market cap, and 24h volume (granularity auto).
//
// Data granularity is automatic (cannot be adjusted): 1 day from current time = 5 minute interval data;
// 2-90 days from current time = hourly data; above 90 days from current time = daily data (00:00 UTC).
//
// Cache/Update Frequency: every 5 minutes.
//
// The last completed UTC day (00:00) is available 35 minutes after midnight on the next UTC day (00:35).
//
// Path parameters:
//
// id(required): pass the coin id (can be obtained from /coins) eg. bitcoin.
//
// Query parameters:
//
// id(required): pass the coin id (can be obtained from /coins) eg. bitcoin.
//
// vs_currency(required): the target currency of market data (usd, eur, jpy, etc.).
//
// days(required): data up to number of days ago (eg. 1,14,30,max).
//
// interval(optional): data interval. Possible value: daily.
//
// precision(optional): full or any value between 0-18 to specify decimal place for currency price value.
func (c *Client) GetCoinMarketChartByCoinID(ctx context.Context, id, vsCurrency, days, interval, precision string) (
	*CoinMarketChartDataResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("coin id should not be empty")
	}
	if vsCurrency == "" {
		return nil, fmt.Errorf("vs_currency should not be empty")
	}
	if days == "" {
		return nil, fmt.Errorf("days should not be empty")
	}

	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("days", days)
	if interval != "" {
		params.Add("interval", interval)
	}
	if precision != "" {
		params.Add("precision", precision)
	}

	path := fmt.Sprintf(coinsMarketChartPath, id)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get coin market chart api", "error", err)
		return nil, err
	}

	var data CoinMarketChartDataResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coin market chart data response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetCoinMarketChartRangeByCoinID gets historical market data include price, market cap, and 24h volume (granularity auto).
//
// Data granularity is automatic (cannot be adjusted): 1 day from current time = 5 minute interval data;
// 2-90 days from current time = hourly data; above 90 days from current time = daily data (00:00 UTC).
//
// Cache/Update Frequency: every 5 minutes.
//
// The last completed UTC day (00:00) is available 35 minutes after midnight on the next UTC day (00:35).
//
// Path parameters:
//
// id(required): pass the coin id (can be obtained from /coins) eg. bitcoin.
//
// Query parameters:
//
// id(required): pass the coin id (can be obtained from /coins) eg. bitcoin.
//
// vs_currency(required): the target currency of market data (usd, eur, jpy, etc.).
//
// from(required): from date in UNIX Timestamp (eg. 1392577232).
//
// to(required): to date in UNIX Timestamp (eg. 1422577232).
//
// precision(optional): full or any value between 0-18 to specify decimal place for currency price value.
func (c *Client) GetCoinMarketChartRangeByCoinID(ctx context.Context, id, vsCurrency, from, to, precision string) (
	*CoinMarketChartDataResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("coin id should not be empty")
	}
	if vsCurrency == "" {
		return nil, fmt.Errorf("vs_currency should not be empty")
	}
	if from == "" {
		return nil, fmt.Errorf("from should not be empty")
	}
	if to == "" {
		return nil, fmt.Errorf("to should not be empty")
	}

	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("from", from)
	params.Add("to", to)
	if precision != "" {
		params.Add("precision", precision)
	}

	path := fmt.Sprintf(coinsMarketChartRangePath, id)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get market chart range api", "error", err)
		return nil, err
	}

	var data CoinMarketChartDataResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coin market chart data response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetCoinOHLCByCoinID gets coin's OHLC.
//
// Candle's body - data granularity is automatic (cannot be adjusted for public api users):
// 1-2 days: 30 minutes; 3-30 days: 4 hours; 31 days and beyond: 4 days/.
//
// Daily candle interval parameter is available for paid plan users only (Analyst/Lite/Pro/Enterprise), use
// interval=daily parameter in your request: 'daily' interval: available for 1/7/14/30/90/180 days/.
//
// Cache/Update Frequency: every 30 minutes.
//
// The last completed UTC day (00:00) is available 35 minutes after midnight on the next UTC day (00:35).
//
// Path parameters:
//
// id(required): pass the coin id (can be obtained from /coins) eg. bitcoin.
//
// Query parameters:
//
// vs_currency(required): the target currency of market data (usd, eur, jpy, etc.).
//
// days(required): data up to number of days ago (1/7/14/30/90/180/365/max).
//
// precision(optional): full or any value between 0-18 to specify decimal place for currency price value.
func (c *Client) GetCoinOHLCByCoinID(ctx context.Context, id, vsCurrency, days, precision string) (*[]CoinOHLCResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("coin id should not be empty")
	}
	if vsCurrency == "" {
		return nil, fmt.Errorf("vs_currency should not be empty")
	}
	if days == "" {
		return nil, fmt.Errorf("days should not be empty")
	}

	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("days", days)
	if precision != "" {
		params.Add("precision", precision)
	}

	path := fmt.Sprintf(coinsOHLCPath, id)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get coin ohlc api", "error", err)
		return nil, err
	}

	var data []CoinOHLCResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coin ohlc response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetCoinInfoByContractAddress gets coin info from contract address.
//
// Cache/Update Frequency: every 60 seconds.
//
// Path parameters:
//
// id(required): asset platform (See asset_platforms endpoint for list of options).
//
// contract_address(required): token's contract address.
func (c *Client) GetCoinInfoByContractAddress(ctx context.Context, id, contractAddress string) (*CoinDataResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id should not be empty")
	}
	if contractAddress == "" {
		return nil, fmt.Errorf("contract_address should not be empty")
	}

	path := fmt.Sprintf(coinsContractPath, id, contractAddress)
	endpoint := fmt.Sprintf("%s%s", c.apiURL, path)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get coin info api", "error", err)
		return nil, err
	}

	var data CoinDataResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coin data response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetMarketChartByContractAddress Get historical market data include price, market cap, and 24h volume (granularity auto)
// from a contract address.
//
// Data granularity is automatic (cannot be adjusted): 1 day from current time = 5 minute interval data;
// 2-90 days from current time = hourly data; above 90 days from current time = daily data (00:00 UTC).
//
// Cache/Update Frequency: every 5 minutes.
//
// The last completed UTC day (00:00) is available 35 minutes after midnight on the next UTC day (00:35).
//
// Path parameters:
//
// id(required): asset platform (See asset_platforms endpoint for list of options).
//
// contract_address(required): token's contract address.
//
// Path parameters:
//
// vs_currency(required): the target currency of market data (usd, eur, jpy, etc.).
//
// days(required): data up to number of days ago (eg. 1,14,30,max).
//
// precision(optional): full or any value between 0-18 to specify decimal place for currency price value.
func (c *Client) GetMarketChartByContractAddress(ctx context.Context, id, contractAddress, vsCurrency, days, precision string) (
	*CoinMarketChartDataResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("asset_platform id should not be empty")
	}
	if contractAddress == "" {
		return nil, fmt.Errorf("contract_address should not be empty")
	}
	if vsCurrency == "" {
		return nil, fmt.Errorf("vs_currency should not be empty")
	}
	if days == "" {
		return nil, fmt.Errorf("days should not be empty")
	}

	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("days", days)
	if precision != "" {
		params.Add("precision", precision)
	}

	path := fmt.Sprintf(coinsContractMarketChartPath, id, contractAddress)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get market chart api", "error", err)
		return nil, err
	}

	var data CoinMarketChartDataResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coin market chart data response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetMarketChartRangeByContractAddress Get historical market data include price, market cap, and 24h volume (granularity auto)
// from a contract address.
//
// Data granularity is automatic (cannot be adjusted): 1 day from current time = 5 minute interval data;
// 2-90 days from current time = hourly data; above 90 days from current time = daily data (00:00 UTC).
//
// Cache/Update Frequency: every 5 minutes.
//
// The last completed UTC day (00:00) is available 35 minutes after midnight on the next UTC day (00:35).
//
// Path parameters:
//
// id(required): asset platform (See asset_platforms endpoint for list of options).
//
// contract_address(required): token's contract address.
//
// Path parameters:
//
// vs_currency(required): the target currency of market data (usd, eur, jpy, etc.).
//
// from(required): from date in UNIX Timestamp (eg. 1392577232).
//
// to(required): to date in UNIX Timestamp (eg. 1422577232).
//
// precision(optional): full or any value between 0-18 to specify decimal place for currency price value.
func (c *Client) GetMarketChartRangeByContractAddress(ctx context.Context, id, contractAddress, vsCurrency, from, to,
	precision string) (*CoinMarketChartDataResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("asset_platform id should not be empty")
	}
	if contractAddress == "" {
		return nil, fmt.Errorf("contract_address should not be empty")
	}
	if vsCurrency == "" {
		return nil, fmt.Errorf("vs_currency should not be empty")
	}
	if from == "" {
		return nil, fmt.Errorf("from should not be empty")
	}
	if to == "" {
		return nil, fmt.Errorf("to should not be empty")
	}

	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	params.Add("from", from)
	params.Add("to", to)
	if precision != "" {
		params.Add("precision", precision)
	}

	path := fmt.Sprintf(coinsContractMarketChartRangePath, id, contractAddress)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get market chart range api", "error", err)
		return nil, err
	}

	var data CoinMarketChartDataResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coin market chart data response", "error", err)
		return nil, err
	}
	return &data, nil
}

// ListAllAssetPlatforms lists all asset platforms(Blockchain networks).
//
// Query parameters:
//
// filter(optional): apply relevant filters to results; valid values: "nft" (asset_platform nft-support).
func (c *Client) ListAllAssetPlatforms(ctx context.Context, filter string) (*[]AssetPlatformsResponse, error) {
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
		slog.Error("failed to send request to list all asset platforms api", "error", err)
		return nil, err
	}

	var data []AssetPlatformsResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal asset platforms response", "error", err)
		return nil, err
	}
	return &data, nil
}

// ListAllCategories lists all categories.
//
// Cache/Update Frequency: every 5 minutes.
func (c *Client) ListAllCategories(ctx context.Context) (*[]ListAllCategoriesResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.apiURL, coinsCategoriesListPath)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to list all categories api", "error", err)
		return nil, err
	}

	var data []ListAllCategoriesResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal list all categories response", "error", err)
		return nil, err
	}
	return &data, nil
}

// ListAllCategoriesWithMarketData lists all categories with market data.
//
// Cache/Update Frequency: every 5 minutes.
//
// Query parameters:
//
// order(optional): valid values: market_cap_desc(default), market_cap_asc, name_desc, name_asc,
// market_cap_change_24h_desc, market_cap_change_24h_asc.
func (c *Client) ListAllCategoriesWithMarketData(ctx context.Context, order string) (*[]ListAllCategoriesWithMarketDataResponse, error) {
	params := url.Values{}
	if order == "" {
		params.Add("order", "market_cap_desc")
	}

	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, coinsCategoriesPath, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to list all categories with market data api", "error", err)
		return nil, err
	}

	var data []ListAllCategoriesWithMarketDataResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal list categories with market data response", "error", err)
		return nil, err
	}
	return &data, nil
}

// ListAllExchanges lists all exchanges(active with trading volumes).
//
// Cache/Update Frequency: every 60 seconds.
//
// Query parameters:
//
// per_page(optional): Valid values: 1...250; total results per page. Default value: 100.
//
// page(optional): page through results.
func (c *Client) ListAllExchanges(ctx context.Context, perPage, page uint) (*[]ExchangesResponse, int, error) {
	params := url.Values{}
	if perPage == 0 {
		perPage = 100
	}
	params.Add("per_page", strconv.Itoa(int(perPage)))
	if page == 0 {
		page = 1
	}
	params.Add("page", strconv.Itoa(int(page)))

	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, exchangesPath, params.Encode())
	resp, header, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to list all exchanges api", "error", err)
		return nil, -1, err
	}

	total := header.Get(totalHeader)
	totalInt, err := strconv.Atoi(total)
	if err != nil {
		slog.Error("failed to parse total http response header", "total", totalInt)
		return nil, -1, err
	}
	pageCount := util.CalculateTotalPages(totalInt, int(perPage))

	var data []ExchangesResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal exchanges response", "error", err)
		return nil, -1, err
	}
	return &data, pageCount, nil
}

// ListAllMarketsInfo lists all supported markets id and name(no pagination required).
//
// Use this to obtain all the markets' id in order to make API calls.
//
// Cache/Update Frequency: every 5 minutes.
func (c *Client) ListAllMarketsInfo(ctx context.Context) (*[]ExchangeMarketsInfoResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.apiURL, exchangesListPath)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to list all markets info api", "error", err)
		return nil, err
	}

	var data []ExchangeMarketsInfoResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal exchange markets info response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetExchangeVolumeAndTickersByExchangeID gets exchange volume in BTC and top 100 tickers only.
//
// For derivatives (e.g. bitmex, binance_futures), please use /derivatives/exchange/{id} endpoint.
//
// IMPORTANT:
// Ticker <object> is limited to 100 items, to get more tickers, use /exchanges/{id}/tickers.
// Ticker <is_stale> is true when ticker that has not been updated/unchanged from the exchange for more than 8 hours.
// Ticker <is_anomaly> is true if ticker's price is outliered by our system.
// You are responsible for managing how you want to display these information(e.g. footnote, different background,
// change opacity, hide)
//
// Dictionary:
//
// [last]: latest unconverted price in the respective pair target currency.
// [volume]: unconverted 24h trading volume in the respective pair target currency.
// [converted_last]: latest converted price in BTC, ETH, and USD.
// [converted_volume]: converted 24h trading volume in BTC, ETH, and USD.
//
// Cache/Update Frequency: every 60 seconds.
//
// Path parameters:
//
// id(required): pass the exchange id(can be obtained from /exchanges/list) eg. binance.
func (c *Client) GetExchangeVolumeAndTickersByExchangeID(ctx context.Context, id string) (*ExchangeVolumeAndTickersResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("exchange id should not be empty")
	}

	path := fmt.Sprintf(exchangesIDPath, id)
	endpoint := fmt.Sprintf("%s%s", c.apiURL, path)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get volume and tickers api", "error", err)
		return nil, err
	}

	var data ExchangeVolumeAndTickersResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal exchanges volume and tickers response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetExchangeTickersByExchangeID gets exchange tickers (paginated, 100 tickers per page).
//
// For derivatives (e.g. bitmex, binance_futures), please use /derivatives/exchange/{id} endpoint.
//
// IMPORTANT:
// Ticker <is_stale> is true when ticker that has not been updated/unchanged from the exchange for more than 8 hours.
// Ticker <is_anomaly> is true if ticker's price is outliered by our system.
// You are responsible for managing how you want to display these information(e.g. footnote, different background,
// change opacity, hide)
//
// Dictionary:
//
// [last]: latest unconverted price in the respective pair target currency.
// [volume]: unconverted 24h trading volume in the respective pair target currency.
// [converted_last]: latest converted price in BTC, ETH, and USD.
// [converted_volume]: converted 24h trading volume in BTC, ETH, and USD.
//
// Cache/Update Frequency: every 60 seconds.
//
// Path parameters:
//
// id(required): pass the exchange id(can be obtained from /exchanges/list) eg. binance.
//
// Query parameters:
//
// coin_ids(optional): filter tickers by coin_ids (ref: v3/coins/list).
//
// include_exchange_logo(optional): flag to show exchange_logo. Valid values: true, false.
//
// page(optional): page through results.
//
// depth(optional): flag to show 2% orderbook depth. i.e., cost_to_move_up_usd and cost_to_move_down_usd.
// Valid values: true, false.
//
// order(optional): valid values: trust_score_desc (default), trust_score_asc and volume_desc.
func (c *Client) GetExchangeTickersByExchangeID(ctx context.Context, id, coinIDs string, includeExchangeLogo bool, page uint, depth bool,
	order string) (*ExchangeTickersResponse, int, error) {
	if id == "" {
		return nil, -1, fmt.Errorf("exchange id should not be empty")
	}

	params := url.Values{}
	if coinIDs != "" {
		params.Add("coin_ids", coinIDs)
	}
	params.Add("include_exchange_logo", strconv.FormatBool(includeExchangeLogo))
	if page == 0 {
		params.Add("page", "1")
	} else {
		params.Add("page", strconv.Itoa(int(page)))
	}
	params.Add("depth", strconv.FormatBool(depth))
	if order != "" {
		params.Add("order", order)
	}

	path := fmt.Sprintf(exchangesTickerPath, id)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, header, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get exchange tickers api", "error", err)
		return nil, -1, err
	}

	total := header.Get(totalHeader)
	totalInt, err := strconv.Atoi(total)
	if err != nil {
		slog.Error("failed to parse total http response header", "total", totalInt)
		return nil, -1, err
	}
	pageCount := util.CalculateTotalPages(totalInt, 100)

	var data ExchangeTickersResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal exchange tickers response", "error", err)
		return nil, -1, err
	}
	return &data, pageCount, nil
}

// GetExchangeVolumeChartByExchangeID gets volume_chart data (in BTC) for a given exchange.
//
// Data granularity is automatic(cannot be adjusted): 1 day = 10-minutely; 2-90 days = hourly; 91 days above = daily.
//
// Note: exclusive endpoint is available for paid users to query more than 1 year of historical data.
//
// Cache/Update Frequency: every 60 seconds.
//
// Path parameters:
//
// id(required): pass the exchange id(can be obtained from /exchanges/list) eg. binance.
//
// Query parameters:
//
// days(required): data up to number of days ago (1/7/14/30/90/180/365).
func (c *Client) GetExchangeVolumeChartByExchangeID(ctx context.Context, id string, days uint) (*[]ExchangeVolumeChartResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("exchange id should not be empty")
	}

	params := url.Values{}
	params.Add("days", strconv.Itoa(int(days)))

	path := fmt.Sprintf(exchangesVolumeChartPath, id)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get volume chart api", "error", err)
		return nil, err
	}

	var data []ExchangeVolumeChartResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal exchange volume chart response", "error", err)
		return nil, err
	}
	return &data, nil
}

// ListAllDerivativesTickers lists all derivative tickers.
//
// Note: 'open_interest' and 'volume_24h' data are in USD.
//
// Cache/Update Frequency: every 30 seconds.
//
// Query parameters:
//
// include_tickers(optional): ['all', 'unexpired'] - expired to show unexpired tickers, all to list all tickers;
// defaults to unexpired.
func (c *Client) ListAllDerivativesTickers(ctx context.Context, includeTickers string) (*[]DerivativesTickersResponse, error) {
	params := url.Values{}
	if includeTickers == "" {
		params.Add("include_tickers", "unexpired")
	} else {
		params.Add("include_tickers", includeTickers)
	}

	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, derivativesPath, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to list all derivatives tickers api", "error", err)
		return nil, err
	}

	var data []DerivativesTickersResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal derivatives tickers response", "error", err)
		return nil, err
	}
	return &data, nil
}

// ListAllDerivativesExchanges lists all derivative exchanges.
//
// Cache/Update Frequency: every 30 seconds.
//
// Query parameters:
//
// order(optional): order results using following params name_asc，name_desc，open_interest_btc_asc，
// open_interest_btc_desc，trade_volume_24h_btc_asc，trade_volume_24h_btc_desc.
//
// per_page: total results per page.
//
// page(optional): page through results.
func (c *Client) ListAllDerivativesExchanges(ctx context.Context, order string, perPage, page uint) (
	*[]DerivativesExchangesResponse, int, error) {
	params := url.Values{}
	if order == "" {
		params.Add("order", "open_interest_btc_desc")
	} else {
		params.Add("order", order)
	}
	if perPage == 0 {
		perPage = 50
	}
	params.Add("per_page", strconv.Itoa(int(perPage)))
	if page == 0 {
		page = 1
	}
	params.Add("page", strconv.Itoa(int(page)))

	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, derivativesExchangesPath, params.Encode())
	resp, header, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to list all derivatives exchanges api", "error", err)
		return nil, -1, err
	}

	total := header.Get(totalHeader)
	totalInt, err := strconv.Atoi(total)
	if err != nil {
		slog.Error("failed to parse total http response header", "total", totalInt)
		return nil, -1, err
	}
	pageCount := util.CalculateTotalPages(totalInt, int(perPage))

	var data []DerivativesExchangesResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal derivatives exchanges response", "error", err)
		return nil, -1, err
	}
	return &data, pageCount, nil
}

// ListDerivativesExchangeData shows derivative exchange data.
//
// Dictionary:
//
// [last]: latest unconverted price in the respective pair target currency.
// [volume]: unconverted 24h trading volume in the respective pair target currency.
// [converted_last]: latest converted price in BTC, ETH, and USD.
// [converted_volume]: converted 24h trading volume in BTC, ETH, and USD.
//
// Cache/Update Frequency: every 30 seconds.
//
// Path parameters:
//
// id: pass the exchange id (can be obtained from derivatives/exchanges/list) eg. bitmex.
//
// Query parameters:
//
// include_tickers(optional): ['all', 'unexpired'] - expired to show unexpired tickers, all to list all tickers,
// leave blank to omit tickers data in response.
func (c *Client) ListDerivativesExchangeData(ctx context.Context, id, includeTickers string) (*DerivativesExchangeTickersResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("exchange id should not be empty")
	}

	params := url.Values{}
	if includeTickers != "" {
		params.Add("include_tickers", includeTickers)
	}

	path := fmt.Sprintf(derivativesIDPath, id)
	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", c.apiURL, path)
	}
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to list derivatives exchange data api", "error", err)
		return nil, err
	}

	var data DerivativesExchangeTickersResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal derivatives exchange tickers response", "error", err)
		return nil, err
	}
	return &data, nil
}

// ListAllDerivativeExchangeInfo lists all derivative exchanges name and identifier.
//
// Cache/Update Frequency: every 5 minutes.
func (c *Client) ListAllDerivativeExchangeInfo(ctx context.Context) (*[]DerivativesExchangeInfoResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.apiURL, derivativesListPath)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to list all derivative exchange info api", "error", err)
		return nil, err
	}

	var data []DerivativesExchangeInfoResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal derivatives exchange info response", "error", err)
		return nil, err
	}
	return &data, nil
}

// ListAllNFTInfo lists all supported NFT ids, paginated by 100 items per page, paginated to 100 items.
//
// Use this to obtain all the NFT ids in order to make API calls, paginated to 100 items.
//
// Cache/Update Frequency: every 5 minutes.
//
// Query parameters:
//
// order(optional): valid values: h24_volume_native_asc, h24_volume_native_desc, floor_price_native_asc, floor_price_native_desc,
// market_cap_native_asc, market_cap_native_desc, market_cap_usd_asc, market_cap_usd_desc.
//
// asset_platform_id(optional): the id of the platform issuing tokens(see asset_platforms endpoint for list of options).
//
// per_page(optional): valid values: 1..250; total results per page; example: 100.
//
// page(optional): page through results; example: 1.
func (c *Client) ListAllNFTInfo(ctx context.Context, order, assetPlatformID string, perPage, page uint) (*[]NFTInfoResponse, int, error) {
	params := url.Values{}
	if order != "" {
		params.Add("order", order)
	}
	if assetPlatformID != "" {
		params.Add("asset_platform_id", assetPlatformID)
	}
	if perPage == 0 {
		perPage = 100
	}
	params.Add("per_page", strconv.Itoa(int(perPage)))
	if page == 0 {
		page = 1
	}
	params.Add("page", strconv.Itoa(int(page)))

	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, nftsListPath, params.Encode())
	resp, header, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to list all nft info api", "error", err)
		return nil, -1, err
	}

	total := header.Get(totalHeader)
	totalInt, err := strconv.Atoi(total)
	if err != nil {
		slog.Error("failed to parse total http response header", "total", totalInt)
		return nil, -1, err
	}
	pageCount := util.CalculateTotalPages(totalInt, int(perPage))

	var data []NFTInfoResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal nft info response", "error", err)
		return nil, -1, err
	}
	return &data, pageCount, nil
}

// GetNFTDataByNFTID gets current data (name, price_floor, volume_24h ...) for an NFT collection. native_currency (string) is
// only a representative of the currency.
//
// Cache/Update Frequency: every 60 seconds.
//
// Path parameters:
//
// id(required): id of nft collection (can be obtained from /nfts/list).
func (c *Client) GetNFTDataByNFTID(ctx context.Context, id string) (*NFTDataResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("nft id should not be empty")
	}

	path := fmt.Sprintf(nftsIDPath, id)
	endpoint := fmt.Sprintf("%s%s", c.apiURL, path)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get nft data api", "error", err)
		return nil, err
	}

	var data NFTDataResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal nft data response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetNFTDataByAssetPlatformIDAndContractAddress gets current data (name, price_floor, volume_24h ...) for an NFT collection.
// *Solana NFT & Art Blocks are not supported for this endpoint, please use /nfts/{id} endpoint instead.
//
// Cache/Update Frequency: every 60 seconds.
//
// Path parameters:
//
// asset_platform_id(required): the id of the platform issuing tokens (see asset_platforms endpoint for list of
// options, use filter=nft param).
//
// contract_address(required): the contract_address of the nft collection (/nfts/list for list of nft collection with metadata).
func (c *Client) GetNFTDataByAssetPlatformIDAndContractAddress(ctx context.Context, assetPlatformID, contractAddress string) (
	*NFTDataResponse, error) {
	if assetPlatformID == "" {
		return nil, fmt.Errorf("asset_platform_id should not be empty")
	}
	if contractAddress == "" {
		return nil, fmt.Errorf("contract_address should not be empty")
	}

	path := fmt.Sprintf(nftsContractPath, assetPlatformID, contractAddress)
	endpoint := fmt.Sprintf("%s%s", c.apiURL, path)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get nft data api", "error", err)
		return nil, err
	}

	var data NFTDataResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal nft data response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetExchangeRates gets BTC-to-Currency exchange rates.
//
// Cache/Update Frequency: every 60 seconds.
func (c *Client) GetExchangeRates(ctx context.Context) (*ExchangeRatesResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.apiURL, exchangeRatesPath)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get exchange rates api", "error", err)
		return nil, err
	}

	var data ExchangeRatesResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal exchange rates response", "error", err)
		return nil, err
	}
	return &data, nil
}

// Search for coins, categories and markets listed on CoinGecko ordered by largest Market Cap first.
//
// Cache/Update Frequency: every 15 minutes.
//
// Query parameters:
//
// query(optional): search string.
func (c *Client) Search(ctx context.Context, query string) (*SearchResponse, error) {
	params := url.Values{}
	if query != "" {
		params.Add("query", query)
	}

	var endpoint string
	if len(params) != 0 {
		endpoint = fmt.Sprintf("%s%s?%s", c.apiURL, searchPath, params.Encode())
	} else {
		endpoint = fmt.Sprintf("%s%s", c.apiURL, searchPath)
	}
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to search api", "error", err)
		return nil, err
	}

	var data SearchResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal search response", "error", err)
		return nil, err
	}
	return &data, nil
}

// SearchTrending gets Top-7 trending coins on CoinGecko as searched by users in the last 24 hours(ordered by most
// popular first).
//
// Cache/Update Frequency: every 10 minutes.
func (c *Client) SearchTrending(ctx context.Context) (*SearchTrendingResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.apiURL, trendingPath)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to search trending api", "error", err)
		return nil, err
	}

	var data SearchTrendingResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal search trending response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetGlobalCryptocurrencyData gets cryptocurrency global data.
//
// Cache/Update Frequency: every 10 minutes.
func (c *Client) GetGlobalCryptocurrencyData(ctx context.Context) (*GlobalCryptocurrencyResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.apiURL, globalPath)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get global cryptocurrency data api", "error", err)
		return nil, err
	}

	var data GlobalCryptocurrencyResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal global cryptocurrency response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetGlobalTop100DefiData gets Top 100 Cryptocurrency Global Decentralized Finance(defi) data.
//
// Cache/Update Frequency: every 60 minutes.
func (c *Client) GetGlobalTop100DefiData(ctx context.Context) (*GlobalDefiResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.apiURL, globalDefiPath)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get global top 100 defi data api", "error", err)
		return nil, err
	}

	var data GlobalDefiResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal global defi response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetCompaniesPublicTreasury gets public companies bitcoin or ethereum holdings (Ordered by total holdings descending).
//
// Path parameters:
//
// coin_id(required): bitcoin or ethereum.
func (c *Client) GetCompaniesPublicTreasury(ctx context.Context, coinID string) (*CompaniesPublicTreasuryResponse, error) {
	if coinID == "" {
		return nil, fmt.Errorf("coin id should be empty")
	}

	path := fmt.Sprintf(companiesPath, coinID)
	endpoint := fmt.Sprintf("%s%s", c.apiURL, path)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get companies public treasury api", "error", err)
		return nil, err
	}

	var data CompaniesPublicTreasuryResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal companies public treasury response", "error", err)
		return nil, err
	}
	return &data, nil
}
