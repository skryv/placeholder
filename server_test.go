package main

import (
	"fmt"
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
	router := CreateHandler()
	server := httptest.NewServer(router)

	URIS := []string{
		"/",
		"/api/health",
		"/api/public/health",
		"/api/public/health-container",
	}

	for _, uri := range URIS {
		url := fmt.Sprintf("%s%s", server.URL, uri)
		assert.HTTPStatusCode(t, HealthcheckHandler, "GET", url, nil, 200)
	}
}
