package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const APIKEY = "a6cc01f8b7350c685d5afe22c3b63a04"

type Weather struct {
	Country     string  `json:"country"`
	Area        string  `json:"area"`
	Temperature float32 `json:"temperature"`
	Description string  `json:"desc"`
}

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
	params.Add("appid", apikey)

	u, err := url.ParseRequestURI(domain)

	if err != nil {
		return "", err
	}

	u.Path = resource
	u.RawQuery = params.Encode()

	return u.String(), nil
}

func RequestToWeather(URL string) error {

	resp, err := http.Get(URL)

	if err != nil {
		return err
	}

	resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	fmt.Println(body)

	return nil

}

func main() {

	URL, err := CreateURL("https://openweathermap.org", "/data/2.5/weather", 10, 5, APIKEY)

	fmt.Println(URL)

	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}

	RequestToWeather(URL)
}
