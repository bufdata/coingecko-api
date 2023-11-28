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

// ListCoinsInfoResponse returned by ListCoinsInfo API.
type ListCoinsInfoResponse struct {
	coinsStruct
	Platforms *PlatformsItem `json:"platforms,omitempty"`
}

// ListCoinsMarketsDataResponse returned by ListCoinsMarketsData API.
type ListCoinsMarketsDataResponse struct {
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
	AthDate                             string         `json:"ath_date"`
	Atl                                 float64        `json:"atl"`
	AtlChangePercentage                 float64        `json:"atl_change_percentage"`
	AtlDate                             string         `json:"atl_date"`
	ROI                                 *ROIItem       `json:"roi"`
	LastUpdated                         string         `json:"last_updated"`
	SparklineIn7d                       *SparklineItem `json:"sparkline_in_7d,omitempty"`
	PriceChangePercentage1hInCurrency   float64        `json:"price_change_percentage_1h_in_currency,omitempty"`
	PriceChangePercentage24hInCurrency  float64        `json:"price_change_percentage_24h_in_currency,omitempty"`
	PriceChangePercentage7dInCurrency   float64        `json:"price_change_percentage_7d_in_currency,omitempty"`
	PriceChangePercentage14dInCurrency  float64        `json:"price_change_percentage_14d_in_currency,omitempty"`
	PriceChangePercentage30dInCurrency  float64        `json:"price_change_percentage_30d_in_currency,omitempty"`
	PriceChangePercentage200dInCurrency float64        `json:"price_change_percentage_200d_in_currency,omitempty"`
	PriceChangePercentage1yInCurrency   float64        `json:"price_change_percentage_1y_in_currency,omitempty"`
}

// CoinDataResponse returned by GetCoinDataByCoinID API.
type CoinDataResponse struct {
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

// CoinTickersResponse returned by GetCoinTickersByCoinID API.
type CoinTickersResponse struct {
	Name    string        `json:"name"`
	Tickers []TickersItem `json:"tickers"`
}

// CoinHistoryDataResponse returned by GetCoinHistoryDataByCoinID API.
type CoinHistoryDataResponse struct {
	coinsStruct
	Localization        LocalizationItem        `json:"localization,omitempty"`
	Image               ImageItem               `json:"image"`
	MarketData          MarketDataForHistory    `json:"market_data"`
	CommunityData       *CommunityDataItem      `json:"community_data"`
	DeveloperData       *DeveloperDataItem      `json:"developer_data"`
	PublicInterestStats PublicInterestStatsItem `json:"public_interest_stats"`
}

// CoinMarketChartDataResponse returned by GetCoinMarketChartByCoinID or GetCoinMarketChartRangeByCoinID API.
type CoinMarketChartDataResponse struct {
	Prices       []ChartItem `json:"prices"`
	MarketCaps   []ChartItem `json:"market_caps"`
	TotalVolumes []ChartItem `json:"total_volumes"`
}

// AssetPlatformsResponse returned by AssetPlatforms API.
type AssetPlatformsResponse struct {
	ID              string   `json:"id"`
	ChainIdentifier *float64 `json:"chain_identifier"`
	Name            string   `json:"name"`
	Shortname       string   `json:"shortname"`
}

// CoinOHLCResponse returned by GetCoinOHLCByCoinID API.
// It consists of five elements: time in unix millisecond(int64), coins opening price(float64), high price(float64),
// low price(float64) and closing price(float64).
type CoinOHLCResponse [5]float64

// ListAllCategoriesResponse returned by ListAllCategories API.
type ListAllCategoriesResponse struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

// ListAllCategoriesWithMarketDataResponse returned by ListAllCategoriesWithMarketData API.
type ListAllCategoriesWithMarketDataResponse struct {
	ID                 string     `json:"id"`
	Name               string     `json:"name"`
	MarketCap          *float64   `json:"market_cap"`
	MarketCapChange24h *float64   `json:"market_cap_change_24_h"`
	Content            string     `json:"content"`
	Top3Coins          []string   `json:"top_3_coins"`
	Volume24h          *float64   `json:"volume_24h"`
	UpdatedAt          *time.Time `json:"updated_at"`
}

// ExchangesResponse returned by ListAllExchanges API.
type ExchangesResponse struct {
	ID string `json:"id"`
	ExchangesItem
}

// ExchangeMarketsInfoResponse returned by ListAllMarketsInfo API.
type ExchangeMarketsInfoResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ExchangeVolumeAndTickersResponse returned by GetExchangeVolumeAndTickersByExchangeID API.
type ExchangeVolumeAndTickersResponse struct {
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

// ExchangeTickersResponse returned by GetExchangeTickersByExchangeID API.
type ExchangeTickersResponse struct {
	Name    string
	Tickers []TickersItem `json:"tickers"`
}

// ExchangeVolumeChartResponse returned by GetExchangeVolumeChartByExchangeID or GetVolumeChartRangeByExchangeID API.
// It consists of two elements: one is time represents unix millisecond(int64); another is volume chart, type is string.
type ExchangeVolumeChartResponse [2]json.Number

// DerivativesTickersResponse returned by ListAllDerivativesTickers API.
type DerivativesTickersResponse struct {
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

// DerivativesExchangesResponse returned by ListAllDerivativesExchanges API.
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

// DerivativesExchangeTickersResponse returned by ListDerivativesExchangeData API.
type DerivativesExchangeTickersResponse struct {
	DerivativesExchangesResponse
	Tickers []DerivativesExchangesTickersItem `json:"tickers,omitempty"`
}

// DerivativesExchangeInfoResponse returned by ListAllDerivativeExchangeInfo API.
type DerivativesExchangeInfoResponse struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// NFTInfoResponse returned by ListAllNFTInfo API.
type NFTInfoResponse struct {
	ID              string `json:"id"`
	ContractAddress string `json:"contract_address"`
	Name            string `json:"name"`
	AssetPlatformID string `json:"asset_platform_id"`
	Symbol          string `json:"symbol"`
}

// NFTDataResponse returned by NFTsID or GetNFTDataByNFTID API.
// TODO: FloorPrice24hPercentageChange, MarketCap24hPercentageChange, Volume24hPercentageChange,FloorPrice7dPercentageChange,
// FloorPrice14dPercentageChange,FloorPrice30dPercentageChange,FloorPrice1yPercentageChange type is NativeCurrencyUSDItem;
// but in golang cannot unmarshal them into {} when they are empty.
type NFTDataResponse struct {
	ID              string `json:"id"`
	ContractAddress string `json:"contract_address"`
	AssetPlatformID string `json:"asset_platform_id"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Image           struct {
		Small string `json:"small"`
	} `json:"image"`
	Description                                string                `json:"description"`
	NativeCurrency                             string                `json:"native_currency"`
	NativeCurrencySymbol                       string                `json:"native_currency_symbol"`
	FloorPrice                                 NativeCurrencyUSDItem `json:"floor_price"`
	MarketCap                                  NativeCurrencyUSDItem `json:"market_cap"`
	Volume24h                                  NativeCurrencyUSDItem `json:"volume_24h"`
	FloorPriceInUSD24hPercentageChange         float64               `json:"floor_price_in_usd_24h_percentage_change"`
	FloorPrice24hPercentageChange              *map[string]any       `json:"floor_price_24h_percentage_change"`
	MarketCap24hPercentageChange               *map[string]any       `json:"market_cap_24h_percentage_change"`
	Volume24hPercentageChange                  *map[string]any       `json:"volume_24h_percentage_change"`
	NumberOfUniqueAddresses                    float64               `json:"number_of_unique_addresses"`
	NumberOfUniqueAddresses24hPercentageChange float64               `json:"number_of_unique_addresses_24h_percentage_change"`
	VolumeInUSD24hPercentageChange             float64               `json:"volume_in_usd_24h_percentage_change"`
	TotalSupply                                float64               `json:"total_supply"`
	OneDaySales                                float64               `json:"one_day_sales"`
	OneDaySales24hPercentageChange             float64               `json:"one_day_sales_24h_percentage_change"`
	OneDayAverageSalePrice                     float64               `json:"one_day_average_sale_price"`
	OneDayAverageSalePrice24hPercentageChange  float64               `json:"one_day_average_sale_price_24h_percentage_change"`
	Links                                      struct {
		Homepage *string `json:"homepage"`
		Twitter  *string `json:"twitter"`
		Discord  *string `json:"discord"`
	} `json:"links"`
	FloorPrice7dPercentageChange  *map[string]any `json:"floor_price_7d_percentage_change"`
	FloorPrice14dPercentageChange *map[string]any `json:"floor_price_14d_percentage_change"`
	FloorPrice30dPercentageChange *map[string]any `json:"floor_price_30d_percentage_change"`
	FloorPrice60dPercentageChange *map[string]any `json:"floor_price_60d_percentage_change"`
	FloorPrice1yPercentageChange  *map[string]any `json:"floor_price_1y_percentage_change"`
	Explorers                     []ExplorerItem  `json:"explorers"`
}

// ExchangeRatesResponse returned by GetExchangeRates API.
type ExchangeRatesResponse struct {
	Rates map[string]ExchangeRatesItem `json:"rates"`
}

// SearchResponse returned by Search API.
type SearchResponse struct {
	Coins      []SearchCoinsItem      `json:"coins"`
	Exchanges  []SearchExchangesItem  `json:"exchanges"`
	ICOs       []any                  `json:"icos"`
	Categories []SearchCategoriesItem `json:"categories"`
	NFTs       []SearchNFTsItem       `json:"nfts"`
}

// SearchTrendingResponse returned by SearchTrending API.
type SearchTrendingResponse struct {
	Coins []struct {
		SearchTrendingCoinItem `json:"item"`
	} `json:"coins"`
	NFTs      []struct{ SearchTrendingNFTItem }
	Exchanges []any `json:"exchanges"`
}

// GlobalCryptocurrencyResponse returned by GetGlobalCryptocurrencyData API.
type GlobalCryptocurrencyResponse struct {
	Data struct {
		ActiveCryptoCurrencies          int                `json:"active_crypto_currencies"`
		UpcomingICOs                    int                `json:"upcoming_icos"`
		OngoingICOs                     int                `json:"ongoing_icos"`
		EndedICOs                       int                `json:"ended_icos"`
		Markets                         int                `json:"markets"`
		TotalMarketCap                  map[string]float64 `json:"total_market_cap"`
		TotalVolume                     map[string]float64 `json:"total_volume"`
		MarketCapPercentage             map[string]float64 `json:"market_cap_percentage"`
		MarketCapChangePercentage24hUSD float64            `json:"market_cap_change_percentage_24h_usd"`
		UpdatedAt                       int64              `json:"updated_at"`
	} `json:"data"`
}

// GlobalDefiResponse returned by GlobalDefi API.
type GlobalDefiResponse struct {
	Data struct {
		DefiMarketCap        string  `json:"defi_market_cap"`
		EthMarketCap         string  `json:"eth_market_cap"`
		DefiToEthRatio       string  `json:"defi_to_eth_ratio"`
		TradingVolume24h     string  `json:"trading_volume_24h"`
		DefiDominance        string  `json:"defi_dominance"`
		TopCoinName          string  `json:"top_coin_name"`
		TopCoinDefiDominance float64 `json:"top_coin_defi_dominance"`
	} `json:"data"`
}

// CompaniesPublicTreasuryResponse returned by CompaniesPublicTreasury API.
type CompaniesPublicTreasuryResponse struct {
	TotalHoldings      float64         `json:"total_holdings"`
	TotalValueUSD      float64         `json:"total_value_usd"`
	MarketCapDominance float64         `json:"market_cap_dominance"`
	Companies          []CompaniesItem `json:"companies"`
}

// ListLatest200CoinsResponse returned by ListLatest200Coins API.
type ListLatest200CoinsResponse struct {
	coinsStruct
	ActivatedAt int64 `json:"activated_at"`
}

// CoinsTopGainersLosersResponse returned by GetTopGainersLosers API.
type CoinsTopGainersLosersResponse struct {
	TopGainers []TopGainerLosersItem `json:"top_gainers"`
	TopLosers  []TopGainerLosersItem `json:"top_losers"`
}

// GlobalMarketCapChartResponse returned by GetGlobalMarketCapChartData API.
type GlobalMarketCapChartResponse struct {
	MarketCapChart MarketCapChartItem `json:"market_cap_chart"`
}

// NFTsMarketsResponse returned by ListAllNFTsMarketsData API.
type NFTsMarketsResponse struct {
	ID              string `json:"id"`
	ContractAddress string `json:"contract_address"`
	AssetPlatformID string `json:"asset_platform_id"`
	Name            string `json:"name"`
	Image           struct {
		Small string `json:"small"`
	} `json:"image"`
	Description                                string                `json:"description"`
	NativeCurrency                             string                `json:"native_currency"`
	FloorPrice                                 NativeCurrencyUSDItem `json:"floor_price"`
	MarketCap                                  NativeCurrencyUSDItem `json:"market_cap"`
	Volume24h                                  NativeCurrencyUSDItem `json:"volume_24h"`
	FloorPriceInUSD24hPercentageChange         float64               `json:"floor_price_in_usd_24h_percentage_change"`
	NumberOfUniqueAddresses                    float64               `json:"number_of_unique_addresses"`
	NumberOfUniqueAddresses24hPercentageChange float64               `json:"number_of_unique_addresses_24h_percentage_change"`
	TotalSupply                                float64               `json:"total_supply"`
}

// NFTsIDMarketChartResponse returned by GetMarketChartByNFTID or GetMarketChartByNFTContractAddress API.
type NFTsIDMarketChartResponse struct {
	FloorPriceUSD    []ChartItem `json:"floor_price_usd"`
	FloorPriceNative []ChartItem `json:"floor_price_native"`
	H24VolumeUSD     []ChartItem `json:"h24_volume_usd"`
	H24VolumeNative  []ChartItem `json:"h24_volume_native"`
	MarketCapUSD     []ChartItem `json:"market_cap_usd"`
	MarketCapNative  []ChartItem `json:"market_cap_native"`
}

// NFTTickersResponse returned by GetNFTTickersByNFTID API.
type NFTTickersResponse struct {
	Tickers []NFTsIDTickersItem `json:"tickers"`
}

// CoinCirculatingSupplyChartResponse returned by GetCirculatingSupplyChartByCoinID API.
type CoinCirculatingSupplyChartResponse struct {
	CirculatingSupply []CoinsIDCirculatingSupplyChartItem `json:"circulating_supply"`
}

// ListAllTokensResponse returned by ListAllTokensByAssetPlatformID API.
type ListAllTokensResponse struct {
	Name      string              `json:"name"`
	LogoURI   string              `json:"logoURI"`
	Keywords  []string            `json:"keywords"`
	Timestamp string              `json:"timestamp"`
	Tokens    []TokensListAllItem `json:"tokens"`
	Version   struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
}
