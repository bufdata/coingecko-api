package coingecko

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewCoinGecko(t *testing.T) {
	cases := []struct {
		name         string
		apiKey       string
		isPro        bool
		httpClient   *http.Client
		wantedResult string
	}{
		{
			name:         "pro api key",
			apiKey:       "test_api_key",
			isPro:        true,
			httpClient:   nil,
			wantedResult: proAPIEndpoint,
		},
		{
			name:         "public api key",
			apiKey:       "test_api_key",
			isPro:        false,
			httpClient:   nil,
			wantedResult: publicAPIEndpoint,
		},
		{
			name:         "api key is empty and non pro",
			apiKey:       "",
			isPro:        false,
			httpClient:   nil,
			wantedResult: publicAPIEndpoint,
		},
		{
			name:         "api key is empty and pro",
			apiKey:       "",
			isPro:        true,
			httpClient:   nil,
			wantedResult: publicAPIEndpoint,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCoinGecko(tt.apiKey, tt.isPro, tt.httpClient)
			if c.apiURL != tt.wantedResult {
				t.Fatalf("incorrect api url, wanted url: %s, got url: %s", tt.wantedResult, c.apiURL)
			}
		})
	}
}

func Test_checkAPIKey(t *testing.T) {
	c := NewCoinGecko("test", true, nil)
	req := httptest.NewRequest(http.MethodGet, publicAPIEndpoint, nil)
	c.checkAPIKey(req)
	result := req.Header.Get(proAPIKeyHeader)
	if result != "test" {
		t.Fatalf("incorrect http header, wanted header: %s", result)
	}
}
