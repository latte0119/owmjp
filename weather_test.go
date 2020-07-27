package owmjp

import (
	"os"
	"testing"
)

func TestMainJP(t *testing.T) {
	Init()
	api := NewAPIWithKey(os.Getenv("OWM_KEY"))
	w, _ := api.GetCurrentWeatherData("Tokyo", nil)
	t.Fatal(w.MainJP())
}
