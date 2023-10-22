package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
)

// CoinsIDCirculatingSupplyChart gets historical circulating supply of a coin, by number of days away from now.
//
// CoinGecko equivalent page: https://www.coingecko.com/en/exchanges/binance#statistics (Exchange Trade Volume chart)
//
// Data Granularity (auto): 1 day from now = 5-minutely data. 2-90 days from now = hourly data. 91 days & above
// from now = daily data (00:00 UTC).
//
// Data Availability: from 22 June 2019.
//
// Update Frequency: 5 minutes. The last completed UTC day (00:00) is available 35 minutes after midnight on
// the next UTC day (00:35).
//
// Path parameters:
//
// id(required): this is path to pass the coin id, not query parameter. Valid values: any coin id, e.g. bitcoin ...
//
// Query parameters:
//
// days(required): data up to number of days ago. Valid values: any integer, e.g. 1, 14, 30 ...
//
// interval(optional): data interval. Valid values: daily. if interval is not specified, auto data granularity will apply.
func (c *Client) CoinsIDCirculatingSupplyChart(ctx context.Context, id string, days int, interval string) (*CoinsIDCirculatingSupplyChartResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id should not be empty")
	}

	params := url.Values{}
	params.Add("days", strconv.Itoa(days))
	if interval != "" {
		params.Add("interval", interval)
	}

	path := fmt.Sprintf(coinsCirculatingSupplyChartPath, id)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to coins id circulating supply chart api", "error", err)
		return nil, err
	}

	var data CoinsIDCirculatingSupplyChartResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coins id circulating supply chart response", "error", err)
		return nil, err
	}
	return &data, nil
}

// CoinsIDCirculatingSupplyChartRange gets historical circulating supply of a coin, within a range of timestamp.
//
// CoinGecko equivalent page: https://www.coingecko.com/en/exchanges/binance#statistics (Exchange Trade Volume chart)
//
// Data Granularity (auto): 1 day from now = 5-minutely data. 2-90 days from now = hourly data. 91 days & above
// from now = daily data (00:00 UTC).
//
// Data Availability: from 22 June 2019.
//
// Update Frequency: 5 minutes. The last completed UTC day (00:00) is available 35 minutes after midnight on
// the next UTC day (00:35).
//
// Path parameters:
//
// id(required): this is path to pass the coin id, not query parameter. Valid values: any coin id, e.g. bitcoin ...
//
// Query parameters:
//
// from(required): from date in UNIX Timestamp. Valid values: UNIX timestamp, e.g. 1633046400.
//
// to(required): to date in UNIX Timestamp. Valid values: UNIX timestamp, e.g. 1635724799.
func (c *Client) CoinsIDCirculatingSupplyChartRange(ctx context.Context, id string, from, to int64) (*CoinsIDCirculatingSupplyChartResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id should not be empty")
	}

	params := url.Values{}
	params.Add("from", strconv.Itoa(int(from)))
	params.Add("to", strconv.Itoa(int(to)))

	path := fmt.Sprintf(coinsCirculatingSupplyChartRangePath, id)
	endpoint := fmt.Sprintf("%s%s?%s", c.apiURL, path, params.Encode())
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to coins id circulating supply chart range api", "error", err)
		return nil, err
	}

	var data CoinsIDCirculatingSupplyChartResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal coins id circulating supply chart range response", "error", err)
		return nil, err
	}
	return &data, nil
}

// TokensListAll gets full list of tokens of a blockchain network (asset platform) that is supported by Ethereum
// token list standard(https://tokenlists.org/).
//
// Update Frequency: 5 minutes.
//
// Note: a token will only be included in the list if the contract address is added by CoinGecko team.
//
// Path parameters:
//
// asset_platform_id(required): this is path to pass the asset platform id, not query parameter. Valid values: any
// asset platform id, e.g. polygon-pos , arbitrum-nova, ethereum ...
func (c *Client) TokensListAll(ctx context.Context, assetPlatformID string) (*TokensListAllResponse, error) {
	if assetPlatformID == "" {
		return nil, fmt.Errorf("asset_platform_id should not be empty")
	}

	path := fmt.Sprintf(tokenListAllPath, assetPlatformID)
	endpoint := fmt.Sprintf("%s%s", c.apiURL, path)
	resp, _, err := c.sendReq(ctx, endpoint)
	if err != nil {
		slog.Error("failed to send request to tokens list all api", "error", err)
		return nil, err
	}

	var data TokensListAllResponse
	if err = json.Unmarshal(resp, &data); err != nil {
		slog.Error("failed to unmarshal tokens list all response", "error", err)
		return nil, err
	}
	return &data, nil
}
