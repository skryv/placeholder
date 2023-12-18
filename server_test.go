package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPortWhenNotSetReturns8080(t *testing.T) {
	assert.Equal(t, ":8080", GetAddress())
}

func TestGetPortWhen5000SetReturns5000(t *testing.T) {
	os.Setenv("PORT", "5000")
	assert.Equal(t, ":5000", GetAddress())
}

func TestServerEndToEnd(t *testing.T) {
	URIS := []string{
		"/",
		"/api/health",
		"/api/public/health",
		"/api/public/health-container",
	}

	for _, uri := range URIS {
		req, _ := http.NewRequest("GET", uri, nil)
		response := executeRequest(req)
		assert.Equal(t, 200, response.Code)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	mux := CreateHandler()
	mux.ServeHTTP(rr, req)

	return rr
}
