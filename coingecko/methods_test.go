package coingecko

import (
	"context"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestClient_Ping(t *testing.T) {
	cases := []struct {
		name         string
		server       *httptest.Server
		wantedIsErr  bool
		wantedResult *PingResponse
		wantedErrStr string
	}{
		{
			name:         "success",
			server:       mockHTTPServer(t, "", `{"gecko_says":"(V3) To the Moon!"}`),
			wantedIsErr:  false,
			wantedResult: &PingResponse{GeckoSays: "(V3) To the Moon!"},
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
			result, err := client.Ping(context.TODO())
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

func TestClient_SimplePrice(t *testing.T) {
	cases := []struct {
		name         string
		ids          []string
		vsCurrencies []string
		server       *httptest.Server
		wantedIsErr  bool
		wantedResult *map[string]map[string]float64
		wantedErrStr string
	}{
		{
			name:         "success",
			ids:          []string{"ethereum"},
			vsCurrencies: []string{"usd"},
			server:       mockHTTPServer(t, "", `{"ethereum": {"usd": 2055.6988786308198,"usd_market_cap": 246511850975.8151,"usd_24h_vol": 23563719178.773373,"usd_24h_change": 1.8256318228221318,"last_updated_at": 1700138165}}`),
			wantedIsErr:  false,
			wantedResult: &map[string]map[string]float64{
				"ethereum": {
					"usd":             2055.6988786308198,
					"usd_market_cap":  246511850975.8151,
					"usd_24h_vol":     23563719178.773373,
					"usd_24h_change":  1.8256318228221318,
					"last_updated_at": 1700138165,
				},
			},
			wantedErrStr: "",
		},
		{
			name:         "empty ids param",
			ids:          nil,
			server:       mockHTTPServer(t, "", ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "the length of ids should be greater than 0",
		},
		{
			name:         "empty vsCurrencies param",
			ids:          []string{"ethereum"},
			vsCurrencies: nil,
			server:       mockHTTPServer(t, "", ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "the length of vsCurrencies should be greater than 0",
		},
		{
			name:         "failed to call api",
			ids:          []string{"ethereum"},
			vsCurrencies: []string{"usd"},
			server:       mockErrorHTTPServer(t, ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			ids:          []string{"ethereum"},
			vsCurrencies: []string{"usd"},
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
			result, err := client.SimplePrice(context.TODO(), tt.ids, tt.vsCurrencies, "true", "true", "true", "true", "true")
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

func TestClient_SimpleTokenPrice(t *testing.T) {
	cases := []struct {
		name              string
		id                string
		contractAddresses []string
		vsCurrencies      []string
		server            *httptest.Server
		wantedIsErr       bool
		wantedResult      *map[string]map[string]float64
		wantedErrStr      string
	}{
		{
			name:              "success",
			id:                "ethereum",
			contractAddresses: []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"},
			vsCurrencies:      []string{"usd"},
			server:            mockHTTPServer(t, "", `{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984": {"usd": 5.369703752217275,"usd_market_cap": 4048630216.552925,"usd_24h_vol": 187988702.82637835,"usd_24h_change": 2.1612574448635384,"last_updated_at": 1700141164}}`),
			wantedIsErr:       false,
			wantedResult: &map[string]map[string]float64{
				"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984": {
					"usd":             5.369703752217275,
					"usd_market_cap":  4048630216.552925,
					"usd_24h_vol":     187988702.82637835,
					"usd_24h_change":  2.1612574448635384,
					"last_updated_at": 1700141164,
				},
			},
			wantedErrStr: "",
		},
		{
			name:         "empty id param",
			id:           "",
			server:       mockHTTPServer(t, "", ""),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "id should not be empty",
		},
		{
			name:              "empty contractAddresses param",
			id:                "ethereum",
			contractAddresses: nil,
			server:            mockHTTPServer(t, "", ""),
			wantedIsErr:       true,
			wantedResult:      nil,
			wantedErrStr:      "the length of contractAddresses should be greater than 0",
		},
		{
			name:              "empty vsCurrencies param",
			id:                "ethereum",
			contractAddresses: []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"},
			vsCurrencies:      nil,
			server:            mockHTTPServer(t, "", ""),
			wantedIsErr:       true,
			wantedResult:      nil,
			wantedErrStr:      "the length of vsCurrencies should be greater than 0",
		},
		{
			name:              "failed to call api",
			id:                "ethereum",
			contractAddresses: []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"},
			vsCurrencies:      []string{"usd"},
			server:            mockErrorHTTPServer(t, ""),
			wantedIsErr:       true,
			wantedResult:      nil,
			wantedErrStr:      statusCode400ErrStr,
		},
		{
			name:              "failed to unmarshal json",
			id:                "ethereum",
			contractAddresses: []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"},
			vsCurrencies:      []string{"usd"},
			server:            mockHTTPServer(t, "", invalidJSONString),
			wantedIsErr:       true,
			wantedResult:      nil,
			wantedErrStr:      invalidCharacterJSONErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.SimpleTokenPrice(context.TODO(), tt.id, tt.contractAddresses, tt.vsCurrencies, "true", "true", "true", "true", "full")
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

func TestClient_SimpleSupportedVSCurrencies(t *testing.T) {
	cases := []struct {
		name         string
		server       *httptest.Server
		wantedIsErr  bool
		wantedResult *SimpleSupportedVSCurrenciesResponse
		wantedErrStr string
	}{
		{
			name:         "success",
			server:       mockHTTPServer(t, "", `["btc","eth","ltc","bch","bnb","eos","xrp","xlm","link","dot","yfi","usd","aed"]`),
			wantedIsErr:  false,
			wantedResult: &SimpleSupportedVSCurrenciesResponse{"btc", "eth", "ltc", "bch", "bnb", "eos", "xrp", "xlm", "link", "dot", "yfi", "usd", "aed"},
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
			result, err := client.SimpleSupportedVSCurrencies(context.TODO())
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

func TestClient_ListCoinsInfo(t *testing.T) {
	cases := []struct {
		name         string
		server       *httptest.Server
		wantedIsErr  bool
		wantedResult *[]ListCoinsInfoResponse
		wantedErrStr string
	}{
		{
			name:        "success",
			server:      mockHTTPServer(t, "", `[{"id": "ethereum","symbol": "eth","name": "Ethereum","platforms": {}},{"id": "uniswap","symbol": "uni","name": "Uniswap","platforms": {"ethereum": "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984","binance-smart-chain": "0xbf5140a22578168fd562dccf235e5d43a02ce9b1","arbitrum-one": "0xfa7f8980b0f1e64a2062791cc3b0871572f1f7f0","optimistic-ethereum": "0x6fd9d7ad17242c41f7131d257212c54a0e816691","xdai": "0x4537e328bf7e4efa29d05caea260d7fe26af9d74","near-protocol": "1f9840a85d5af5bf1d1762f925bdaddc4201f984.factory.bridge.near","energi": "0x665b3a802979ec24e076c80025bff33c18eb6007","sora": "0x009be848df92a400da2f217256c88d1a9b1a0304f9b3e90991a67418e1d3b08c","harmony-shard-0": "0x90d81749da8867962c760414c1c25ec926e889b6","avalanche": "0x8ebaf22b6f053dffeaf46f4dd9efa95d89ba8580","huobi-token": "0x22c54ce8321a4015740ee1109d9cbc25815c46e6","polygon-pos": "0xb33eaad8d922b1083446dc23f610c2567fb5180f"}}]`),
			wantedIsErr: false,
			wantedResult: &[]ListCoinsInfoResponse{
				{
					coinsStruct{
						ID:     "ethereum",
						Symbol: "eth",
						Name:   "Ethereum",
					},
					&PlatformsItem{},
				},
				{
					coinsStruct{
						ID:     "uniswap",
						Symbol: "uni",
						Name:   "Uniswap",
					},
					&PlatformsItem{
						"ethereum":            "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
						"binance-smart-chain": "0xbf5140a22578168fd562dccf235e5d43a02ce9b1",
						"arbitrum-one":        "0xfa7f8980b0f1e64a2062791cc3b0871572f1f7f0",
						"optimistic-ethereum": "0x6fd9d7ad17242c41f7131d257212c54a0e816691",
						"xdai":                "0x4537e328bf7e4efa29d05caea260d7fe26af9d74",
						"near-protocol":       "1f9840a85d5af5bf1d1762f925bdaddc4201f984.factory.bridge.near",
						"energi":              "0x665b3a802979ec24e076c80025bff33c18eb6007",
						"sora":                "0x009be848df92a400da2f217256c88d1a9b1a0304f9b3e90991a67418e1d3b08c",
						"harmony-shard-0":     "0x90d81749da8867962c760414c1c25ec926e889b6",
						"avalanche":           "0x8ebaf22b6f053dffeaf46f4dd9efa95d89ba8580",
						"huobi-token":         "0x22c54ce8321a4015740ee1109d9cbc25815c46e6",
						"polygon-pos":         "0xb33eaad8d922b1083446dc23f610c2567fb5180f",
					},
				},
			},
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
			result, err := client.ListCoinsInfo(context.TODO(), true)
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

// func TestClient_ListCoinsMarketsData(t *testing.T) {
// 	cases := []struct {
// 		name         string
// 		server       *httptest.Server
// 		wantedIsErr  bool
// 		wantedResult *[]ListCoinsMarketsDataResponse
// 		wantedErrStr string
// 	}{
// 		{
// 			name:         "success",
// 			server:       nil,
// 			wantedIsErr:  false,
// 			wantedResult: nil,
// 			wantedErrStr: "",
// 		},
// 		{
// 			name:         "failed to call api",
// 			server:       mockErrorHTTPServer(t, ""),
// 			wantedIsErr:  true,
// 			wantedResult: nil,
// 			wantedErrStr: statusCode400ErrStr,
// 		},
// 		{
// 			name:         "failed to unmarshal json",
// 			server:       mockHTTPServer(t, "", invalidJSONString),
// 			wantedIsErr:  true,
// 			wantedResult: nil,
// 			wantedErrStr: invalidCharacterJSONErrStr,
// 		},
// 	}
// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			client := setup(t)
// 			client.apiURL = tt.server.URL
// 			result, err := client.ListCoinsMarketsData(context.TODO(), false)
// 			if tt.wantedIsErr {
// 				if !strings.Contains(err.Error(), tt.wantedErrStr) {
// 					t.Fatalf("incorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
// 				}
// 			} else {
// 				if err != nil {
// 					t.Fatalf("error should be nil, got: %v", err)
// 				}
// 			}
// 			if !reflect.DeepEqual(result, tt.wantedResult) {
// 				t.Fatalf("incorrect response, wanted result: %+v, got result: %+v", tt.wantedResult, result)
// 			}
// 		})
// 	}
// }
