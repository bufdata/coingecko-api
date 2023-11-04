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
			name:         "api key is empty",
			apiKey:       "",
			isPro:        false,
			httpClient:   nil,
			wantedResult: publicAPIEndpoint,
		},
		{
			name:         "api key is nonempty but no pro",
			apiKey:       "test_api_key",
			isPro:        false,
			httpClient:   nil,
			wantedResult: publicAPIEndpoint,
		},
		{
			name:         "api key is nonempty and pro",
			apiKey:       "test_api_key",
			isPro:        true,
			httpClient:   nil,
			wantedResult: proAPIEndpoint,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCoinGecko(tt.apiKey, tt.isPro, tt.httpClient)
			if c.apiURL != tt.wantedResult {
				t.Fatalf("uncorrect api url, wanted url: %s, got url: %s", tt.wantedResult, c.apiKey)
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
		t.Fatalf("uncorrect http header, wanted header: %s", result)
	}
}
