package proxy

import (
	"errors"
	"io"
	"net"
	"net/http"
	"time"
)

var headers = []string{
	"Proxy-Connection",
}

func copyHeader(t, f http.Header) {
	for h, vv := range f {
		for _, v := range vv {
			t.Add(h, v)
		}
	}
}

func delHeaders(t http.Header) {
	for _, h := range headers {
		t.Del(h)
	}
}

func connectHandshake(w http.ResponseWriter, r *http.Request) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return nil, err
	}

	w.WriteHeader(http.StatusOK)
	return conn, nil
}

func connectHijacker(w http.ResponseWriter) (net.Conn, error) {
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return nil, errors.New("hijacking not supported")
	}

	conn, _, err := hijacker.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	return conn, nil
}

func transfer(dest io.WriteCloser, src io.ReadCloser) {
	defer dest.Close()
	defer src.Close()

	io.Copy(dest, src)
}
