package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const APIKEY = "a6cc01f8b7350c685d5afe22c3b63a04"

func CreateURL(
	domain string,
	resource string,
	lat float32,
	lon float32,
	apikey string,
) (string, error) {

	params := url.Values{}
	params.Add("lat", fmt.Sprintf("%v", lat))
	params.Add("lon", fmt.Sprintf("%v", lon))
	params.Add("APPID", apikey)

	u, err := url.ParseRequestURI(domain)

	if err != nil {
		return "", err
	}

	u.Path = resource
	u.RawQuery = params.Encode()

	log.Println(u.String())

	return u.String(), nil
}

func RequestToWeather(URL string) ([]byte, error) {

	resp, err := http.Get(URL)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func UnmarshalWeatherJSON(body []byte) (NeccesaryWeather, error) {

	jsonData := JsonStructureParse{}

	err := json.Unmarshal(body, &jsonData)

	if err != nil {
		return NeccesaryWeather{}, err
	}

	return NeccesaryWeather{
		Country:     jsonData.Sys.Country,
		Name:        jsonData.Name,
		Temperature: jsonData.Main.Temperature,
		Description: jsonData.Weather[0].Description,
	}, nil
}

func main() {

	URL, err := CreateURL("https://api.openweathermap.org", "/data/2.5/weather", 10, 5, APIKEY)

	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}

	body, err := RequestToWeather(URL)

	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}

	data, err := UnmarshalWeatherJSON(body)

	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}

	log.Println(data)
}
