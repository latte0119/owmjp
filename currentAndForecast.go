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
