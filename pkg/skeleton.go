package skeleton

type NeccesaryWeather struct {
	Country     string  `json:"country"`
	Name        string  `json:"name"`
	Temperature float32 `json:"temp"`
	Description string  `json:"description"`
}

type Coordinate struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type MainInfo struct {
	Temperature    float32 `json:"temp"`
	FeelsLike      float32 `json:"feels_like"`
	TemperatureMin float32 `json:"temp_min"`
	TemperatureMax float32 `json:"temp_max"`
	Pressure       float32 `json:"pressure"`
	Humidity       float32 `json:"humidity"`
	SeaLevel       float32 `json:"sea_level"`
	GrundLevel     float32 `json:"grnd_level"`
}

type Wind struct {
	Speed  float32 `json:"speed"`
	Degree float32 `json:"deg"`
	Gust   float32 `json:"gust"`
}

type Clouds struct {
	All int8 `json:"all"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type JsonStructureParse struct {
	Coord      Coordinate `json:"coord"`
	Weather    []Weather  `json:"weather"`
	Base       string     `json:"base"`
	Main       MainInfo   `json:"main"`
	Visibility float32    `json:"visibility"`
	Wind       Wind       `json:"wind"`
	Clouds     Clouds     `json:"clouds"`
	Dt         int64      `json:"dt"`
	Sys        Sys        `json:"sys"`
	Timezone   int16      `json:"timezone"`
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Code       int16      `json:"cod"`
}
