package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/alSergey/TechMain_2021_http_proxy/parser"
	"github.com/alSergey/TechMain_2021_http_proxy/proxy"
)

func main() {
	var options parser.Parser
	if err := options.Create(); err != nil {
		fmt.Println(err.Error())
		options.Usage()
		return
	}

	p := &proxy.Proxy{}

	server := http.Server{
		Addr:    ":8080",
		Handler: p,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	switch options.Protocol {
	case "http":
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf(err.Error())
		}

		break
	case "https":
		if err := server.ListenAndServeTLS(options.CertFile, options.KeyFile); err != nil {
			log.Fatalf(err.Error())
		}

		break
	default:
		log.Fatal("select http or https")
	}
}
