package coingecko

import (
	"encoding/json"
	"time"
)

type coinsStruct struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// PlatformsItem maps platforms into contract address.
type PlatformsItem map[string]string

// ROIItem describes roi item.
type ROIItem struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}

// SparklineItem describes sparkline item.
type SparklineItem struct {
	Price []float64 `json:"price"`
}

type DetailPlatformsItem map[string]DetailPlatformsInfo

// DetailPlatformsInfo describes detail platforms info
type DetailPlatformsInfo struct {
	DecimalPlace    *uint  `json:"decimal_place"`
	ContractAddress string `json:"contract_address"`
}

// LocalizationItem maps all locale (en, zh...) into respective string.
type LocalizationItem map[string]string

// DescriptionItem maps all description (in locale) into respective string.
type DescriptionItem map[string]string

// LinksItem maps all links
type LinksItem map[string]any

// ImageItem describes image item.
type ImageItem struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large,omitempty"`
}

// ICODataItem describes ico data.
type ICODataItem map[string]any

// AllCurrencies maps all currencies to its price.
type AllCurrencies map[string]float64

// CurrenciesToDate maps all currencies to date.
type CurrenciesToDate map[string]time.Time

// MarketDataItem describes market data items.
type MarketDataItem struct {
	CurrentPrice                           AllCurrencies      `json:"current_price"`
	TotalValueLocked                       map[string]float64 `json:"total_value_locked"`
	McapToTVLRatio                         *float32           `json:"mcap_to_tvl_ratio"`
	FDVToTVLRatio                          *float32           `json:"fdv_to_tvl_ratio"`
	ROI                                    *ROIItem           `json:"roi"`
	Ath                                    AllCurrencies      `json:"ath"`
	AthChangePercentage                    AllCurrencies      `json:"ath_change_percentage"`
	AthDate                                CurrenciesToDate   `json:"ath_date"`
	Atl                                    AllCurrencies      `json:"atl"`
	AtlChangePercentage                    AllCurrencies      `json:"atl_change_percentage"`
	AtlDate                                CurrenciesToDate   `json:"atl_date"`
	MarketCap                              AllCurrencies      `json:"market_cap"`
	MarketCapRank                          int                `json:"market_cap_rank"`
	FullyDilutedValuation                  AllCurrencies      `json:"fully_diluted_valuation"`
	MarketCapFdvRatio                      float32            `json:"market_cap_fdv_ratio"`
	TotalVolume                            AllCurrencies      `json:"total_volume"`
	High24h                                AllCurrencies      `json:"high_24h"`
	Low24h                                 AllCurrencies      `json:"low_24h"`
	PriceChange24h                         float64            `json:"price_change_24h"`
	PriceChangePercentage24h               float64            `json:"price_change_percentage_24h"`
	PriceChangePercentage7d                float64            `json:"price_change_percentage_7d"`
	PriceChangePercentage14d               float64            `json:"price_change_percentage_14d"`
	PriceChangePercentage30d               float64            `json:"price_change_percentage_30d"`
	PriceChangePercentage60d               float64            `json:"price_change_percentage_60d"`
	PriceChangePercentage200d              float64            `json:"price_change_percentage_200d"`
	PriceChangePercentage1y                float64            `json:"price_change_percentage_1y"`
	MarketCapChange24h                     float64            `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h           float64            `json:"market_cap_change_percentage_24h"`
	PriceChange24hInCurrency               AllCurrencies      `json:"price_change_24h_in_currency"`
	PriceChangePercentage1hInCurrency      AllCurrencies      `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24hInCurrency     AllCurrencies      `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7dInCurrency      AllCurrencies      `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage14dInCurrency     AllCurrencies      `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentage60dInCurrency     AllCurrencies      `json:"price_change_percentage_60d_in_currency"`
	PriceChangePercentage30dInCurrency     AllCurrencies      `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage200dInCurrency    AllCurrencies      `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentage1yInCurrency      AllCurrencies      `json:"price_change_percentage_1y_in_currency"`
	MarketCapChange24hInCurrency           AllCurrencies      `json:"market_cap_change_24h_in_currency"`
	MarketCapChangePercentage24hInCurrency AllCurrencies      `json:"market_cap_change_percentage_24h_in_currency"`
	TotalSupply                            float64            `json:"total_supply"`
	MaxSupply                              *float64           `json:"max_supply"`
	CirculatingSupply                      float64            `json:"circulating_supply"`
	SparklineIn7d                          *SparklineItem     `json:"sparkline_in_7d,omitempty"`
	LastUpdated                            time.Time          `json:"last_updated"`
}

// CommunityDataItem describes community data.
type CommunityDataItem struct {
	FacebookLikes            *uint           `json:"facebook_likes"`
	TwitterFollowers         int             `json:"twitter_followers"`
	RedditAveragePosts48h    float64         `json:"reddit_average_posts_48h"`
	RedditAverageComments48h float64         `json:"reddit_average_comments_48h"`
	RedditSubscribers        uint            `json:"reddit_subscribers"`
	RedditAccountsActive48h  json.RawMessage `json:"reddit_accounts_active_48h"`
	TelegramChannelUserCount *uint           `json:"telegram_channel_user_count,omitempty"`
}

// DeveloperDataItem describes developer data.
type DeveloperDataItem struct {
	Forks                        uint `json:"forks"`
	Stars                        uint `json:"stars"`
	Subscribers                  uint `json:"subscribers"`
	TotalIssues                  uint `json:"total_issues"`
	ClosedIssues                 uint `json:"closed_issues"`
	PullRequestsMerged           uint `json:"pull_requests_merged"`
	PullRequestContributors      uint `json:"pull_request_contributors"`
	CodeAdditionsDeletions4Weeks struct {
		Additions *int `json:"additions"`
		Deletions *int `json:"deletions"`
	} `json:"code_additions_deletions_4_weeks"`
	CommitCount4Weeks              uint  `json:"commit_count_4_weeks"`
	Last4WeeksCommitActivitySeries []int `json:"last_4_weeks_commit_activity_series,omitempty"`
}

// PublicInterestStatsItem map all public interest stats.
type PublicInterestStatsItem struct {
	AlexaRank   *uint `json:"alexa_rank"`
	BingMatches *uint `json:"bing_matches"`
}

// TickersItem describes tickers.
type TickersItem struct {
	Base   string `json:"base"`
	Target string `json:"target"`
	Market struct {
		Name                string `json:"name"`
		Identifier          string `json:"identifier"`
		HasTradingIncentive bool   `json:"has_trading_incentive"`
		Logo                string `json:"logo,omitempty"`
	} `json:"market"`
	Last                   float64            `json:"last"`
	Volume                 float64            `json:"volume"`
	CostToMoveUpUsd        float64            `json:"cost_to_move_up_usd,omitempty"`
	CostToMoveDownUsd      float64            `json:"cost_to_move_down_usd,omitempty"`
	ConvertedLast          map[string]float64 `json:"converted_last"`
	ConvertedVolume        map[string]float64 `json:"converted_volume"`
	TrustScore             string             `json:"trust_score"`
	BidAskSpreadPercentage float64            `json:"bid_ask_spread_percentage"`
	Timestamp              time.Time          `json:"timestamp"`
	LastTradedAt           time.Time          `json:"last_traded_at"`
	LastFetchAt            time.Time          `json:"last_fetch_at"`
	IsAnomaly              bool               `json:"is_anomaly"`
	IsStale                bool               `json:"is_stale"`
	TradeURL               string             `json:"trade_url"`
	TokenInfoURL           *string            `json:"token_info_url"`
	CoinID                 string             `json:"coin_id"`
	TargetCoinID           string             `json:"target_coin_id,omitempty"`
}

// MarketDataForHistory used for CoinsHistory API.
type MarketDataForHistory struct {
	CurrentPrice AllCurrencies `json:"current_price"`
	MarketCap    AllCurrencies `json:"market_cap"`
	TotalVolume  AllCurrencies `json:"total_volume"`
}

// ChartItem
type ChartItem [2]float64

// ExchangesItem
type ExchangesItem struct {
	Name                        string  `json:"name"`
	YearEstablished             *int    `json:"year_established"`
	Country                     *string `json:"country"`
	Description                 string  `json:"description"`
	URL                         string  `json:"url"`
	Image                       string  `json:"image"`
	HasTradingIncentive         *bool   `json:"has_trading_incentive"`
	TrustScore                  float32 `json:"trust_score"`
	TrustScoreRank              uint    `json:"trust_score_rank"`
	TradeVolume24hBTC           float64 `json:"trade_volume_24h_btc"`
	TradeVolume24hBTCNormalized float64 `json:"trade_volume_24h_btc_normalized"`
}

// StatusUpdatesItem
type StatusUpdatesItem struct {
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	User        string    `json:"user"`
	UserTitle   string    `json:"user_title"`
	Pin         bool      `json:"pin"`
	Project     struct {
		Type  string    `json:"type"`
		ID    string    `json:"id"`
		Name  string    `json:"name"`
		Image ImageItem `json:"image"`
	} `json:"project"`
}

// DerivativesExchangesTickersItem
type DerivativesExchangesTickersItem struct {
	Symbol               string            `json:"symbol"`
	Base                 string            `json:"base"`
	Target               string            `json:"target"`
	TradeURL             string            `json:"trade_url"`
	ContractType         string            `json:"contract_type"`
	Last                 float64           `json:"last"`
	H24PercentageChange  float64           `json:"h24_percentage_change"`
	Index                float64           `json:"index"`
	IndexBasisPercentage float64           `json:"index_basis_percentage"`
	BidAskSpread         float64           `json:"bid_ask_spread"`
	FundingRate          float64           `json:"funding_rate"`
	OpenInterestUSD      float64           `json:"open_interest_usd"`
	H24Volume            float64           `json:"h24_volume"`
	ConvertedVolume      map[string]string `json:"converted_volume"`
	ConvertedLast        map[string]string `json:"converted_last"`
	LastTraded           int64             `json:"last_traded"`
	ExpiredAt            *int64            `json:"expired_at"`
}

// CompaniesItem
type CompaniesItem struct {
	Name                    string  `json:"name"`
	Symbol                  string  `json:"symbol"`
	Country                 string  `json:"country"`
	TotalHoldings           float64 `json:"total_holdings"`
	TotalEntryValueUSD      float64 `json:"total_entry_value_usd"`
	TotalCurrentValueUSD    float64 `json:"total_current_value_usd"`
	PercentageOfTotalSupply float64 `json:"percentage_of_total_supply"`
}

type SearchTrendingCoinItem struct {
	ID            string `json:"id"`
	CoinID        int    `json:"coin_id"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	MarketCapRank int    `json:"market_cap_rank"`
	ImageItem
	Slug     string  `json:"slug"`
	PriceBTC float64 `json:"price_btc"`
	Score    int     `json:"score"`
}

// SearchTrendingNFTItem
type SearchTrendingNFTItem struct {
	ID                            string  `json:"id"`
	Name                          string  `json:"name"`
	Symbol                        string  `json:"symbol"`
	Thumb                         string  `json:"thumb"`
	NftContractID                 int     `json:"nft_contract_id"`
	NativeCurrencySymbol          string  `json:"native_currency_symbol"`
	FloorPriceInNativeCurrency    float64 `json:"floor_price_in_native_currency"`
	FloorPrice24HPercentageChange float64 `json:"floor_price_24h_percentage_change"`
}

// SearchCoinsItem
type SearchCoinsItem struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	APISymbol     string `json:"api_symbol"`
	Symbol        string `json:"symbol"`
	MarketCapRank *int   `json:"market_cap_rank"`
	Thumb         string `json:"thumb"`
	Large         string `json:"large"`
}

// SearchExchangesItem
type SearchExchangesItem struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MarketType string `json:"market_type"`
	Thumb      string `json:"thumb"`
	Large      string `json:"large"`
}

// SearchCategoriesItem
type SearchCategoriesItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// SearchNFTsItem
type SearchNFTsItem struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Thumb  string `json:"thumb"`
}

// ExchangeRatesItemStruct
type ExchangeRatesItem struct {
	Name  string  `json:"name"`
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

// NativeCurrencyUSDItem used for NFTsIDResponse.
type NativeCurrencyUSDItem struct {
	NativeCurrency float64 `json:"native_currency"`
	USD            float64 `json:"usd"`
}

type ExplorerItem struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

// TopGainerLosersItem used for CoinsTopGainersLosersResponse.
type TopGainerLosersItem struct {
	coinsStruct
	Image         string  `json:"image"`
	MarketCapRank *int64  `json:"market_cap_rank"`
	USD           float64 `json:"usd"`
	USD24hVol     float64 `json:"usd_24h_vol"`
	USD1hChange   float64 `json:"usd_1h_change"`
}

// NFTsIDTickersItem used for NFTsIDTickersResponse.
type NFTsIDTickersItem struct {
	FloorPriceInNativeCurrency float64   `json:"floor_price_in_native_currency"`
	H24VolumeInNativeCurrency  float64   `json:"h24_volume_in_native_currency"`
	NativeCurrency             string    `json:"native_currency"`
	UpdatedAt                  time.Time `json:"updated_at"`
	NFTMarketplaceID           string    `json:"nft_marketplace_id"`
}

// CoinsIDCirculatingSupplyChartItem used for CoinsIDCirculatingSupplyChartResponse.
type CoinsIDCirculatingSupplyChartItem [2]json.Number

// TokensListAllItem used for TokensListAllResponse.
type TokensListAllItem struct {
	ChainID  int    `json:"chainId"`
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	LogoURI  string `json:"logoURI"`
}

// {
//            "chainId": 137,
//            "address": "0x52468c88e8b4f5bcca20a6a7813355637dc5e3ad",
//            "name": "Power Of Deep Ocean",
//            "symbol": "PODO",
//            "decimals": 18,
//            "logoURI": "https://assets.coingecko.com/coins/images/27645/thumb/PODO_TICKER_200.png?1665020330"
//        },
