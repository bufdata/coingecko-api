package coingecko

import (
	"encoding/json"
	"time"
)

// PingResponse returned by Ping API.
type PingResponse struct {
	GeckoSays string `json:"gecko_says"`
}

// SimpleSupportedVSCurrenciesResponse returned by SimpleSupportedVSCurrencies API.
type SimpleSupportedVSCurrenciesResponse []string

// CoinsListResponse returned by CoinsList API.
type CoinsListResponse struct {
	coinsStruct
	Platforms *PlatformsItem `json:"platforms,omitempty"`
}

// AssetPlatformsResponse returned by AssetPlatforms API.
type AssetPlatformsResponse struct {
	ID              string   `json:"id"`
	ChainIdentifier *float64 `json:"chain_identifier"`
	Name            string   `json:"name"`
	Shortname       string   `json:"shortname"`
}

// CoinsMarketsResponse returned by CoinsMarkets API.
type CoinsMarketsResponse struct {
	coinsStruct
	Image                               string         `json:"image"`
	CurrentPrice                        float64        `json:"current_price"`
	MarketCap                           float64        `json:"market_cap"`
	MarketCapRank                       int16          `json:"market_cap_rank"`
	FullyDilutedValuation               float64        `json:"fully_diluted_valuation"`
	TotalVolume                         float64        `json:"total_volume"`
	High24h                             float64        `json:"high_24h"`
	Low24h                              float64        `json:"low_24h"`
	PriceChange24h                      float64        `json:"price_change_24h"`
	PriceChangePercentage24h            float64        `json:"price_change_percentage_24h"`
	MarketCapChange24h                  float64        `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h        float64        `json:"market_cap_change_percentage_24h"`
	CirculatingSupply                   float64        `json:"circulating_supply"`
	TotalSupply                         float64        `json:"total_supply"`
	MaxSupply                           *float64       `json:"max_supply"`
	Ath                                 float64        `json:"ath"`
	AthChangePercentage                 float64        `json:"ath_change_percentage"`
	AthDate                             time.Time      `json:"ath_date"`
	Atl                                 float64        `json:"atl"`
	AtlChangePercentage                 float64        `json:"atl_change_percentage"`
	AtlDate                             time.Time      `json:"atl_date"`
	ROI                                 *ROIItem       `json:"roi"`
	LastUpdated                         time.Time      `json:"last_updated"`
	SparklineIn7d                       *SparklineItem `json:"sparkline_in_7d,omitempty"`
	PriceChangePercentage1hInCurrency   float64        `json:"price_change_percentage_1h_in_currency,omitempty"`
	PriceChangePercentage24hInCurrency  float64        `json:"price_change_percentage_24h_in_currency,omitempty"`
	PriceChangePercentage7dInCurrency   float64        `json:"price_change_percentage_7d_in_currency,omitempty"`
	PriceChangePercentage14dInCurrency  float64        `json:"price_change_percentage_14d_in_currency,omitempty"`
	PriceChangePercentage30dInCurrency  float64        `json:"price_change_percentage_30d_in_currency,omitempty"`
	PriceChangePercentage200dInCurrency float64        `json:"price_change_percentage_200d_in_currency,omitempty"`
	PriceChangePercentage1yInCurrency   float64        `json:"price_change_percentage_1y_in_currency,omitempty"`
}

// CoinsIDResponse returned by CoinsID API.
type CoinsIDResponse struct {
	coinsStruct
	AssetPlatformID              string                  `json:"asset_platform_id"`
	Platforms                    *PlatformsItem          `json:"platforms"`
	DetailPlatforms              *DetailPlatformsItem    `json:"detail_platforms"`
	BlockTimeInMinutes           int32                   `json:"block_time_in_minutes"`
	HashingAlgorithm             *string                 `json:"hashing_algorithm"`
	Categories                   []string                `json:"categories"`
	PreviewListing               bool                    `json:"preview_listing"`
	PublicNotice                 json.RawMessage         `json:"public_notice"`
	AdditionalNotices            json.RawMessage         `json:"additional_notices"`
	Localization                 LocalizationItem        `json:"localization,omitempty"`
	Description                  DescriptionItem         `json:"description"`
	Links                        LinksItem               `json:"links"`
	Image                        ImageItem               `json:"image"`
	CountryOrigin                string                  `json:"country_origin"`
	GenesisDate                  *string                 `json:"genesis_date"`
	ContractAddress              string                  `json:"contract_address,omitempty"`
	SentimentVotesUpPercentage   float64                 `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64                 `json:"sentiment_votes_down_percentage"`
	ICOData                      *ICODataItem            `json:"ico_data,omitempty"`
	WatchlistPortfolioUsers      float64                 `json:"watchlist_portfolio_users"`
	MarketCapRank                int16                   `json:"market_cap_rank"`
	CoingeckoRank                int16                   `json:"coingecko_rank"`
	CoingeckoScore               float32                 `json:"coingecko_score"`
	DeveloperScore               float32                 `json:"developer_score"`
	CommunityScore               float32                 `json:"community_score"`
	LiquidityScore               float32                 `json:"liquidity_score"`
	PublicInterestScore          float32                 `json:"public_interest_score"`
	MarketData                   *MarketDataItem         `json:"market_data,omitempty"`
	CommunityData                *CommunityDataItem      `json:"community_data,omitempty"`
	DeveloperData                *DeveloperDataItem      `json:"developer_data,omitempty"`
	PublicInterestStats          PublicInterestStatsItem `json:"public_interest_stats"`
	StatusUpdates                json.RawMessage         `json:"status_updates"`
	LastUpdated                  time.Time               `json:"last_updated"`
	Tickers                      *[]TickersItem          `json:"tickers,omitempty"`
}

// CoinsIDTickersResponse returned by CoinsIDTickers API.
type CoinsIDTickersResponse struct {
	Name    string        `json:"name"`
	Tickers []TickersItem `json:"tickers"`
}

// CoinsIDHistoryResponse returned by CoinsIDHistory API.
type CoinsIDHistoryResponse struct {
	coinsStruct
	Localization        LocalizationItem        `json:"localization,omitempty"`
	Image               ImageItem               `json:"image"`
	MarketData          MarketDataForHistory    `json:"market_data"`
	CommunityData       *CommunityDataItem      `json:"community_data"`
	DeveloperData       *DeveloperDataItem      `json:"developer_data"`
	PublicInterestStats PublicInterestStatsItem `json:"public_interest_stats"`
}

// CoinsIDMarketChartResponse returned by CoinsIDMarketChart or CoinsIDMarketChartRange API.
type CoinsIDMarketChartResponse struct {
	Prices       []ChartItem `json:"prices"`
	MarketCaps   []ChartItem `json:"market_caps"`
	TotalVolumes []ChartItem `json:"total_volumes"`
}

// CoinsOHLCResponse returned by CoinsOHLC API.
// It consists of five elements: time in unix millisecond(int64), coins opening price(float64), high price(float64),
// low price(float64) and closing price(float64).
type CoinsOHLCResponse [5]json.Number
