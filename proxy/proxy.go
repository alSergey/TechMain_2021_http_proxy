package proxy

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Proxy struct {
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodConnect {
		handleHTTPS(w, r)
	} else {
		handleHTTP(w, r)
	}
}

func handleHTTP(w http.ResponseWriter, r *http.Request) {
	connectID := uuid.New()
	fmt.Println("HTTP Connect", connectID, r.Method, r.URL)

	r.RequestURI = ""
	delHeaders(r.Header)

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal("handleHTTP:", err)
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Response", connectID, resp.Status)

	delHeaders(resp.Header)
	copyHeader(w.Header(), resp.Header)

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)

	fmt.Println("HTTP Close", connectID)
}

func handleHTTPS(w http.ResponseWriter, r *http.Request) {
	connectID := uuid.New()
	fmt.Println("HTTPS Connect", connectID, r.Method, r.URL)

	destConn, err := connectHandshake(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("connectHandshake:", err)
		return
	}

	fmt.Println("HTTPS handshake:", connectID)

	srcConn, err := connectHijacker(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("connectHijacker:", err)
		return
	}

	fmt.Println("HTTPS hijacker", connectID)

	go transfer(destConn, srcConn)
	go transfer(srcConn, destConn)

	fmt.Println("HTTPS open transfer", connectID)
}
