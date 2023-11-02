package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
)

// ListLatest200Coins gets the latest 200 coins (id & activated time) that recently listed on CoinGecko.com.
//
// CoinGecko equivalent page: https://www.coingecko.com/en/new-cryptocurrencies
// Tips: if you're looking to get the latest coins listed on CoinGecko, this is the best endpoint to do the job.
//
// Update frequency: 30 sec.
func (c *Client) ListLatest200Coins(ctx context.Context) (*[]CoinsListNewResponse, error) {
	endpoint := fmt.Sprintf("%s%s", c.apiURL, coinsListNewPath)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to list latest 200 coins api", "error", err)
		return nil, err
	}

	var data []CoinsListNewResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coins list new response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetTopGainersLosers gets the top 30 coins with the largest price gain and loss by a specific time duration.
//
// CoinGecko equivalent page: https://www.coingecko.com/en/crypto-gainers-losers.
// Note: only coins with at least $50,000 24hour trading volume will be included.
//
// Update frequency: 5 minutes.
//
// Query parameters:
//
// vs_currency(required): filter result by currency. Valid values: usd, jpy, krw, eur, ...
//
// duration(optional): filter result by the time range, from 1 hour up to 1 year.
// Valid values: 1h, 24h, 7d, 14d, 30d, 60d, 1y. Default value: 24h.
//
// top_coins(optional): filter result by MarketCap ranking (top 300 to 1000), or all coins (including coins that do
// not have MarketCap ranking). Valid values: 300, 500, 1000, all. Default value: 1000.
func (c *Client) GetTopGainersLosers(ctx context.Context, vsCurrency, duration, topCoins string) (*CoinsTopGainersLosersResponse, error) {
	if vsCurrency == "" {
		return nil, fmt.Errorf("vs_currency should not be empty")
	}

	params := url.Values{}
	params.Add("vs_currency", vsCurrency)
	if duration != "" {
		params.Add("duration", duration)
	}
	if topCoins != "" {
		params.Add("top_coins", topCoins)
	}

	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, topGainersLoserPath, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to get top gainers losers api", "error", err)
		return nil, err
	}

	var data CoinsTopGainersLosersResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coins top gainers losers response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetGlobalMarketCapChartData gets historical global market cap and volume data, by number of days away from now.
//
// CoinGecko equivalent page: https://www.coingecko.com/en/global-charts.
//
// Data Granularity (auto): 1 day from now = hourly data; 2 days & above from now = daily data (00:00 UTC).
//
// Update Frequency: 60 minutes.
//
// Query parameters:
//
// days(required): data up to number of days ago. Valid values: any integer e.g. 1, 14, 30 , … or max.
//
// vs_currency(optional): filter result by currency. Valid values: jpy, krw, eur, ...
func (c *Client) GetGlobalMarketCapChartData(ctx context.Context, days, vsCurrency string) (*GlobalMarketCapChartResponse, error) {
	if days == "" {
		return nil, fmt.Errorf("days should not be empty")
	}

	params := url.Values{}
	params.Add("days", days)
	if vsCurrency != "" {
		params.Add("vs_currency", vsCurrency)
	}

	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, globalMarketCapChartPath, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to global market cap chart api", "error", err)
		return nil, err
	}

	var data GlobalMarketCapChartResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal global market cap chart response", "error", err)
		return nil, err
	}
	return &data, nil
}

// ListAllNFTsMarketsData gets the list of all supported NFT floor price, market cap, volume and market related data on CoinGecko.
//
// CoinGecko equivalent page: https://www.coingecko.com/en/nft.
//
// Tips: Other nfts endpoints are also available on our Free API documentation page.
// By default, this endpoint will return 100 results per page and only 1 page.
// To get the number 251-500 NFTs ranked by 24hr volume as seen on CoinGecko NFT page, you may include these
// parameters: per_page=250, page=2 and order=h24_volume_usd_desc
//
// Update Frequency: 5 minutes.
//
// Query parameters:
//
// asset_platform_id(optional): filter result by asset platform (blockchain network).
// Valid values: ethereum, avalanche, polygon-pos, arbitrum-one, optimistic-ethereum, klay-token.
//
// order(optional): sort results by field. Valid values: h24_volume_native_asc, h24_volume_native_desc, h24_volume_usd_asc,
// h24_volume_usd_desc, market_cap_usd_asc, market_cap_usd_desc. Default order: market_cap_usd_desc.
//
// per_page(optional): total results per page. Valid values: 1...250. Default value: 100.
// Max value is 250. You can only get up to 250 results per page.
//
// page(optional): page through results. Valid values: any integer e.g. 1, 2, 10, ... Default value: 1.
func (c *Client) ListAllNFTsMarketsData(ctx context.Context, assetPlatformID, order string, perPage, page uint) (
	*[]NFTsMarketsResponse, int, error) {
	params := url.Values{}
	if assetPlatformID == "" {
		assetPlatformID = "ethereum"
	}
	params.Add("asset_platform_id", assetPlatformID)
	if order == "" {
		order = "market_cap_usd_desc"
	}
	params.Add("order", order)
	if perPage == 0 {
		perPage = 100
	}
	params.Add("per_page", strconv.Itoa(int(perPage)))
	if page == 0 {
		page = 1
	}
	params.Add("page", strconv.Itoa(int(page)))

	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, nftsMarketPath, params.Encode())
	resp, header, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to list all nft markets api", "error", err)
		return nil, -1, err
	}

	total := header.Get("total")
	totalInt, err := strconv.Atoi(total)
	if err != nil {
		slog.Error("failed to parse total http response header", "total", totalInt)
		return nil, -1, err
	}
	pageCount := calculateTotalPages(totalInt, int(perPage))

	var data []NFTsMarketsResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal global market cap chart response", "error", err)
		return nil, -1, err
	}
	return &data, pageCount, nil
}

// GetMarketChartByNFTID gets historical market data of a NFT collection, including floor price, market cap, and 24h volume,
// by number of days away from now.
//
// CoinGecko equivalent page: our NFT price floor chart, e.g. as seen on https://www.coingecko.com/en/nft/bored-ape-yacht-club
//
// Data Granularity (auto): 1-14 days from now = 10-minutely data; 15 days & above from now = daily data (00:00 UTC).
//
// Update Frequency: 5 minutes. The last completed UTC day (00:00) is available 5 minutes after midnight on the next
// UTC day (00:05).
//
// Path parameters:
//
// id(required): id of NFT collection. Valid values: cryptopunks, bored-ape-yacht-club, ...
//
// Query parameters:
//
// days(required): data up to number of days ago. Valid values: any integer e.g. 1, 14, 30 , 90 , … or max.
func (c *Client) GetMarketChartByNFTID(ctx context.Context, id, days string) (*NFTsIDMarketChartResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("nft id should not be empty")
	}
	if days == "" {
		return nil, fmt.Errorf("days should not be empty")
	}

	params := url.Values{}
	params.Add("days", days)

	path := fmt.Sprintf(nftsMarketChartPath, id)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to nfts id market chart api", "error", err)
		return nil, err
	}

	var data NFTsIDMarketChartResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal nfts id market chart response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetMarketChartByNFTContractAddress gets historical market data of a NFT collection using contract address, including floor
// price, market cap, and 24h volume, by number of days away from now.
//
// CoinGecko equivalent page: our NFT price floor chart, e.g. as seen on  https://www.coingecko.com/en/nft/bored-ape-yacht-club
//
// Data Granularity (auto): 1-14 days from now = 10-minutely data; 15 days & above from now = daily data (00:00 UTC).
//
// Update Frequency: 5 minutes. The last completed UTC day (00:00) is available 5 minutes after midnight on the next
// UTC day (00:05).
//
// Tips: this endpoint doesn't support Solana NFT and Art Blocks, please use nfts/{id}/market_chart endpoint instead.
//
// Path parameters:
//
// asset_platform_id(required): id of asset platform (blockchain network) note: this is path to pass the blockchain
// network, not query parameter. Valid values: ethereum, avalanche, polygon-pos, arbitrum-one, optimistic-ethereum, klay-token
//
// contract_address(required): contract address of the NFT collection note: this is path to pass the coin id, not query parameter.
//
// Query parameters:
//
// days(required): data up to number of days ago. Valid values: any integer e.g. 1, 14, 30 , 90 , … or max.
func (c *Client) GetMarketChartByNFTContractAddress(ctx context.Context, assetPlatformID, contractAddress, days string) (
	*NFTsIDMarketChartResponse, error) {
	if assetPlatformID == "" {
		return nil, fmt.Errorf("asset_platform_id should not be empty")
	}
	if contractAddress == "" {
		return nil, fmt.Errorf("contract_address should not be empty")
	}

	params := url.Values{}
	params.Add("days", days)

	path := fmt.Sprintf(nftsContractMarketChartPath, assetPlatformID, contractAddress)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to nfts contract market chart api", "error", err)
		return nil, err
	}

	var data NFTsIDMarketChartResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal nfts contract market chart response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetNFTTickersByNFTID gets the latest floor price and 24h volume of a NFT collection, on each NFT marketplace, e.g. OpenSea
// and Looksrare.
//
// CoinGecko equivalent page: https://www.coingecko.com/en/nft/otherdeed-for-otherside (table)
//
// Update Frequency: 30 seconds.
//
// Path parameters:
//
// id(required): id of NFT collection. Valid values: cryptopunks, bored-ape-yacht-club, ...
func (c *Client) GetNFTTickersByNFTID(ctx context.Context, id string) (*NFTsIDTickersResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("nft id should not be empty")
	}

	path := fmt.Sprintf(nftsTickersPath, id)
	endpoint := fmt.Sprintf("%s%s", c.apiURL, path)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to nfts id tickers api", "error", err)
		return nil, err
	}

	var data NFTsIDTickersResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal nfts id tickers response", "error", err)
		return nil, err
	}
	return &data, nil
}

// GetVolumeChartRangeByExchangeID gets historical volume data (in BTC) of an exchange, by specifying a date range
// (up to 31 days per call).
//
// CoinGecko equivalent page: https://www.coingecko.com/en/exchanges/binance#statistics (Exchange Trade Volume chart)
//
// Update Frequency: 5 minutes.
//
// Note: you can query the full historical volume of an exchange with this endpoint, the data interval will be daily.
// the date range (between 'from' and 'to') has to be within 31 days.
//
// Path parameters:
//
// id(required): this is path to pass the exchange id, not query parameter. Valid values: binance, uniswap_v3, ...
//
// Query parameters:
//
// from(required): from date in UNIX Timestamp. Valid values: UNIX timestamp e.g. 1672531200.
//
// to(required): to date in UNIX Timestamp. Valid values: UNIX timestamp e.g. 1672531200.
func (c *Client) GetVolumeChartRangeByExchangeID(ctx context.Context, id string, from, to int64) (*[]ExchangesIDVolumeChartResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("exchange id should not be empty")
	}
	if from == 0 {
		return nil, fmt.Errorf("from should not be empty")
	}
	if to == 0 {
		return nil, fmt.Errorf("to should not be empty")
	}

	params := url.Values{}
	params.Add("from", strconv.Itoa(int(from)))
	params.Add("to", strconv.Itoa(int(to)))

	path := fmt.Sprintf(exchangeVolumeChartRangePath, id)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to exchanges id volume chart range api", "error", err)
		return nil, err
	}

	var data []ExchangesIDVolumeChartResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal exchanges id volume chart range response", "error", err)
		return nil, err
	}
	return &data, nil
}
