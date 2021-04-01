package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	server := CreateServer()
	server.ListenAndServe()
}

func CreateServer() *http.Server {
	mux := http.NewServeMux()

	paths := []string{
		"/",
		"/api/health",
		"/api/public/health",
	}

	for _, uri := range paths {
		mux.HandleFunc(uri, HealthcheckHandler)
	}

	server := &http.Server{
		Addr:    GetAddress(),
		Handler: mux,
	}

	return server
}

func GetAddress() string {
	var port string

	value, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	} else {
		port = value
	}

	return fmt.Sprintf(":%s", port)
}

func HealthcheckHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
}
