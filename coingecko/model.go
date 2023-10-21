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
	MarketCapRank                       int            `json:"market_cap_rank"`
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
	BlockTimeInMinutes           int64                   `json:"block_time_in_minutes"`
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
	MarketCapRank                int                     `json:"market_cap_rank"`
	CoingeckoRank                int                     `json:"coingecko_rank"`
	CoingeckoScore               float32                 `json:"coingecko_score"`
	DeveloperScore               float32                 `json:"developer_score"`
	CommunityScore               float32                 `json:"community_score"`
	LiquidityScore               float32                 `json:"liquidity_score"`
	PublicInterestScore          float32                 `json:"public_interest_score"`
	MarketData                   *MarketDataItem         `json:"market_data,omitempty"`
	CommunityData                *CommunityDataItem      `json:"community_data,omitempty"`
	DeveloperData                *DeveloperDataItem      `json:"developer_data,omitempty"`
	PublicInterestStats          PublicInterestStatsItem `json:"public_interest_stats"`
	StatusUpdates                StatusUpdatesItem       `json:"status_updates"`
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
type CoinsOHLCResponse [5]float64

// CoinsContractResponse returned by CoinsContract API.
type CoinsContractResponse struct {
	CoinsIDResponse
}

// CoinsContractMarketChartResponse returned by CoinsContractMarketChart and CoinsContractMarketChartRange API.
type CoinsContractMarketChartResponse struct {
	CoinsIDMarketChartResponse
}

// CoinsCategoriesListResponse returned by CoinsCategoriesList API.
type CoinsCategoriesListResponse struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

// CoinsCategoriesResponse returned by CoinsCategories API.
type CoinsCategoriesResponse struct {
	ID                 string     `json:"id"`
	Name               string     `json:"name"`
	MarketCap          *float64   `json:"market_cap"`
	MarketCapChange24h *float64   `json:"market_cap_change_24_h"`
	Content            string     `json:"content"`
	Top3Coins          []string   `json:"top_3_coins"`
	Volume24h          *float64   `json:"volume_24h"`
	UpdatedAt          *time.Time `json:"updated_at"`
}

// ExchangesResponse returned by Exchanges API.
type ExchangesResponse struct {
	ID string `json:"id"`
	ExchangesItem
}

// ExchangesListResponse returned by ExchangesList API.
type ExchangesListResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ExchangesIDResponse returned by ExchangesID API.
type ExchangesIDResponse struct {
	ExchangesItem
	FacebookURL   string              `json:"facebook_url"`
	RedditURL     string              `json:"reddit_url"`
	TelegramURL   string              `json:"telegram_url"`
	SlackURL      string              `json:"slack_url"`
	OtherURL1     string              `json:"other_url_1"`
	OtherURL2     string              `json:"other_url_2"`
	TwitterHandle string              `json:"twitter_handle"`
	Centralized   bool                `json:"centralized"`
	PublicNotice  string              `json:"public_notice"`
	AlertNotice   string              `json:"alert_notice"`
	Tickers       []TickersItem       `json:"tickers"`
	StatusUpdates []StatusUpdatesItem `json:"status_updates"`
}

// ExchangesIDTickersResponse returned by ExchangesIDTickers API.
type ExchangesIDTickersResponse struct {
	Name    string
	Tickers []TickersItem `json:"tickers"`
}

// ExchangesIDVolumeChartResponse returned by ExchangesIDVolumeChart API.
// It consists of two elements: one is time represents unix millisecond(int64); another is volume chart, type is string.
type ExchangesIDVolumeChartResponse [2]json.Number

// DerivativesResponse returned by Derivatives API.
type DerivativesResponse struct {
	Market                   string  `json:"market"`
	Symbol                   string  `json:"symbol"`
	IndexID                  string  `json:"index_id"`
	Price                    string  `json:"price"`
	PricePercentageChange24h float64 `json:"price_percentage_change_24h"`
	ContractType             string  `json:"contract_type"`
	Index                    float64 `json:"index"`
	Basis                    float64 `json:"basis"`
	Spread                   float64 `json:"spread"`
	FundingRate              float64 `json:"funding_rate"`
	OpenInterest             float64 `json:"open_interest"`
	Volume24h                float64 `json:"volume_24h"`
	LastTradedAt             int64   `json:"last_traded_at"`
	ExpiredAt                *int64  `json:"expired_at"`
}

// DerivativesExchangesResponse returned by DerivativesExchanges API.
type DerivativesExchangesResponse struct {
	Name                   string  `json:"name"`
	ID                     string  `json:"id,omitempty"`
	OpenInterestBTC        float64 `json:"open_interest_btc"`
	TradeVolume24hBTC      string  `json:"trade_volume_24h_btc"`
	NumberOfPerpetualPairs int     `json:"number_of_perpetual_pairs"`
	NumberOfFuturesPairs   int     `json:"number_of_futures_pairs"`
	Image                  string  `json:"image"`
	YearEstablished        int     `json:"year_established"`
	Country                *string `json:"country"`
	Description            string  `json:"description"`
	URL                    string  `json:"url"`
}

// DerivativesExchangesIDResponse returned by DerivativesExchangesID API.
type DerivativesExchangesIDResponse struct {
	DerivativesExchangesResponse
	Tickers []DerivativesExchangesTickersItem `json:"tickers,omitempty"`
}

// DerivativesExchangesListResponse returned by DerivativesExchangesList API.
type DerivativesExchangesListResponse struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func calculateTotalPages(totalCount, pageSize int) int {
	return (totalCount + pageSize - 1) / pageSize
}
