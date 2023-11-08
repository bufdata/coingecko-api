package coingecko

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func setup(t *testing.T) *Client {
	t.Helper()
	return NewCoinGecko("", false, nil)
}

func TestClient_Ping(t *testing.T) {
	cases := []struct {
		name         string
		server       *httptest.Server
		wantedIsErr  bool
		wantedResult *PingResponse
		wantedErrStr string
	}{
		{
			name: "success",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				resp := &PingResponse{GeckoSays: "ok"}
				data, _ := json.Marshal(resp)
				_, _ = w.Write(data)
			})),
			wantedIsErr:  false,
			wantedResult: &PingResponse{GeckoSays: "ok"},
			wantedErrStr: "",
		},
		{
			name: "failure",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusBadRequest)
				_, _ = w.Write([]byte("mock error"))
			})),
			wantedIsErr:  true,
			wantedResult: nil,
			wantedErrStr: "status code: 400, error message: mock error",
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			client := setup(t)
			client.apiURL = tt.server.URL
			result, err := client.Ping(context.TODO())
			if tt.wantedIsErr {
				if !strings.Contains(err.Error(), tt.wantedErrStr) {
					t.Fatalf("uncorrect error, wanted error: %v, got error: %v", tt.wantedErrStr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("error should be nil, got: %v", err)
				}
			}
			if !reflect.DeepEqual(result, tt.wantedResult) {
				t.Fatalf("uncorrect response, wanted result: %+v, got result: %+v", tt.wantedResult, result)
			}
		})
	}
}
