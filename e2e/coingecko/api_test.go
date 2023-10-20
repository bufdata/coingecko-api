package coingecko

import (
	"context"
	"log"
	"log/slog"
	"testing"

	"github.com/neuprotron/coingecko-api/coingecko"
)

const emptyString = ""

func TestClient_Ping(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.Ping(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call ping successfully", "response data", *data)
}

func TestClient_SimplePriceOneCoin(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.SimplePrice(context.Background(), []string{"bitcoin"}, []string{"usd"}, emptyString, emptyString,
		emptyString, emptyString, emptyString)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call SimplePrice successfully", "response data", *data)
}

func TestClient_SimplePriceMultiCoins(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.SimplePrice(context.Background(), []string{"bitcoin", "ethereum"}, []string{"usd", "eur"},
		"true", "true", "true", "true", "18")
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call SimplePrice successfully", "response data", *data)
}

func TestClient_SimpleTokenPriceOneContractAddress(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.SimpleTokenPrice(context.Background(), "ethereum", []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"},
		[]string{"usd"}, emptyString, emptyString, emptyString, emptyString, emptyString)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call SimpleTokenPrice successfully", "response data", *data)
}

func TestClient_SimpleTokenPriceMultiContractAddresses(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.SimpleTokenPrice(context.Background(), "ethereum", []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
		"0xd533a949740bb3306d119cc777fa900ba034cd52"}, []string{"usd", "eur"}, "true", "true", "true", "true", "18")
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call SimpleTokenPrice successfully", "response data", *data)
}

func TestClient_SimpleSupportedVSCurrencies(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.SimpleSupportedVSCurrencies(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call SimpleSupportedVSCurrencies successfully", "response data", *data)
}

func TestClient_CoinsListTrue(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.CoinsList(context.Background(), true)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call CoinsList successfully", "response data", (*data)[0])
}

func TestClient_CoinsListFalse(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.CoinsList(context.Background(), false)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call CoinsList successfully", "response data", (*data)[0])
}

func TestClient_CoinsMarkets(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.CoinsMarkets(context.Background(), "usd", []string{"bitcoin", "ethereum"}, emptyString,
		emptyString, 0, 0, false, []string{"1h", "24h", "7d"}, emptyString, emptyString)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call CoinsMarkets successfully", "response data", (*data)[0])
}

func TestClient_CoinsID(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.CoinsID(context.Background(), "ethereum", true, true, true, true, true, false)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call CoinsID successfully", "response data", data.Name)
}

func TestClient_CoinsIDTickers(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, pageCount, err := api.CoinsIDTickers(context.Background(), "ethereum", "", true, 1, emptyString, true)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call CoinsIDTickers successfully", "response data", data.Name, "page count", pageCount)
}

func TestClient_CoinsIDHistory(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.CoinsIDHistory(context.Background(), "ethereum", "01-10-2023", true)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call CoinsIDHistory successfully", "response data", data.Name)
}

func TestClient_CoinsIDMarketChart(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.CoinsIDMarketChart(context.Background(), "ethereum", "usd", "max", "daily", "full")
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call CoinsIDMarketChart successfully", "response data", len(data.MarketCaps))
}

func TestClient_CoinsIDMarketChartRange(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.CoinsIDMarketChartRange(context.Background(), "ethereum", "usd", "1682477232", "1682577232", "full")
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call CoinsIDMarketChartRange successfully", "response data", len(data.MarketCaps))
}

func TestClient_CoinsIDOHLC(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.CoinsIDOHLC(context.Background(), "ethereum", "usd", "1", "full")
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call CoinsIDOHLC successfully", "response data", len(*data))
}

func TestClient_CoinsContract(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.CoinsContract(context.Background(), "ethereum", "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984")
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call CoinsContract successfully", "response data", data.Name)
}

func TestClient_CoinsContractMarketChart(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.CoinsContractMarketChart(context.Background(), "ethereum", "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
		"usd", "1", "full")
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call CoinsContract successfully", "response data", len(data.Prices))
}

func TestClient_CoinsContractMarketChartRange(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.CoinsContractMarketChartRange(context.Background(), "ethereum", "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
		"usd", "1682477232", "1682577232", "full")
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call CoinsContractMarketChartRange successfully", "response data", len(data.Prices))
}

func TestClient_AssetPlatforms(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, nil)
	data, err := api.AssetPlatforms(context.Background(), "")
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call AssetPlatforms successfully", "response data", *data)
}
