package owmjp

import (
	"net/url"
	"os"
	"testing"
)

func TestCurrentWeather(t *testing.T) {
	Init()
	api := NewAPIWithKey(os.Getenv("OWM_KEY"))

	v := url.Values{}
	v.Set("units", "metric")

	_, err := api.GetCurrentWeatherData("Tokyo", v)
	if err != nil {
		t.Fatal(err)
	}
}

func TestForecast(t *testing.T) {
	Init()
	api := NewAPIWithKey(os.Getenv("OWM_KEY"))
	v := url.Values{}
	v.Set("units", "metric")

	_, err := api.GetForecastData("Tokyo", v)
	if err != nil {
		t.Fatal(err)
	}
}
