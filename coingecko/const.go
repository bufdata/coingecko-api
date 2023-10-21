package coingecko

// define public and pro api endpoint
const (
	publicAPIEndpoint = "https://api.coingecko.com/api/v3"
	proAPIEndpoint    = "https://pro-api.coingecko.com/api/v3"
)

const (
	applicationJSONHeader = "application/json"
	proAPIKeyQueryParam   = "x_cg_pro_api_key"
	proAPIKeyHeader       = "x-cg-pro-api-key"
)

// CoinGecko API path
// both public and paid can use these paths
const (
	// ping
	pingPath = "/ping"

	// simple
	simplePricePath = "/simple/price"
	// the format specifier is id of the platform issuing tokens.
	simpleTokenPricePath      = "/simple/token_price/%s"
	supportedVsCurrenciesPath = "/simple/supported_vs_currencies"

	// coins
	coinsListPath             = "/coins/list"
	coinsMarketsPath          = "/coins/markets"
	coinsIDPath               = "/coins/%s"
	coinsTickersPath          = "/coins/%s/tickers"
	coinsHistoryPath          = "/coins/%s/history"
	coinsMarketChartPath      = "/coins/%s/market_chart"
	coinsMarketChartRangePath = "/coins/%s/market_chart/range"
	coinsOHLCPath             = "/coins/%s/ohlc"

	// contract
	coinsContractPath                 = "/coins/%s/contract/%s"
	coinsContractMarketChartPath      = "/coins/%s/contract/%s/market_chart/"
	coinsContractMarketChartRangePath = "/coins/%s/contract/%s/market_chart/range"

	// asset platforms
	assetPlatformsPath = "/asset_platforms"

	// categories
	coinsCategoriesListPath = "/coins/categories/list"
	coinsCategoriesPath     = "/coins/categories"

	// exchanges
	exchangesPath            = "/exchanges"
	exchangesListPath        = "/exchanges/list"
	exchangesIDPath          = "/exchanges/%s"
	exchangesTickerPath      = "/exchanges/%s/tickers"
	exchangesVolumeChartPath = "/exchanges/%s/volume_chart"

	// derivatives
	derivativesPath          = "/derivatives"
	derivativesExchangesPath = "/derivatives/exchanges"
	derivativesIDPath        = "/derivatives/exchanges/%s"
	derivativesListPath      = "/derivatives/exchanges/list"

	// nfts
	nftsListPath     = "/nfts/list"
	nftsIDPath       = "/nfts/%s"
	nftsContractPath = "/nfts/%s/contract/%s"

	// exchange rates
	exchangeRatesPath = "/exchange_rates"

	// search
	searchPath = "/search"

	// trending
	trendingPath = "/search/trending"

	// global
	globalPath = "/global"
	globalDefi = "/global/decentralized_finance_defi"

	// companies
	companiesPath = "/companies/public_treasury/%s"
)

// paid plan apis
const (
	coinsListNewPath             = "/coins/list/new"
	topGainersLoserPath          = "/coins/top_gainers_losers"
	globalMarketCapChartPath     = "/global/market_cap_chart"
	nftsMarketPath               = "/nfts/markets"
	nftsMarketChartPath          = "/nfts/%s/market_chart"
	nftsContractMarketChartPath  = "/nfts/%s/contract/%s/market_chart"
	nftsTickersPath              = "/nfts/%s/tickers"
	exchangeVolumeChartRangePath = "/exchange/%s/volume_chart/range"
)

// enterprise plan apis
const (
	coinsCirculatingSupplyChartPath      = "/coins/%s/circulating_supply_chart"
	coinsCirculatingSupplyChartRangePath = "/coins/%s/circulating_supply_chart/range"
	tokenListAllPath                     = "/token_lists/%s/all.json"
)
