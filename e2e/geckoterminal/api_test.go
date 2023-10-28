package geckoterminal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"testing"

	"github.com/neuprotron/coingecko-api/geckoterminal"
)

const emptyString = ""

func TestClient_GetNetworks(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetNetworks(context.Background(), 0)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call GetNetworks successfully", "response data", *data)
}

func TestClient_GetDexes(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetDexes(context.Background(), "bsc", 0)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call GetDexes successfully", "response data", *data)
}

func TestClient_GetSpecificPool(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetSpecificPool(context.Background(), "eth", "0x60594a405d53811d3bc4766596efd80fd545a270",
		[]string{"base_token", "quote_token", "dex"})
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call GetSpecificPool successfully", "response data", *data)
}

func TestClient_GetMultiPools(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetMultiPools(context.Background(), "eth", []string{"base_token", "quote_token", "dex"},
		[]string{"0x60594a405d53811d3bc4766596efd80fd545a270", "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640"})
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call GetMultiPools successfully", "response data", *data)
}

func TestClient_GetTop20PoolsOnOneNetwork(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetTop20PoolsOnOneNetwork(context.Background(), "eth", []string{"base_token", "quote_token", "dex"})
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call GetTop20PoolsOnOneNetwork successfully", "response data", *data)
}

func TestClient_GetTop20PoolsOnOneDex(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetTop20PoolsOnOneDex(context.Background(), "eth", "sushiswap", []string{"base_token", "quote_token", "dex"})
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call GetTop20PoolsOnOneDex successfully", "response data", *data)
}

func TestClient_GetLatest20PoolsOnOneNetwork(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetLatest20PoolsOnOneNetwork(context.Background(), "eth", []string{"base_token", "quote_token", "dex"})
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call GetLatest20PoolsOnOneNetwork successfully", "response data", *data)
}

func TestClient_SearchPools(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.SearchPools(context.Background(), "ETH", "eth", []string{"base_token", "quote_token", "dex"})
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call SearchPools successfully", "response data", *data)

	b, err := json.Marshal(data)
	if err != nil {
		slog.Error("marshal error", "error", err)
	}
	err = os.WriteFile("./logs/search_pools.json", b, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("json file is saved")
}
