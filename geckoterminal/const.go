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
	getNetworksPath = "/networks"

	// dexes path
	getDexesPath = "/networks/%s/dexes"

	// pool path
	getSpecificPoolPath              = "/networks/%s/pools/%s"
	getMultiPoolsPath                = "/networks/%s/pools/multi/%s"
	getTop20PoolsPath                = "/networks/%s/pools"
	getTop20PoolsOnOneDexPath        = "/networks/%s/dexes/%s/pools"
	getLatest20PoolsOnOneNetworkPath = "/networks/%s/new_pools"
	getLatest20PoolsOnAllNetworkPath = "/networks/new_pools"
	searchPoolsPath                  = "/search/pools"

	// tokens path
	getTop20PoolsForOneTokenPath         = "/networks/%s/tokens/%s/pools"
	getSpecificTokenOnOneNetworkPath     = "/networks/%s/tokens/%s"
	getMultiTokensOnOneNetworkPath       = "/networks/%s/tokens/multi/%s"
	getSpecificTokenInfoOnOneNetworkPath = "/networks/%s/tokens/%s/info"
	getPoolTokensInfoOnOneNetworkPath    = "/networks/%s/pools/%s/info"
	getRecentlyUpdated100TokensInfoPath  = "/tokens/info_recently_updated"

	// ohlcvs path
	getOHLCVPath = "/networks/%s/pools/%s/ohlcv/%s"
)
