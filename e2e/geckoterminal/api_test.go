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

func TestClient_Networks(t *testing.T) {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetAllNetworks(context.Background(), 0)
	if err != nil {
		log.Fatalln(err)
	}
	slog.Info("call networks successfully", "response data", *data)

	b, err := json.Marshal(data)
	if err != nil {
		slog.Error("marshal error", "error", err)
	}
	err = os.WriteFile("./logs/networks.json", b, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("json file is saved")
}

// https://api.geckoterminal.com/api/v2/networks?page=1
