package weather

import "testing"

func TestLoadFromJSON(t *testing.T) {
	LoadFromJSON("weather_data.json")
}
