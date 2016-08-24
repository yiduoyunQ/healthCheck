package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net"
	"strconv"

	"github.com/fugr/healthCheck/parser"
)

func getDBAddr(file string) (string, error) {
	config, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	var (
		addr string
		port string

		section   = []byte("mysqld")
		bind      = []byte("bind_address")
		portBytes = []byte("port")
	)

	body := parser.GetSectionBody(config, section)
	lines := bytes.Split(body, []byte{'\n'})

	for i := range lines {
		if addr != "" && port != "" {

			return addr + ":" + port, nil
		}

		if str := parser.GetValue(lines[i], bind); str != "" {
			ip := net.ParseIP(str)
			if ip != nil {
				addr = str
				continue
			}
		}

		if str := parser.GetValue(lines[i], portBytes); str != "" {

			n, err := strconv.Atoi(str)
			if err == nil && n > 0 {
				port = str
			}
		}
	}

	return "", errors.New("Not Found Address")
}
