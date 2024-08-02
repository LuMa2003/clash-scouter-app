package main

import (
	"crypto/tls"
	"fmt"
	"github.com/LuMa2003/clash-scouter-app/internal/clash"
	"github.com/LuMa2003/clash-scouter-app/internal/cli"
	"github.com/LuMa2003/clash-scouter-app/pkg/lcu"
	"github.com/manifoldco/promptui"
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

	for {
		connInfo, err := lcu.GetAuth()

		if err == nil {
			// If no error, print success and break out of the loop
			data, err := lcu.LCU(lcu.Request{
				Conn:     &connInfo,
				Method:   "GET",
				Endpoint: "/riotclient/region-locale",
				Body:     nil,
			})
			check(err)
			region := gjson.GetBytes(data, "region").String()

			summoner_array, err := clash.ClashOpponent(&connInfo)
			check(err)
			cli.Cli(&summoner_array, region)
			break
		}

		prompt := promptui.Select{
			Label:    err,
			HideHelp: true,
			Size:     10,
			Items:    []string{"Retry", "Exit"},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if result == "Exit" {
			fmt.Println("Exiting...")
			return
		}

	}
	duration := time.Since(start)
	fmt.Println(duration)
}
