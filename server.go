package main

import (
	"fmt"
	"net/http"
	"os"
)

const PORT_ENV_KEY = "PORT"

func main() {
	server := CreateServer()
	server.ListenAndServe()
}

func CreateServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HealthcheckHandler)
	mux.HandleFunc("/api/health", HealthcheckHandler)

	server := &http.Server{
		Addr:    GetAddress(),
		Handler: mux,
	}

	return server
}

func GetAddress() string {
	var port string

	value, ok := os.LookupEnv(PORT_ENV_KEY)

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
