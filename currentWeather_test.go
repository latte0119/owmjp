package owmjp

import (
	"net/url"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestCurrentWeather(t *testing.T) {
	godotenv.Load("envfiles/.env")
	api := NewAPIWithKey(os.Getenv("OWM_KEY"))

	v := url.Values{}
	v.Set("units", "metric")

	_, err := api.GetCurrentWeatherData("Tokyo", v)
	if err != nil {
		t.Fatal(err)
	}
}

func TestForecast(t *testing.T) {
	godotenv.Load("envfiles/.env")
	api := NewAPIWithKey(os.Getenv("OWM_KEY"))
	v := url.Values{}
	v.Set("units", "metric")

	_, err := api.GetForecastData("Tokyo", v)
	if err != nil {
		t.Fatal(err)
	}
}
