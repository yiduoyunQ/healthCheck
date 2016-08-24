package main

import (
	"bytes"
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
		_err   error
	)

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", "", 0, err
	}

	body := parser.GetSectionBody(content, domainSection)

	lines := bytes.Split(body, []byte{'\n'})
	for i := range lines {

		value := parser.GetValue(lines[i], domainKey)
		if value != "" {
			domain = value
			break
		}
	}

	body = parser.GetSectionBody(content, adminSection)
	lines = bytes.Split(body, []byte{'\n'})
	for i := range lines {

		value := parser.GetValue(lines[i], addrKey)
		if value != "" {

			ip, portStr, err := net.SplitHostPort(value)
			if err != nil {
				continue
			}

			n, err := strconv.Atoi(portStr)
			if err != nil {
				_err = err
				continue
			}

			addr = ip
			port = n
			_err = nil
			break
		}
	}

	return domain, addr, port, _err
}
