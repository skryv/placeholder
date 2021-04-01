package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerReturnsOkOnRootAndHealthcheck(t *testing.T) {
	for _, uri := range []string{
		"/",
		"/api/health",
	} {
		req, _ := http.NewRequest(http.MethodGet, uri, nil)

		requestRecorder := httptest.NewRecorder()
		handler := http.HandlerFunc(HealthcheckHandler)

		handler.ServeHTTP(requestRecorder, req)

		assert.Equal(t, http.StatusOK, requestRecorder.Code)
	}
}

func TestGetPortWhenNotSetReturns8080(t *testing.T) {
	assert.Equal(t, ":8080", GetAddress())
}

func TestGetPortWhen5000SetReturns5000(t *testing.T) {
	os.Setenv("PORT", "5000")
	assert.Equal(t, ":5000", GetAddress())
}
