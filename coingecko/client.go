package coingecko

import (
	"bytes"
	"context"
	"encoding/json"
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
func NewCoinGecko(apiKey string, httpClient *http.Client) *Client {
	var apiURL string
	if apiKey == "" {
		apiURL = publicAPIEndpoint
	} else {
		apiURL = proAPIEndpoint
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		httpClient: httpClient,
		apiURL:     apiURL,
		apiKey:     apiKey,
	}
}

func (c *Client) sendReq(ctx context.Context, endpoint string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		slog.Error("failed to new request with context", "error", err)
		return nil, err
	}
	data, err := c.doAPI(ctx, req)
	if err != nil {
		slog.Error("failed to do api", "url", req.URL.String(), "error", err)
	}
	return data, nil
}

func (c *Client) doAPI(ctx context.Context, req *http.Request) ([]byte, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		slog.Error("failed to do", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			slog.Error("failed to read error response", "error", err)
			return nil, err
		}
		if err = json.Unmarshal(data, &errResp); err != nil {
			slog.Error("failed to unmarshal error response", "error", err)
			return nil, err
		}
		return nil, fmt.Errorf("failed to call %s, status code: %d, error message: %s", req.URL.String(),
			resp.StatusCode, errResp.Error)
	}

	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		slog.Error("failed to parse resp body", "error", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

// checkAPIKey should be called after all query params are filled
func (c *Client) checkAPIKey(apiURL string) string {
	if c.apiKey != "" {
		return fmt.Sprintf("%s%s=%s", apiURL, proAPIKeyQueryParam, c.apiKey)
	} else {
		return apiURL
	}
}
