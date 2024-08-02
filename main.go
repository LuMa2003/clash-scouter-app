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
	"os"
	"time"
)

func main() {
	start := time.Now()

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	prompt := promptui.Select{
		Label:    "Something went wrong",
		HideHelp: true,
		Items:    []string{"Retry", "Exit"},
	}

	var connInfo lcu.ConnInfo
	var err error
	for {
		connInfo, err = lcu.GetAuth()
		if err != nil {
			prompt.Label = "League of Legends is not running"
			choice, _, _ := prompt.Run()
			if choice == 1 {
				os.Exit(0)
			}
		} else {
			break
		}
	}

	data, err := lcu.LCU(lcu.Request{
		Conn:     &connInfo,
		Method:   "GET",
		Endpoint: "/riotclient/region-locale",
		Body:     nil,
	})
	if err != nil {
		panic(err)
	}
	region := gjson.GetBytes(data, "region").String()

	var summonerArray []clash.Summoner
	for {
		summonerArray, err = clash.ClashOpponent(&connInfo)
		if err != nil {
			prompt.Label = "No Clash opponent found"
			choice, _, _ := prompt.Run()
			if choice == 1 {
				os.Exit(0)
			}
		} else {
			break
		}
	}

	cli.Cli(&summonerArray, region)

	duration := time.Since(start)
	fmt.Println(duration)
}
