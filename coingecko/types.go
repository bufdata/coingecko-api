package coingecko

import "time"

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
	Large string `json:"large"`
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
	MarketCapRank                          int16              `json:"market_cap_rank"`
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
	FacebookLikes            *uint   `json:"facebook_likes"`
	TwitterFollowers         int32   `json:"twitter_followers"`
	RedditAveragePosts48h    float64 `json:"reddit_average_posts_48h"`
	RedditAverageComments48h float64 `json:"reddit_average_comments_48h"`
	RedditSubscribers        uint    `json:"reddit_subscribers"`
	RedditAccountsActive48h  float64 `json:"reddit_accounts_active_48h"`
	TelegramChannelUserCount *uint   `json:"telegram_channel_user_count"`
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
	Last4WeeksCommitActivitySeries []int `json:"last_4_weeks_commit_activity_series"`
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
	} `json:"market"`
	Last                   float64            `json:"last"`
	Volume                 float64            `json:"volume"`
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
