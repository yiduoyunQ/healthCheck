package main

import (
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

	validIP := func(str string) bool {
		return net.ParseIP(str) != nil
	}

	if str := parser.GetString(body, bind, validIP); str != "" {
		addr = str
	}

	validPort := func(str string) bool {
		n, err := strconv.Atoi(str)
		if err == nil && n > 0 {
			return true
		}

		return false
	}

	if str := parser.GetString(body, portBytes, validPort); str != "" {
		port = str
	}

	if addr != "" && port != "" {
		return addr + ":" + port, nil
	}

	return "", errors.New("not found address in " + file)
}
