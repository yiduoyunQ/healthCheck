package main

import (
	"io/ioutil"
	"net"
	"strconv"

	"github.com/fugr/healthCheck/parser"
)

var (
	domainSection = []byte("upsql-proxy")
	domainKey     = []byte("proxy-domain")

	adminSection = []byte("adm-cli")
	addrKey      = []byte("adm-cli-address")
)

func getConfigValue(file string) (string, string, int, error) {
	var (
		domain string
		addr   string
		port   int
	)

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", "", 0, err
	}

	body := parser.GetSectionBody(content, domainSection)

	validDomain := func(str string) bool {
		return str != ""
	}

	value := parser.GetString(body, domainKey, validDomain)
	if value != "" {
		domain = value
	}

	body = parser.GetSectionBody(content, adminSection)

	validAddr := func(str string) bool {
		_, portStr, err := net.SplitHostPort(value)
		if err != nil {
			return false
		}

		n, err := strconv.Atoi(portStr)
		if err != nil && n > 0 {
			return false
		}

		return true
	}

	value = parser.GetString(body, addrKey, validAddr)
	if value != "" {

		ip, portStr, err := net.SplitHostPort(value)
		if err != nil {
			return domain, "", 0, err
		}

		n, err := strconv.Atoi(portStr)
		if err != nil {
			return domain, ip, n, err
		}

		addr = ip
		port = n
	}

	return domain, addr, port, nil
}
