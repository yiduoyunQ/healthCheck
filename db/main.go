package main

import (
	"log"
	"time"
)

func main() {
	log.Println("version:", version)

	addr, err := getDBAddr(fConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Get Addr:", addr)

	errCh := healthCheck(addr, timeout, readTimeout)

	select {
	case <-time.After(timeout):

		log.Fatal("Timeout:" + timeout.String())

	case err := <-errCh:
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println(addr, "db health check done!")
}
