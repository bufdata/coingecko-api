package coingecko

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	invalidJSONString = `{"name":what?}`

	invalidCharacterJSONErrStr   = "invalid character 'w' looking for beginning of value"
	unexpectedEndJSONInputErrStr = "unexpected end of JSON input"
	statusCode400ErrStr          = "status code: 400, error message: invalid request params"
	incorrectJSONTypeErrStr      = "json: cannot unmarshal string into Go value of type"
)

func setup(t *testing.T) *Client {
	t.Helper()
	return NewCoinGecko("", false, nil)
}

// func mockHTTPServer(t *testing.T, resp any) *httptest.Server {
// 	t.Helper()
// 	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		data, _ := json.Marshal(resp)
// 		_, _ = w.Write(data)
// 	}))
// 	return svr
// }

func mockHTTPServer(t *testing.T, totalCount, resp string) *httptest.Server {
	t.Helper()

	var svr *httptest.Server
	if totalCount == "" {
		svr = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(resp))
		}))
	} else {
		svr = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(totalHeader, totalCount)
			_, _ = w.Write([]byte(resp))
		}))
	}
	return svr
}

func mockErrorHTTPServer(t *testing.T, totalCount string) *httptest.Server {
	t.Helper()

	var svr *httptest.Server
	if totalCount == "" {
		svr = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("invalid request params"))
		}))
	} else {
		svr = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add(totalHeader, totalCount)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("invalid request params"))
		}))
	}
	return svr
}
