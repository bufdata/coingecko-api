package geckoterminal

const (
	// GeckoTerminal endpoint
	geckoTerminalAPIEndpoint = "https://api.geckoterminal.com/api/v2"
)

const (
	acceptHeader = "accept"
	jsonHeader   = "application/json;version=20230302"
)

const (
	// networks path
	networksPath = "/networks"

	// dexes path
	dexesPath = "/networks/%s/dexes"

	// pool path
	networksIDPoolsAddressPath = "/networks/%s/pools/%s"
	networksIDPoolsMultiPath   = "/networks/%s/pools/multi/%s"
	networksIDPoolsPath        = "/networks/%s/pools"
	networksIDDexesPoolsPath   = "/networks/%s/dexes/%s/pools"
	networksIDNewPoolsPath     = "/networks/%s/new_pools"
	networksNewPoolsPath       = "/networks/new_pools"
	searchPoolsPath            = "/search/pools"

	// tokens path
	networksIDTokensPoolsPath          = "/networks/%s/tokens/%s/pools"
	networksIDTokensPath               = "/networks/%s/tokens/%s"
	networksIDTokensMultiAddressesPath = "/networks/%s/tokens/multi/%s"
	networksIDTokensInfoPath           = "/networks/%s/tokens/%s/info"
	networksIDPoolsInfoPath            = "/networks/%s/pools/%s/info"
	tokensInfoRecentlyUpdatedPath      = "/tokens/info_recently_updated"

	// ohlcvs path
	networksIDPoolsOHLCVPath = "/networks/%s/pools/%s/ohlcv/%s"
)
