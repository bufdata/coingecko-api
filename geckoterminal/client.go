package geckoterminal

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
	httpClient *http.Client
}

// NewGeckoTerminal create a new GeckoTerminal API client.
func NewGeckoTerminal(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) sendReq(ctx context.Context, endpoint string) ([]byte, http.Header, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		slog.Error("failed to new request with context", "endpoint", endpoint, "error", err)
		return nil, nil, err
	}

	data, header, err := c.doAPI(req)
	if err != nil {
		slog.Error("failed to do api", "url", req.URL.String(), "header", req.Header, "error", err)
		return nil, nil, err
	}
	return data, header, nil
}

func (c *Client) doAPI(req *http.Request) ([]byte, http.Header, error) {
	req.Header.Add(acceptHeader, jsonHeader)
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
