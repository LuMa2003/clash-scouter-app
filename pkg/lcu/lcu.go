package lcu

import (
	"bytes"
	b64 "encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
)

type ConnInfo struct {
	Port uint16
	Auth string
}

type Request struct {
	Conn *ConnInfo
	Method   string
	Endpoint string
	Body     io.Reader
}

func GetAuth() (ConnInfo, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("wmic", "PROCESS", "WHERE", "name='LeagueClientUx.exe'", "GET", "commandline")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	if stderr.Len() != 0 {
		return ConnInfo{}, errors.New("League of Legends is not running")
	}

	port, _ := strconv.Atoi(string(regexp.MustCompile(`--app-port=([0-9]*)`).FindAllSubmatch(stdout.Bytes(), 1)[0][1]))
	token := string(regexp.MustCompile(`--remoting-auth-token=([\w-_]*)`).FindAllSubmatch(stdout.Bytes(), 1)[0][1])

	return ConnInfo{Port: uint16(port), Auth: b64.StdEncoding.EncodeToString([]byte("riot:" + token))}, nil
}

func LCU(request *Request) ([]byte, error) {
	requestURL := fmt.Sprint("https://127.0.0.1:", request.Conn.Port, request.Endpoint)
	req, err := http.NewRequest(request.Method, requestURL, request.Body)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "Basic "+request.Conn.Auth)
	res, err := http.DefaultClient.Do(req)
	if err != nil || res.StatusCode != 200 {
		return nil, errors.New("request failed")
	}
	resBody, err := io.ReadAll(res.Body)

	return resBody, nil
}
