package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CoronaResponse struct {
	Uid      string
	Uf       string
	State    string
	Cases    int
	Deaths   int
	Suspects int
	Refuses  int
	Datetime string
}

func coronaVirus() (CoronaResponse, error) {
	resp, err := http.Get("https://covid19-brazil-api.now.sh/api/report/v1/brazil/uf/sp")
	var jsonResp CoronaResponse

	if err != nil {
		return jsonResp, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	strBody := string(body)

	json.Unmarshal([]byte(strBody), &jsonResp)
	return jsonResp, nil
}
