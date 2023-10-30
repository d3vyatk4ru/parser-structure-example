package skeleton

import "fmt"

type Weather struct {
	Description string `json:"description"`
}

type Main struct {
	Temperature float32 `json:"temp"`
}

type Sys struct {
	Country string `json:"country"`
}

type NeccesaryWeather struct {
	Sys     Sys       `json:"sys"`
	Name    string    `json:"name"`
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
}

func (w NeccesaryWeather) String() string {
	return fmt.Sprintf(
		"Country: %s, Name: %s, Temperature: %v, Description: %s",
		w.Sys.Country, w.Name, w.Main.Temperature, w.Weather[0].Description,
	)
}
