package weather

import "testing"

func TestFilterByCountryJSON(t *testing.T) {
	stations, _ := LoadFromJSON("weather_data.json")
	result := FilterByCountry(stations, "FR")
	t.Logf("stations ES: %d", len(result))
}

func TestFilterByCountryXML(t *testing.T) {
	stations, _ := LoadFromXML("weather_data.xml")
	result := FilterByCountry(stations, "FR")
	t.Logf("stations ES: %d", len(result))
}

func TestAvgTemperatureJSON(t *testing.T) {
	stations, _ := LoadFromJSON("weather_data.json")
	avg := AvgTemperature(stations[8])
	t.Logf("avg temp station JSON: %.2f", avg)
}

func TestAvgTemperatureXML(t *testing.T) {
	stations, _ := LoadFromXML("weather_data.xml")
	avg := AvgTemperature(stations[4])
	t.Logf("avg temp station XML: %.2f", avg)
}

func TestMaxWindGustJSON(t *testing.T) {
	stations, _ := LoadFromJSON("weather_data.json")
	station, speed := MaxWindGust(stations)
	t.Logf("station: %s, vitesse max: %.2f", station.Name, speed)
}

func TestMaxWindGustXML(t *testing.T) {
	stations, _ := LoadFromXML("weather_data.xml")
	station, speed := MaxWindGust(stations)
	t.Logf("station: %s, vitesse max: %.2f", station.Name, speed)
}
