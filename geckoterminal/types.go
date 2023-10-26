package geckoterminal

import "time"

type basicStruct struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// NetworksItem
type NetworksItem struct {
	basicStruct
	Attributes AttributesItem `json:"attributes"`
}

// AttributesItem
type AttributesItem struct {
	Name                     string  `json:"name"`
	CoingeckoAssetPlatformID *string `json:"coingecko_asset_platform_id,omitempty"`
}

// LinksItem used for pagination
type LinksItem struct {
	First string  `json:"first"`
	Prev  *string `json:"prev"`
	Next  *string `json:"next"`
	Last  string  `json:"last"`
}

// PoolDataItem
type PoolDataItem struct {
	basicStruct
	Attributes    PoolAttributesItem     `json:"attributes"`
	Relationships map[string]basicStruct `json:"relationships"`
}

// PoolAttributesItem
type PoolAttributesItem struct {
	BaseTokenPriceUSD             *string                     `json:"base_token_price_usd"`
	BaseTokenPriceNativeCurrency  *string                     `json:"base_token_price_native_currency"`
	QuotaTokenPriceUSD            *string                     `json:"quota_token_price_usd"`
	QuoteTokenPriceNativeCurrency *string                     `json:"quote_token_price_native_currency"`
	BaseTokenPriceQuoteToken      *string                     `json:"base_token_price_quote_token"`
	QuoteTokenPriceBaseToken      *string                     `json:"quote_token_price_base_token"`
	Address                       string                      `json:"address"`
	Name                          string                      `json:"name"`
	PoolCreatedAt                 *time.Time                  `json:"pool_created_at"`
	FDVUsed                       *string                     `json:"fdv_used"`
	MarketCapUSD                  *string                     `json:"market_cap_usd"`
	PriceChangePercentage         map[string]string           `json:"price_change_percentage"`
	Transactions                  map[string]TransactionsItem `json:"transactions"`
	VolumeUSD                     map[string]string           `json:"volume_usd"`
	ReserveInUSD                  *string                     `json:"reserve_in_usd"`
}

// TransactionsItem
type TransactionsItem struct {
	Buys  int64 `json:"buys"`
	Sells int64 `json:"sells"`
}

// PoolIncludedItem
type PoolIncludedItem struct {
	basicStruct
	Attributes struct {
		Address         string  `json:"address,omitempty"`
		Name            string  `json:"name"`
		Symbol          string  `json:"symbol,omitempty"`
		CoingeckoCoinID *string `json:"coingecko_coin_id,omitempty"`
	} `json:"attributes"`
}
