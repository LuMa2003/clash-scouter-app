package main

import (
	"crypto/tls"
	"fmt"
	"github.com/LuMa2003/clash-scouter-app/pkg/lcu"
	"github.com/tidwall/gjson"
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

	connInfo, err := lcu.GetAuth()

	check(err)

	data, err := lcu.LCU(lcu.Request{
		Conn:     &connInfo,
		Method:   "GET",
		Endpoint: "/riotclient/region-locale",
		Body:     nil,
	})
	check(err)

	region := gjson.GetBytes(data, "region").String()

	fmt.Println(region)

	duration := time.Since(start)
	fmt.Println(duration)
}
