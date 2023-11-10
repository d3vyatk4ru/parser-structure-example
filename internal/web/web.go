package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	s "parser/internal/skeleton"
)

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

	return u.String(), nil
}

func RequestToWeather(URL string) ([]byte, error) {

	client := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   time.Second,
			ResponseHeaderTimeout: time.Second,
		},
	}

	resp, err := client.Get(URL)

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

func UnmarshalWeatherJSON(body []byte) (s.NeccesaryWeather, error) {

	jsonData := s.NeccesaryWeather{}

	err := json.Unmarshal(body, &jsonData)

	if err != nil {
		return s.NeccesaryWeather{}, err
	}

	return jsonData, nil
}
