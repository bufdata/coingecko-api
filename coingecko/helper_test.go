package coingecko

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	invalidCharacterJSONErrStr   = "invalid character 'w' looking for beginning of value"
	unexpectedEndJSONInputErrStr = "unexpected end of JSON input"
	statusCode400ErrStr          = "status code: 400, error message: mock error"
)

func setup(t *testing.T) *Client {
	t.Helper()
	return NewCoinGecko("", false, nil)
}

func mockHTTPServer(t *testing.T) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	return srv
}

func mockErrorHTTPServer(t *testing.T) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("mock error"))
	}))
	return srv
}

func mockInvalidJSONHTTPServer(t *testing.T) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"name":what?}`))
	}))
	return srv
}
