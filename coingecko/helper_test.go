package coingecko

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	invalidCharacterJSONErrStr   = "invalid character 'w' looking for beginning of value"
	unexpectedEndJSONInputErrStr = "unexpected end of JSON input"
	statusCode400ErrStr          = "status code: 400, error message: invalid request params"
	incorrectJSONTypeErrStr      = "json: cannot unmarshal string into Go value of type"
)

func setup(t *testing.T) *Client {
	t.Helper()
	return NewCoinGecko("", false, nil)
}

func mockHTTPServer(t *testing.T, resp any) *httptest.Server {
	t.Helper()
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(resp)
		_, _ = w.Write(data)
	}))
	return svr
}

func mockStringHTTPServer(t *testing.T, resp string) *httptest.Server {
	t.Helper()
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(resp))
	}))
	return svr
}

func mockErrorHTTPServer(t *testing.T) *httptest.Server {
	t.Helper()
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid request params"))
	}))
	return svr
}

//nolint:unused
func mockInvalidJSONHTTPServer(t *testing.T) *httptest.Server {
	t.Helper()
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"name":what?}`))
	}))
	return svr
}
