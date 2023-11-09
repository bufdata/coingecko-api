package coingecko

import (
	"context"
	"encoding/json"
	"math"
	"net/http"
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
			name: "success",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp := &PingResponse{GeckoSays: "ok"}
				data, _ := json.Marshal(resp)
				_, _ = w.Write(data)
			})),
			wantedIsErr:  false,
			wantedResult: &PingResponse{GeckoSays: "ok"},
			wantedErrStr: "",
		},
		{
			name:         "failed to call api",
			server:       mockErrorHTTPServer(t),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			server:       mockInvalidJSONHTTPServer(t),
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
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp := map[string]map[string]float64{
					"ethereum": {"usd": 1901.02},
				}
				data, _ := json.Marshal(resp)
				_, _ = w.Write(data)
			})),
			wantedIsErr: false,
			wantedResult: &map[string]map[string]float64{
				"ethereum": {"usd": 1901.02},
			},
			wantedErrStr: "",
		},
		{
			name:         "incorrect ids param",
			ids:          nil,
			server:       mockHTTPServer(t),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "the length of ids should be greater than 0",
		},
		{
			name:         "incorrect ids vs_currencies param",
			ids:          []string{"ethereum"},
			vsCurrencies: nil,
			server:       mockHTTPServer(t),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "the length of vsCurrencies should be greater than 0",
		},
		{
			name:         "failed to call api",
			ids:          []string{"ethereum"},
			vsCurrencies: []string{"usd"},
			server:       mockErrorHTTPServer(t),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			ids:          []string{"ethereum"},
			vsCurrencies: []string{"usd"},
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp := map[string]map[string]float64{
					"ethereum": {"usd": math.NaN()},
				}
				data, _ := json.Marshal(resp)
				_, _ = w.Write(data)
			})),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: unexpectedEndJSONInputErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.SimplePrice(context.TODO(), tt.ids, tt.vsCurrencies, "", "", "", "", "")
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
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp := map[string]map[string]float64{
					"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984": {"usd": 5.22},
				}
				data, _ := json.Marshal(resp)
				_, _ = w.Write(data)
			})),
			wantedIsErr: false,
			wantedResult: &map[string]map[string]float64{
				"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984": {"usd": 5.22},
			},
			wantedErrStr: "",
		},
		{
			name:         "incorrect id param",
			id:           "",
			server:       mockHTTPServer(t),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "id should not be empty",
		},
		{
			name:              "incorrect contractAddresses param",
			id:                "ethereum",
			contractAddresses: nil,
			server:            mockHTTPServer(t),
			wantedIsErr:       true,
			wantedResult:      nil,
			wantedErrStr:      "the length of contractAddresses should be greater than 0",
		},
		{
			name:              "incorrect vsCurrencies param",
			id:                "ethereum",
			contractAddresses: []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"},
			vsCurrencies:      nil,
			server:            mockHTTPServer(t),
			wantedIsErr:       true,
			wantedResult:      nil,
			wantedErrStr:      "the length of vsCurrencies should be greater than 0",
		},
		{
			name:              "failed to call api",
			id:                "ethereum",
			contractAddresses: []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"},
			vsCurrencies:      []string{"usd"},
			server:            mockErrorHTTPServer(t),
			wantedIsErr:       true,
			wantedResult:      nil,
			wantedErrStr:      statusCode400ErrStr,
		},
		{
			name:              "failed to unmarshal json",
			id:                "ethereum",
			contractAddresses: []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"},
			vsCurrencies:      []string{"usd"},
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp := map[string]map[string]float64{
					"ethereum": {"usd": math.NaN()},
				}
				data, _ := json.Marshal(resp)
				_, _ = w.Write(data)
			})),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: unexpectedEndJSONInputErrStr,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.SimpleTokenPrice(context.TODO(), tt.id, tt.contractAddresses, tt.vsCurrencies, "", "", "", "", "")
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
			name: "success",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp := &SimpleSupportedVSCurrenciesResponse{"usd", "eur"}
				data, _ := json.Marshal(resp)
				_, _ = w.Write(data)
			})),
			wantedIsErr:  false,
			wantedResult: &SimpleSupportedVSCurrenciesResponse{"usd", "eur"},
			wantedErrStr: "",
		},
		{
			name:         "failed to call api",
			server:       mockErrorHTTPServer(t),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: statusCode400ErrStr,
		},
		{
			name:         "failed to unmarshal json",
			server:       mockInvalidJSONHTTPServer(t),
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
