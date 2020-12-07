package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CountryInformation struct {
	Country    string
	Cases      int
	Confirmed  int
	Deaths     int
	Recovered  int
	Updated_at string
}


type StateInformation struct {
	State      string
	UF		   string		
	Cases	   int
	Deaths     int
	Suspects   int
	Refuses    int
	Datetime   string
}

//Country 
type single struct {
    Data CountryInformation `json:"data"`
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


func getCountry(place string) (CountryInformation, error) {

	if place == "" {
		place = "brazil"
	  }

	  
	strBody, _ := GETRequest("https://covid19-brazil-api.now.sh/api/report/v1/"+place)
	var parsed single
	json.Unmarshal([]byte(strBody), &parsed)
	var country CountryInformation
	country.Country = parsed.Data.Country
	country.Cases = parsed.Data.Cases
	country.Confirmed += parsed.Data.Confirmed
	country.Deaths += parsed.Data.Deaths
	country.Recovered += parsed.Data.Recovered
	
	return country, nil


}


/* SOON */
func getState(place string) (StateInformation, error) {

	if place == "" {
		place = "df"
	  }

	  
	strBody, _ := GETRequest("https://covid19-brazil-api.now.sh/api/report/v1/brazil/uf/"+place)
	var state StateInformation
	json.Unmarshal([]byte(strBody), &state)

	return state, nil


}





