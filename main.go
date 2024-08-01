package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	start := time.Now()

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	duration := time.Since(start)
	fmt.Println(duration)
}
