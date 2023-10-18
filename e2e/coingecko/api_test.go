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
	api := coingecko.NewCoinGecko("", nil)
	data, err := api.Ping(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call ping successfully", "response data", *data)
}

func TestClient_SimplePriceOneCoin(t *testing.T) {
	api := coingecko.NewCoinGecko("", nil)
	data, err := api.SimplePrice(context.Background(), []string{"bitcoin"}, []string{"usd"}, emptyString, emptyString,
		emptyString, emptyString, emptyString)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call SimplePrice successfully", "response data", *data)
}

func TestClient_SimplePriceMultiCoins(t *testing.T) {
	api := coingecko.NewCoinGecko("", nil)
	data, err := api.SimplePrice(context.Background(), []string{"bitcoin", "ethereum"}, []string{"usd", "eur"},
		"true", "true", "true", "true", "18")
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call SimplePrice successfully", "response data", *data)
}

func TestClient_SimpleTokenPriceOneContractAddress(t *testing.T) {
	api := coingecko.NewCoinGecko("", nil)
	data, err := api.SimpleTokenPrice(context.Background(), "ethereum", []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"},
		[]string{"usd"}, emptyString, emptyString, emptyString, emptyString, emptyString)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call SimpleTokenPrice successfully", "response data", *data)
}

func TestClient_SimpleTokenPriceMultiContractAddresses(t *testing.T) {
	api := coingecko.NewCoinGecko("", nil)
	data, err := api.SimpleTokenPrice(context.Background(), "ethereum", []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
		"0xd533a949740bb3306d119cc777fa900ba034cd52"}, []string{"usd", "eur"}, "true", "true", "true", "true", "18")
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call SimpleTokenPrice successfully", "response data", *data)
}

func TestClient_SimpleSupportedVSCurrencies(t *testing.T) {
	api := coingecko.NewCoinGecko("", nil)
	data, err := api.SimpleSupportedVSCurrencies(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call SimpleSupportedVSCurrencies successfully", "response data", *data)
}
