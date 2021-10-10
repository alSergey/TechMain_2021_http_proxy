package parser

import (
	"errors"
	"flag"
	"fmt"
)

type Parser struct {
	Protocol string
	CertFile string
	KeyFile  string
}

func (p *Parser) Create() error {
	flag.StringVar(&p.CertFile, "cert", "./certs/root.pem", "Cert file location")
	flag.StringVar(&p.KeyFile, "key", "./certs/root.key", "Key file location")
	flag.StringVar(&p.Protocol, "protocol", "http", "Protocol [http || https]")

	flag.Parse()

	if !p.isCorrect() {
		return errors.New("select http or https protocol")
	}

	return nil
}

func (p Parser) Usage() {
	fmt.Println("cert [cert_file] key [key_file] protocol [http || https]")
	flag.PrintDefaults()
}

func (p Parser) isCorrect() bool {
	if p.Protocol == "http" || p.Protocol == "https" {
		return true
	}

	return false
}
