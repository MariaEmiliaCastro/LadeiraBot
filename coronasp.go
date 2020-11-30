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

type CountryInformation struct {
	Country    string
	Cases      int
	Confirmed  int
	Deaths     int
	Recovered  int
	Updated_at string
}

type AllCountriesResponse struct {
	Data []CountryInformation
}

func GETRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	return strBody, nil
}

func coronaVirus() (CoronaResponse, error) {
	strBody, _ := GETRequest("https://covid19-brazil-api.now.sh/api/report/v1/brazil/uf/sp")
	var jsonResp CoronaResponse
	json.Unmarshal([]byte(strBody), &jsonResp)
	return jsonResp, nil
}

func allCountriesCorona() (CountryInformation, error) {
	strBody, _ := GETRequest("https://covid19-brazil-api.now.sh/api/report/v1/countries")
	var parsed AllCountriesResponse
	json.Unmarshal([]byte(strBody), &parsed)
	var world CountryInformation
	world.Country = "World"
	for _, data := range parsed.Data {
		world.Cases += data.Cases
		world.Confirmed += data.Confirmed
		world.Deaths += data.Deaths
		world.Recovered += data.Recovered
	}

	return world, nil
}
