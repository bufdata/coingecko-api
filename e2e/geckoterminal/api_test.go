package geckoterminal

import (
	"context"
	"log/slog"
	"testing"

	"github.com/bufdata/coingecko-api/geckoterminal"
)

func TestClient_GetNetworks(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetNetworks(context.Background(), 0)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetNetworks successfully", "response data", *data)
}

func TestClient_GetDexes(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetDexes(context.Background(), "bsc", 0)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetDexes successfully", "response data", *data)
}

func TestClient_GetSpecificPool(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetSpecificPool(context.Background(), "eth", "0x60594a405d53811d3bc4766596efd80fd545a270",
		[]string{"base_token", "quote_token", "dex"})
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetSpecificPool successfully", "response data", *data)
}

func TestClient_GetMultiPools(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetMultiPools(context.Background(), "eth", []string{"base_token", "quote_token", "dex"},
		[]string{"0x60594a405d53811d3bc4766596efd80fd545a270", "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"})
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetMultiPools successfully", "response data", *data)
}

func TestClient_GetTop20PoolsOnOneNetwork(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetTop20PoolsOnOneNetwork(context.Background(), "eth", []string{"base_token", "quote_token", "dex"})
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetTop20PoolsOnOneNetwork successfully", "response data", *data)
}

func TestClient_GetTop20PoolsOnOneDex(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetTop20PoolsOnOneDex(context.Background(), "eth", "sushiswap", []string{"base_token", "quote_token", "dex"})
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetTop20PoolsOnOneDex successfully", "response data", *data)
}

func TestClient_GetLatest20PoolsOnOneNetwork(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetLatest20PoolsOnOneNetwork(context.Background(), "eth", []string{"base_token", "quote_token", "dex"})
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetLatest20PoolsOnOneNetwork successfully", "response data", *data)
}

func TestClient_SearchPools(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.SearchPools(context.Background(), "ETH", "eth", []string{"base_token", "quote_token", "dex"})
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call SearchPools successfully", "response data", *data)
}

func TestClient_GetTop20PoolsForOneToken(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetTop20PoolsForOneToken(context.Background(), "eth", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		[]string{"base_token", "quote_token", "dex"})
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetTop20PoolsForOneToken successfully", "response data", *data)
}

func TestClient_GetSpecificTokenOnOneNetwork(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetSpecificTokenOnOneNetwork(context.Background(), "eth", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		[]string{"top_pools"})
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetSpecificTokenOnOneNetwork successfully", "response data", *data)
}

func TestClient_GetMultiTokensOnOneNetwork(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetMultiTokensOnOneNetwork(context.Background(), "eth", []string{"0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"},
		[]string{"top_pools"})
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetMultiTokensOnOneNetwork successfully", "response data", *data)
}

func TestClient_GetSpecificTokenInfoOnOneNetwork(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetSpecificTokenInfoOnOneNetwork(context.Background(), "eth", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetSpecificTokenInfoOnOneNetwork successfully", "response data", *data)
}

func TestClient_GetPoolTokensInfoOnOneNetwork(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetPoolTokensInfoOnOneNetwork(context.Background(), "eth", "0x0d4a11d5eeaac28ec3f61d100daf4d40471f1852")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetPoolTokensInfoOnOneNetwork successfully", "response data", *data)
}

func TestClient_GetRecentlyUpdated100TokensInfo(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetRecentlyUpdated100TokensInfo(context.Background(), []string{"network"})
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetRecentlyUpdated100TokensInfo successfully", "response data", *data)
}

func TestClient_GetOHLCV(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetOHLCV(context.Background(), "eth", "0x60594a405d53811d3bc4766596efd80fd545a270",
		"day", 1, 1697658844, 100, "usd", "base")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetOHLCV successfully", "response data", *data)
}
