package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	web "parser/internal/web"
)

var (
	l         *os.File
	logger    *log.Logger
	FILE_NAME = "logger"
)

func init() {

	locall, err := os.OpenFile(FILE_NAME+".log", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		os.Exit(1)
	}

	localLogger := log.New(locall, "[WEATHER PARSE]: ", log.Lshortfile|log.LstdFlags)

	// иначе перекрывает видимость
	l = locall
	logger = localLogger
}

func GenerateData(
	latStart int,
	latStop int,
	lonStart int,
	lonStop int,
	domain string,
	path string,
	APIKEY string,
) error {
	for lat := latStart; lat <= latStop; lat++ {
		for lon := lonStart; lon <= lonStop; lon++ {
			URL, err := web.CreateURL(domain, path, float32(lat), float32(lon), APIKEY)

			logger.Println(URL)

			if err != nil {
				return err
			}

			body, err := web.RequestToWeather(URL)

			if err != nil {
				return err
			}

			data, err := web.UnmarshalWeatherJSON(body)

			if err != nil {
				return err
			}

			logger.Println(data)

			//TODO: write to json
		}
	}

	return nil

}

func main() {

	defer l.Close()

	APIKEY := os.Getenv("APIKEY")
	domain := flag.String("domain", "https://api.openweathermap.org", "Site for parsing")
	path := flag.String("path", "/data/2.5/weather", "Path to data o nweb resource")

	err := GenerateData(0, 360, 0, 360, *domain, *path, APIKEY)

	if err != nil {
		logger.Printf("Error: %s\n", err.Error())
	}
}
