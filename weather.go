package owmjp

type Weather struct {
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"Icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  float64 `json:"pressure"`
		Humidity  float64 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Clouds struct {
		All float64 `json:"all"`
	} `json:"clouds"`
	Rain struct {
		Rain1h float64 `json:"rain.1h"`
		Rain3h float64 `json:"rain.3h"`
	} `json:"rain"`
	Snow struct {
		Snow1h float64 `json:"snow.1h"`
		Snow3h float64 `json:"snow.3h"`
	} `json:"snow"`
	Dt    int    `json:"dt"`
	DtTxt string `json:"dt_txt"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
}
