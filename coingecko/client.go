package coingecko

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

// Client struct
type Client struct {
	apiURL     string
	apiKey     string
	httpClient *http.Client
}

// NewCoinGecko create a new CoinGecko API client.
//
// For users with Pro API Key, users should use [https://pro-api.coingecko.com/api/v3/] to make API request.
// Therefore, you should provide apiKey and set isProAPIKey to true.
func NewCoinGecko(apiKey string, isProAPIKey bool, httpClient *http.Client) *Client {
	var apiURL string
	if apiKey == "" || !isProAPIKey {
		apiURL = publicAPIEndpoint
	} else {
		apiURL = proAPIEndpoint
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		apiURL:     apiURL,
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

func (c *Client) sendReq(ctx context.Context, endpoint string) ([]byte, http.Header, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		slog.Error("failed to new request with context", "error", err)
		return nil, nil, err
	}

	c.checkAPIKey(req)
	data, header, err := c.doAPI(req)
	if err != nil {
		slog.Error("failed to do api", "url", req.URL.String(), "error", err)
		return nil, nil, err
	}
	return data, header, nil
}

func (c *Client) doAPI(req *http.Request) ([]byte, http.Header, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		slog.Error("failed to do", "error", err)
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			slog.Error("failed to read error response", "error", err)
			return nil, nil, err
		}
		return nil, nil, fmt.Errorf("failed to call %s, status code: %d, error message: %s", req.URL.String(),
			resp.StatusCode, string(data))
	}

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		slog.Error("failed to parse resp body", "error", err)
		return nil, nil, err
	}
	return buf.Bytes(), resp.Header, nil
}

// check user whether provides api key, if provided adds it into http header.
//
// CoinGecko supports supplying API key in one of two ways:
//
// 1. Header: x-cg-pro-api-key
//
// 2. Query string parameter: x_cg_pro_api_key
func (c *Client) checkAPIKey(req *http.Request) {
	if c.apiKey != "" {
		req.Header.Add(proAPIKeyHeader, c.apiKey)
	}
}
