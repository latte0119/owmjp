package owmjp

import (
	"fmt"
	"net/url"
	"testing"
)

func TestCurrentWeather(t *testing.T) {
	api := NewAPIWithAPIKey("d1dcf4b83f0b4d9a5e0dbc0a059028a2")

	v := url.Values{}
	v.Set("units", "metric")

	weather, err := api.GetCurrentWeatherData("Tokyo", v)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(weather)
}
