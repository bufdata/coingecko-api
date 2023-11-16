package coingecko

import (
	"context"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestClient_ListLatest200Coins(t *testing.T) {
	cases := []struct {
		name         string
		server       *httptest.Server
		wantedIsErr  bool
		wantedResult *[]ListLatest200CoinsResponse
		wantedErrStr string
	}{
		{
			name:        "success",
			server:      mockHTTPServer(t, "", `[{"id":"texan","symbol":"texan","name":"Texan","activated_at":1673690316}]`),
			wantedIsErr: false,
			wantedResult: &[]ListLatest200CoinsResponse{{
				coinsStruct: coinsStruct{ID: "texan", Symbol: "texan", Name: "Texan"},
				ActivatedAt: 1673690316,
			}},
			wantedErrStr: "",
		},
		{
			name:         "failed to call api",
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			server:       mockHTTPServer(t, "", invalidJSONString),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: invalidCharacterJSONErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.ListLatest200Coins(context.TODO())
			if tt.wantedIsErr {
				if !strings.Contains(err.Error(), tt.wantedErrStr) {
					t.Fatalf("incorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("error should be nil, got: %v", err)
				}
			}
			if !reflect.DeepEqual(result, tt.wantedResult) {
				t.Fatalf("incorrect response, wanted result: %+v, got result: %+v", tt.wantedResult, result)
			}
		})
	}
}

func TestClient_GetTopGainersLosers(t *testing.T) {
	cases := []struct {
		name         string
		vsCurrency   string
		server       *httptest.Server
		wantedIsErr  bool
		wantedResult *CoinsTopGainersLosersResponse
		wantedErrStr string
	}{
		{
			name:        "success",
			vsCurrency:  "usd",
			server:      mockHTTPServer(t, "", `{"top_gainers":[{"id":"platinx","symbol":"ptx","name":"PlatinX","image":"https://assets.coingecko.com/coins/images/23726/original/logo200x200.png?1645162319","market_cap_rank":null,"usd":0.020361781843315337,"usd_24h_vol":187147.7839535509,"usd_1h_change":96.92603350641804}],"top_losers":[{"id":"nftearth","symbol":"nfte","name":"NFTEarth","image":"https://assets.coingecko.com/coins/images/29116/original/20230223_224134.jpg?1677224110","market_cap_rank":null,"usd":0.013121207474003034,"usd_24h_vol":85887.2358881691,"usd_1h_change":-30.431856675273593}]}`),
			wantedIsErr: false,
			wantedResult: &CoinsTopGainersLosersResponse{
				TopGainers: []TopGainerLosersItem{
					{
						coinsStruct: coinsStruct{
							ID:     "platinx",
							Symbol: "ptx",
							Name:   "PlatinX",
						},
						Image:       "https://assets.coingecko.com/coins/images/23726/original/logo200x200.png?1645162319",
						USD:         0.020361781843315337,
						USD24hVol:   187147.7839535509,
						USD1hChange: 96.92603350641804,
					},
				},
				TopLosers: []TopGainerLosersItem{
					{
						coinsStruct: coinsStruct{
							ID:     "nftearth",
							Symbol: "nfte",
							Name:   "NFTEarth",
						},
						Image:       "https://assets.coingecko.com/coins/images/29116/original/20230223_224134.jpg?1677224110",
						USD:         0.013121207474003034,
						USD24hVol:   85887.2358881691,
						USD1hChange: -30.431856675273593,
					},
				},
			},
			wantedErrStr: "",
		},
		{
			name:         "empty vsCurrency param",
			vsCurrency:   "",
			server:       mockHTTPServer(t, "", ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "vsCurrency should not be empty",
		},
		{
			name:         "failed to call api",
			vsCurrency:   "usd",
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			vsCurrency:   "usd",
			server:       mockHTTPServer(t, "", invalidJSONString),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: invalidCharacterJSONErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.GetTopGainersLosers(context.TODO(), tt.vsCurrency, "24h", "1000")
			if tt.wantedIsErr {
				if !strings.Contains(err.Error(), tt.wantedErrStr) {
					t.Fatalf("incorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("error should be nil, got: %v", err)
				}
			}
			if !reflect.DeepEqual(result, tt.wantedResult) {
				t.Fatalf("incorrect response, wanted result: %+v, got result: %+v", tt.wantedResult, result)
			}
		})
	}
}

func TestClient_GetGlobalMarketCapChartData(t *testing.T) {
	cases := []struct {
		name         string
		days         string
		server       *httptest.Server
		wantedIsErr  bool
		wantedResult *GlobalMarketCapChartResponse
		wantedErrStr string
	}{
		{
			name:        "success",
			days:        "1",
			server:      mockHTTPServer(t, "", `{"market_cap_chart":{"market_cap":[[1666137600000,966408688449.8091],[1666224000000,952627682841.667],[1666310400000,948870473181.127]],"volume":[[1666137600000,54276746009.72646],[1666224000000,51224236590.94324],[1666310400000,52896988973.930435]]}}`),
			wantedIsErr: false,
			wantedResult: &GlobalMarketCapChartResponse{MarketCapChart: MarketCapChartItem{
				MarketCap: []ChartItem{{1666137600000, 966408688449.8091}, {1666224000000, 952627682841.667}, {1666310400000, 948870473181.127}},
				Volume:    []ChartItem{{1666137600000, 54276746009.72646}, {1666224000000, 51224236590.94324}, {1666310400000, 52896988973.930435}},
			}},
			wantedErrStr: "",
		},
		{
			name:         "empty days param",
			days:         "",
			server:       mockHTTPServer(t, "", ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "days should not be empty",
		},
		{
			name:         "failed to call api",
			days:         "1",
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			days:         "1",
			server:       mockHTTPServer(t, "", invalidJSONString),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: invalidCharacterJSONErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.GetGlobalMarketCapChartData(context.TODO(), tt.days, "usd")
			if tt.wantedIsErr {
				if !strings.Contains(err.Error(), tt.wantedErrStr) {
					t.Fatalf("incorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("error should be nil, got: %v", err)
				}
			}
			if !reflect.DeepEqual(result, tt.wantedResult) {
				t.Fatalf("incorrect response, wanted result: %+v, got result: %+v", tt.wantedResult, result)
			}
		})
	}
}

func TestClient_ListAllNFTsMarketsData(t *testing.T) {
	cases := []struct {
		name          string
		server        *httptest.Server
		wantedIsErr   bool
		wantedResult1 *[]NFTsMarketsResponse
		wantedResult2 int
		wantedErrStr  string
	}{
		{
			name:        "success",
			server:      mockHTTPServer(t, "1", `[{"id": "bored-ape-yacht-club","contract_address": "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d","asset_platform_id": "ethereum","name": "Bored Ape Yacht Club","image": {"small": "https://assets.coingecko.com/nft_contracts/images/20/small/bored-ape-yacht-club.png?1629268335"},"description": "The Bored Ape Yacht Club is a collection of 10,000 unique Bored Ape NFTs— unique digital collectibles living on the Ethereum blockchain. Your Bored Ape doubles as your Yacht Club membership card, and grants access to members-only benefits. The ecosystem token for the BAYC project is ApeCoin (https://www.coingecko.com/en/coins/apecoin). If you hold a BAYC NFT, you are allocated 10,094 APE tokens per NFT.","native_currency": "ethereum","floor_price": {"native_currency": 76.98,"usd": 100431},"market_cap": {"native_currency": 769646,"usd": 1004111826},"volume_24h": {"native_currency": 150.05,"usd": 195761},"floor_price_in_usd_24h_percentage_change": 8.6122,"number_of_unique_addresses": 6417.0,"number_of_unique_addresses_24h_percentage_change": 0.03118,"total_supply": 9998.0}]`),
			wantedIsErr: false,
			wantedResult1: &[]NFTsMarketsResponse{
				{
					ID:              "bored-ape-yacht-club",
					ContractAddress: "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d",
					AssetPlatformID: "ethereum",
					Name:            "Bored Ape Yacht Club",
					Image: struct {
						Small string `json:"small"`
					}{Small: "https://assets.coingecko.com/nft_contracts/images/20/small/bored-ape-yacht-club.png?1629268335"},
					Description:    "The Bored Ape Yacht Club is a collection of 10,000 unique Bored Ape NFTs— unique digital collectibles living on the Ethereum blockchain. Your Bored Ape doubles as your Yacht Club membership card, and grants access to members-only benefits. The ecosystem token for the BAYC project is ApeCoin (https://www.coingecko.com/en/coins/apecoin). If you hold a BAYC NFT, you are allocated 10,094 APE tokens per NFT.",
					NativeCurrency: "ethereum",
					FloorPrice: NativeCurrencyUSDItem{
						NativeCurrency: 76.98,
						USD:            100431,
					},
					MarketCap: NativeCurrencyUSDItem{
						NativeCurrency: 769646,
						USD:            1004111826,
					},
					Volume24h: NativeCurrencyUSDItem{
						NativeCurrency: 150.05,
						USD:            195761,
					},
					FloorPriceInUSD24hPercentageChange:         8.6122,
					NumberOfUniqueAddresses:                    6417.0,
					NumberOfUniqueAddresses24hPercentageChange: 0.03118,
					TotalSupply: 9998.0,
				},
			},
			wantedResult2: 1,
			wantedErrStr:  "",
		},
		{
			name:          "failed to call api",
			server:        mockErrorHTTPServer(t, "1"),
			wantedIsErr:   true,
			wantedResult1: nil,
			wantedResult2: -1,
			wantedErrStr:  statusCode400ErrStr,
		},
		{
			name:          "incorrect total header",
			server:        mockHTTPServer(t, "abc", ""),
			wantedIsErr:   true,
			wantedResult1: nil,
			wantedResult2: -1,
			wantedErrStr:  "parsing \"abc\": invalid syntax",
		},
		{
			name:          "failed to unmarshal json",
			server:        mockHTTPServer(t, "1", invalidJSONString),
			wantedIsErr:   true,
			wantedResult1: nil,
			wantedResult2: -1,
			wantedErrStr:  invalidCharacterJSONErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result1, result2, err := client.ListAllNFTsMarketsData(context.TODO(), "", "", 0, 0)
			if tt.wantedIsErr {
				if !strings.Contains(err.Error(), tt.wantedErrStr) {
					t.Fatalf("incorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("error should be nil, got: %v", err)
				}
			}
			if !reflect.DeepEqual(result1, tt.wantedResult1) {
				t.Fatalf("incorrect response, wanted result: %+v, got result: %+v", tt.wantedResult1, result1)
			}
			if result2 != tt.wantedResult2 {
				t.Fatalf("incorrect page count, wanted result: %+v, got result: %+v", tt.wantedResult2, result2)
			}
		})
	}
}

func TestClient_GetMarketChartByNFTID(t *testing.T) {
	cases := []struct {
		name         string
		id           string
		days         string
		server       *httptest.Server
		wantedIsErr  bool
		wantedResult *NFTsIDMarketChartResponse
		wantedErrStr string
	}{
		{
			name:        "success",
			id:          "cryptopunks",
			days:        "1",
			server:      mockHTTPServer(t, "", `{"floor_price_usd": [[1666829111000,103371.67058983991]],"floor_price_native": [[1666829111000,65.95]],"h24_volume_usd": [[1666829111000,1462956.0082642469]],"h24_volume_native": [[1666829111000,933.35]],"market_cap_usd": [[1666829111000,1033509962.5572194]],"market_cap_native": [[1666829111000,659368.1]]}`),
			wantedIsErr: false,
			wantedResult: &NFTsIDMarketChartResponse{
				FloorPriceUSD:    []ChartItem{{1666829111000, 103371.67058983991}},
				FloorPriceNative: []ChartItem{{1666829111000, 65.95}},
				H24VolumeUSD:     []ChartItem{{1666829111000, 1462956.0082642469}},
				H24VolumeNative:  []ChartItem{{1666829111000, 933.35}},
				MarketCapUSD:     []ChartItem{{1666829111000, 1033509962.5572194}},
				MarketCapNative:  []ChartItem{{1666829111000, 659368.1}},
			},
			wantedErrStr: "",
		},
		{
			name:         "empty id path params",
			id:           "",
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "nft id should not be empty",
		},
		{
			name:         "empty contractAddress path params",
			id:           "cryptopunks",
			days:         "",
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "days should not be empty",
		},
		{
			name:         "failed to call api",
			id:           "cryptopunks",
			days:         "1",
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			id:           "cryptopunks",
			days:         "1",
			server:       mockHTTPServer(t, "", invalidJSONString),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: invalidCharacterJSONErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.GetMarketChartByNFTID(context.TODO(), tt.id, tt.days)
			if tt.wantedIsErr {
				if !strings.Contains(err.Error(), tt.wantedErrStr) {
					t.Fatalf("incorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("error should be nil, got: %v", err)
				}
			}
			if !reflect.DeepEqual(result, tt.wantedResult) {
				t.Fatalf("incorrect response, wanted result: %+v, got result: %+v", tt.wantedResult, result)
			}
		})
	}
}

func TestClient_GetMarketChartByNFTContractAddress(t *testing.T) {
	cases := []struct {
		name            string
		assetPlatformID string
		contractAddress string
		server          *httptest.Server
		wantedIsErr     bool
		wantedResult    *NFTsIDMarketChartResponse
		wantedErrStr    string
	}{
		{
			name:            "success",
			assetPlatformID: "ethereum",
			contractAddress: "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d",
			server:          mockHTTPServer(t, "", `{"floor_price_usd": [[1668297912000,72621.54501539188]],"floor_price_native": [[1668297912000,57.99]],"h24_volume_usd": [[1668297912000,1036989.0768312709]],"h24_volume_native": [[1668297912000,828.06]],"market_cap_usd": [[1668297912000,726070207.0638882]],"market_cap_native": [[1668297912000,579784.02]]}`),
			wantedIsErr:     false,
			wantedResult: &NFTsIDMarketChartResponse{
				FloorPriceUSD:    []ChartItem{{1668297912000, 72621.54501539188}},
				FloorPriceNative: []ChartItem{{1668297912000, 57.99}},
				H24VolumeUSD:     []ChartItem{{1668297912000, 1036989.0768312709}},
				H24VolumeNative:  []ChartItem{{1668297912000, 828.06}},
				MarketCapUSD:     []ChartItem{{1668297912000, 726070207.0638882}},
				MarketCapNative:  []ChartItem{{1668297912000, 579784.02}},
			},
			wantedErrStr: "",
		},
		{
			name:            "empty assetPlatformID path params",
			assetPlatformID: "",
			server:          mockErrorHTTPServer(t, ""),
			wantedIsErr:     true,
			wantedResult:    nil,
			wantedErrStr:    "assetPlatformID should not be empty",
		},
		{
			name:            "empty contractAddress path params",
			assetPlatformID: "ethereum",
			server:          mockErrorHTTPServer(t, ""),
			wantedIsErr:     true,
			wantedResult:    nil,
			wantedErrStr:    "contractAddress should not be empty",
		},
		{
			name:            "failed to call api",
			assetPlatformID: "ethereum",
			contractAddress: "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d",
			server:          mockErrorHTTPServer(t, ""),
			wantedIsErr:     true,
			wantedResult:    nil,
			wantedErrStr:    statusCode400ErrStr,
		},
		{
			name:            "failed to unmarshal json",
			assetPlatformID: "ethereum",
			contractAddress: "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d",
			server:          mockHTTPServer(t, "", invalidJSONString),
			wantedIsErr:     true,
			wantedResult:    nil,
			wantedErrStr:    invalidCharacterJSONErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.GetMarketChartByNFTContractAddress(context.TODO(), tt.assetPlatformID, tt.contractAddress, "1")
			if tt.wantedIsErr {
				if !strings.Contains(err.Error(), tt.wantedErrStr) {
					t.Fatalf("incorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("error should be nil, got: %v", err)
				}
			}
			if !reflect.DeepEqual(result, tt.wantedResult) {
				t.Fatalf("incorrect response, wanted result: %+v, got result: %+v", tt.wantedResult, result)
			}
		})
	}
}

func TestClient_GetNFTTickersByNFTID(t *testing.T) {
	cases := []struct {
		name         string
		id           string
		server       *httptest.Server
		wantedIsErr  bool
		wantedResult *NFTTickersResponse
		wantedErrStr string
	}{
		{
			name:        "success",
			id:          "ethereum",
			server:      mockHTTPServer(t, "", `{"tickers": [{"floor_price_in_native_currency": 1.25,"h24_volume_in_native_currency": 205.51,"native_currency": "ethereum","updated_at": "2022-10-28T06:31:42.529Z","nft_marketplace_id": "opensea"}]}`),
			wantedIsErr: false,
			wantedResult: &NFTTickersResponse{Tickers: []NFTsIDTickersItem{
				{
					FloorPriceInNativeCurrency: 1.25,
					H24VolumeInNativeCurrency:  205.51,
					NativeCurrency:             "ethereum",
					UpdatedAt:                  "2022-10-28T06:31:42.529Z",
					NFTMarketplaceID:           "opensea",
				},
			}},
			wantedErrStr: "",
		},
		{
			name:         "empty id path param",
			id:           "",
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "nft id should not be empty",
		},
		{
			name:         "failed to call api",
			id:           "ethereum",
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			id:           "ethereum",
			server:       mockHTTPServer(t, "", invalidJSONString),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: invalidCharacterJSONErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.GetNFTTickersByNFTID(context.TODO(), tt.id)
			if tt.wantedIsErr {
				if !strings.Contains(err.Error(), tt.wantedErrStr) {
					t.Fatalf("incorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("error should be nil, got: %v", err)
				}
			}
			if !reflect.DeepEqual(result, tt.wantedResult) {
				t.Fatalf("incorrect response, wanted result: %+v, got result: %+v", tt.wantedResult, result)
			}
		})
	}
}

func TestClient_GetVolumeChartRangeByExchangeID(t *testing.T) {
	cases := []struct {
		name         string
		id           string
		from         int64
		to           int64
		server       *httptest.Server
		wantedIsErr  bool
		wantedResult *[]ExchangeVolumeChartResponse
		wantedErrStr string
	}{
		{
			name:        "success",
			id:          "binance",
			from:        1673222400,
			to:          1675814400,
			server:      mockHTTPServer(t, "", `[[1672617000000.0,"243150.80176188724454"],[1672703400000.0,"383998.4692726336780669"]]`),
			wantedIsErr: false,
			wantedResult: &[]ExchangeVolumeChartResponse{
				{"1672617000000.0", "243150.80176188724454"},
				{"1672703400000.0", "383998.4692726336780669"},
			},
			wantedErrStr: "",
		},
		{
			name:         "empty exchange id path param",
			id:           "",
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "exchange id should not be empty",
		},
		{
			name:         "incorrect from query param",
			id:           "binance",
			from:         0,
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "from should not be less than or equal to 0",
		},
		{
			name:         "incorrect to query param",
			id:           "binance",
			from:         1673222400,
			to:           0,
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "to should not be less than or equal to 0",
		},
		{
			name:         "failed to call api",
			id:           "binance",
			from:         1673222400,
			to:           1675814400,
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			id:           "binance",
			from:         1673222400,
			to:           1675814400,
			server:       mockHTTPServer(t, "", invalidJSONString),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: invalidCharacterJSONErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.GetVolumeChartRangeByExchangeID(context.TODO(), tt.id, tt.from, tt.to)
			if tt.wantedIsErr {
				if !strings.Contains(err.Error(), tt.wantedErrStr) {
					t.Fatalf("incorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("error should be nil, got: %v", err)
				}
			}
			if !reflect.DeepEqual(result, tt.wantedResult) {
				t.Fatalf("incorrect response, wanted result: %+v, got result: %+v", tt.wantedResult, result)
			}
		})
	}
}
