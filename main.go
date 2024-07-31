package main

import (
	"crypto/tls"
	b64 "encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/LuMa2003/clash-scouter-app/internal"
	"github.com/tidwall/gjson"

)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getAuth() (uint16, string) {
	dat, err := os.ReadFile("C:/Riot Games/League of Legends/lockfile")
	check(err)
	lockfile := strings.Split(string(dat), ":")

	port, _ := strconv.Atoi(lockfile[2])

	auth := b64.StdEncoding.EncodeToString([]byte("riot:" + lockfile[3]))

	return uint16(port), auth
}

func main() {
	start := time.Now()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	port, auth := getAuth()

	fmt.Printf("Port: %v Auth %v \n", port, auth)

	requestURL := fmt.Sprintf("https://127.0.0.1:%d/lol-summoner/v1/current-summoner", port)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	check(err)
	req.Header.Add("Authorization", "Basic "+auth)

	res, err := client.Do(req)
	check(err)

	fmt.Println(res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	check(err)

	Summoners_array := gjson.GetManyBytes(resBody, "gameName", "tagLine")
	cli.Cli(fmt.Sprint(Summoners_array[0].Value() ,"-" , Summoners_array[1].Value()))
	
	duration := time.Since(start)
	fmt.Println(duration)
}
