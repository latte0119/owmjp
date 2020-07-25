package owmjp

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

func (a API) GetCurrentWeatherData(name string, v url.Values) (weather Weather, err error) {
	if v == nil {
		v = url.Values{}
	}
	v.Set("q", name)
	v.Set("appid", a.APIKey)

	res, err := a.HTTPClient.Get(a.baseURL + "/weather" + "?" + v.Encode())
	if err != nil {
		return weather, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return weather, err
	}

	if err := json.Unmarshal(body, &weather); err != nil {
		return weather, err
	}

	return weather, nil
}

type Forecast struct {
	Cnt  int       `json:"cnt"`
	List []Weather `json:"list"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		}
	} `json:"city"`
	Country  string `json:"country"`
	Timezone int    `json:"timezone"`
}

func (a API) GetForecastData(name string, v url.Values) (weathers []Weather, err error) {
	if v == nil {
		v = url.Values{}
	}
	v.Set("q", name)
	v.Set("appid", a.APIKey)

	var forecast Forecast

	res, err := a.HTTPClient.Get(a.baseURL + "/forecast" + "?" + v.Encode())
	if err != nil {
		return weathers, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return weathers, err
	}

	if err := json.Unmarshal(body, &forecast); err != nil {
		return weathers, err
	}

	weathers = forecast.List
	for i := 0; i < forecast.Cnt; i++ {
		weathers[i].ID = forecast.City.ID
		weathers[i].Name = forecast.City.Name
		weathers[i].Coord = forecast.City.Coord
	}
	return weathers, nil
}
