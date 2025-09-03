package openweather

type CoordinatesResponse struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

type Coordinates struct {
	Lat float64
	Lon float64
}

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
		//FeelsLike float64 `json:"feels_like"`
		//TempMin   float64 `json:"temp_min"`
		//TempMax   float64 `json:"temp_max"`
		//Pressure  int     `json:"pressure"`
		//Humidity  int     `json:"humidity"`
		//SeaLevel  int     `json:"sea_level"`
		//GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
}

type Weather struct {
	Temp float64
}
