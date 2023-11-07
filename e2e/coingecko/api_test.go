package coingecko

import (
	"context"
	"log/slog"
	"testing"

	"github.com/bufdata/coingecko-api/coingecko"
)

const emptyString = ""

func TestClient_Ping(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.Ping(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ping successfully", "response data", *data)
}

func TestClient_SimplePriceOneCoin(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.SimplePrice(context.Background(), []string{"bitcoin"}, []string{"usd"}, emptyString, emptyString,
		emptyString, emptyString, emptyString)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call SimplePrice successfully", "response data", *data)
}

func TestClient_SimplePriceMultiCoins(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.SimplePrice(context.Background(), []string{"bitcoin", "ethereum"}, []string{"usd", "eur"},
		"true", "true", "true", "true", "18")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call SimplePrice successfully", "response data", *data)
}

func TestClient_SimpleTokenPriceOneContractAddress(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.SimpleTokenPrice(context.Background(), "ethereum", []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"},
		[]string{"usd"}, emptyString, emptyString, emptyString, emptyString, emptyString)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call SimpleTokenPrice successfully", "response data", *data)
}

func TestClient_SimpleTokenPriceMultiContractAddresses(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.SimpleTokenPrice(context.Background(), "ethereum", []string{"0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
		"0xd533a949740bb3306d119cc777fa900ba034cd52"}, []string{"usd", "eur"}, "true", "true", "true", "true", "18")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call SimpleTokenPrice successfully", "response data", *data)
}

func TestClient_SimpleSupportedVSCurrencies(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.SimpleSupportedVSCurrencies(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call SimpleSupportedVSCurrencies successfully", "response data", *data)
}

func TestClient_ListCoinsInfoTrue(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.ListCoinsInfo(context.Background(), true)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListSupportedCoinsInfo successfully", "response data", (*data)[0])
}

func TestClient_ListCoinsInfoFalse(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.ListCoinsInfo(context.Background(), false)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListCoinsInfo successfully", "response data", (*data)[0])
}

func TestClient_ListCoinsMarketsData(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.ListCoinsMarketsData(context.Background(), "usd", []string{"bitcoin", "ethereum"}, emptyString,
		emptyString, 0, 0, false, []string{"1h", "24h", "7d"}, emptyString, emptyString)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListCoinsMarketsData successfully", "response data", (*data)[0])
}

func TestClient_GetCoinDataByCoinID(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetCoinDataByCoinID(context.Background(), "ethereum", true, true, true, true, true, false)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetCoinDataByCoinID successfully", "response data", data.Name)
}

func TestClient_GetTickersByCoinID(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, pageCount, err := api.GetCoinTickersByCoinID(context.Background(), "ethereum", "", true, 1, emptyString, true)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetCoinTickersByCoinID successfully", "response data", data.Name, "page count", pageCount)
}

func TestClient_GetCoinHistoryDataByCoinID(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetCoinHistoryDataByCoinID(context.Background(), "ethereum", "01-10-2023", true)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetCoinHistoryDataByCoinID successfully", "response data", data.Name)
}

func TestClient_GetCoinMarketChartByCoinID(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetCoinMarketChartByCoinID(context.Background(), "ethereum", "usd", "max", "daily", "full")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetCoinMarketChartByCoinID successfully", "response data", len(data.MarketCaps))
}

func TestClient_GetCoinMarketChartRangeByCoinID(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetCoinMarketChartRangeByCoinID(context.Background(), "ethereum", "usd", "1682477232", "1682577232", "full")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetCoinMarketChartRangeByCoinID successfully", "response data", len(data.MarketCaps))
}

func TestClient_GetCoinOHLCByCoinID(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetCoinOHLCByCoinID(context.Background(), "ethereum", "usd", "1", "full")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetCoinOHLCByCoinID successfully", "response data", len(*data))
}

func TestClient_GetCoinInfoByContractAddress(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetCoinInfoByContractAddress(context.Background(), "ethereum", "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetCoinInfoByContractAddress successfully", "response data", data.Name)
}

func TestClient_GetMarketChartByContractAddress(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetMarketChartByContractAddress(context.Background(), "ethereum", "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
		"usd", "1", "full")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetMarketChartByContractAddress successfully", "response data", len(data.Prices))
}

func TestClient_GetMarketChartRangeByContractAddress(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetMarketChartRangeByContractAddress(context.Background(), "ethereum", "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
		"usd", "1682477232", "1682577232", "full")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetMarketChartRangeByContractAddress successfully", "response data", len(data.Prices))
}

func TestClient_ListAllAssetPlatforms(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.ListAllAssetPlatforms(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListAllAssetPlatforms successfully", "response data", *data)
}

func TestClient_ListAllCategories(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.ListAllCategories(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListAllCategories successfully", "response data", len(*data))
}

func TestClient_ListAllCategoriesWithMarketData(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.ListAllCategoriesWithMarketData(context.Background(), emptyString)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListAllCategoriesWithMarketData successfully", "response data", len(*data))
}

func TestClient_ListAllExchanges(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, pageCount, err := api.ListAllExchanges(context.Background(), 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListAllExchanges successfully", "response data", len(*data), "page count", pageCount)
}

func TestClient_ListAllMarketsInfo(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.ListAllMarketsInfo(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListAllMarketsInfo successfully", "response data", len(*data))
}

func TestClient_GetVolumeAndTickersByExchangeID(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetExchangeVolumeAndTickersByExchangeID(context.Background(), "uniswap_v3")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetExchangeVolumeAndTickersByExchangeID successfully", "response data", data.Name)
}

func TestClient_GetExchangeTickersByExchangeID(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, count, err := api.GetExchangeTickersByExchangeID(context.Background(), "binance", "curve-dao-token", true, 1, true, "")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetExchangeTickersByExchangeID successfully", "response data", data.Name, "count", count)
}

func TestClient_GetExchangeVolumeChartByExchangeID(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetExchangeVolumeChartByExchangeID(context.Background(), "binance", 1)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetExchangeVolumeChartByExchangeID successfully", "response data", len(*data))
}

func TestClient_ListAllDerivativesTickers(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.ListAllDerivativesTickers(context.Background(), "")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListAllDerivativesTickers successfully", "response data", len(*data))
}

func TestClient_ListAllDerivativesExchanges(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, count, err := api.ListAllDerivativesExchanges(context.Background(), "", 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListAllDerivativesExchanges successfully", "response data", len(*data), "count", count)
}

func TestClient_ListDerivativesExchangeData(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.ListDerivativesExchangeData(context.Background(), "binance_futures", "all")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListDerivativesExchangeData successfully", "response data", data.Name)
}

func TestClient_ListAllDerivativeExchangeInfo(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.ListAllDerivativeExchangeInfo(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListAllDerivativeExchangeInfo successfully", "response data", len(*data))
}

func TestClient_ListAllNFTInfo(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, count, err := api.ListAllNFTInfo(context.Background(), "", "", 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call ListAllNFTInfo successfully", "response data", data, "count", count)
}

func TestClient_GetDataByNFTID(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetNFTDataByNFTID(context.Background(), "ag3dnft")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetNFTDataByNFTID successfully", "response data", data)
}

func TestClient_GetNFTDataByAssetPlatformIDAndContractAddress(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetNFTDataByAssetPlatformIDAndContractAddress(context.Background(), "binance-smart-chain", "0x4bafc595a9ff4a5f4936689a0389c148a65456a2")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetNFTDataByAssetPlatformIDAndContractAddress successfully", "response data", data)
}

func TestClient_GetExchangeRates(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetExchangeRates(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetExchangeRates successfully", "response data", data)
}

func TestClient_Search(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.Search(context.Background(), "bnb")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call Search successfully", "response data", data)
}

func TestClient_SearchTrending(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.SearchTrending(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call SearchTrending successfully", "response data", data)
}

func TestClient_GetGlobalCryptocurrencyData(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetGlobalCryptocurrencyData(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetGlobalCryptocurrencyData successfully", "response data", data)
}

func TestClient_GetGlobalTop100DefiData(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetGlobalTop100DefiData(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetGlobalTop100DefiData successfully", "response data", data)
}

func TestClient_GetCompaniesPublicTreasury(t *testing.T) {
	api := coingecko.NewCoinGecko(emptyString, false, nil)
	data, err := api.GetCompaniesPublicTreasury(context.Background(), "ethereum")
	if err != nil {
		t.Fatal(err)
	}
	slog.Info("call GetCompaniesPublicTreasury successfully", "response data", data.MarketCapDominance)
}
