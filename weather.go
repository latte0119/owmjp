package owmjp

import "errors"

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

func (w Weather) MainJP() (string, error) {
	id := w.ID
	if id == 0 {
		return "", errors.New("Error : ID is undefined")
	}

	switch {
	case 200 <= id && id < 300:
		return "雷雨", nil
	case 300 <= id && id < 400:
		return "霧雨", nil
	case 400 <= id && id < 500:
		return "雨", nil
	case 500 <= id && id < 600:
		return "雪", nil
	case id == 701:
		return "靄", nil
	case id == 711:
		return "煙", nil
	case id == 721:
		return "霞", nil
	case id == 731:
		return "砂塵", nil
	case id == 741:
		return "霧", nil
	case id == 751:
		return "砂あらし", nil
	case id == 761:
		return "砂塵", nil
	case id == 762:
		return "灰", nil
	case id == 771:
		return "スコール", nil
	case id == 781:
		return "竜巻", nil
	case id == 800:
		return "快晴", nil
	case 800 < id:
		return "曇り", nil
	default:
		return "", errors.New("Error : ID is invalid")
	}
}
