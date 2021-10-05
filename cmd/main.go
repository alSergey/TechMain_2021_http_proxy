package main

import (
	"log"
	"net/http"

	"github.com/alSergey/TechMain_2021_http_proxy/proxy"
)

func main() {
	p := &proxy.Proxy{}

	server := http.Server{
		Addr:    ":8080",
		Handler: p,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
