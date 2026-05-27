package weather

import (
	"encoding/json"
	"os"
	"time"
)

type rootJSON struct {
	StationJSON []stationJSON `json:"stations"`
}

type stationJSON struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Country     string            `json:"country"`
	Location    locationJSON      `json:"location"`
	Altitude    int               `json:"altitude_m"`
	Device      deviceJSON        `json:"device"`
	Observation []observationJSON `json:"observations"`
}
type locationJSON struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type deviceJSON struct {
	Type string `json:"type"`
}

type observationJSON struct {
	Timestamp   time.Time `json:"timestamp"`
	Temperature float64   `json:"temperature_celsius"` // celsius
	Conditions  string    `json:"conditions"`
	Notes       *string   `json:"notes"`
	Wind        windJSON  `json:"wind"`
}
type windJSON struct {
	Speed     float64 `json:"speed_kmh"`     // km/h
	Direction int     `json:"direction_deg"` // Degrees
}

// conversion map
var countryISO = map[string]string{
	"France":             "FR",
	"Italy":              "IT",
	"Espagne":            "ES",
	"Allemagne":          "DE",
	"Portugal":           "PT",
	"Belgique":           "BE",
	"Pays-bas":           "NL",
	"Suisse":             "CH",
	"Autriche":           "AT",
	"Pologne":            "PL",
	"Danemark":           "DK",
	"Suède":              "SE",
	"Norvège":            "NO",
	"République tchèque": "CZ",
}

func toStation(S stationJSON) Station {
	obs := make([]Observation, 0)
	for _, o := range S.Observation {
		obs = append(obs, toObservation(o))
	}

	return Station{
		ID:      S.ID,
		Name:    S.Name,
		Country: countryISO[S.Country],
		Location: Location{
			Latitude:  S.Location.Latitude,
			Longitude: S.Location.Longitude,
		},
		Altitude: S.Altitude,
		Device: Device{
			Type: S.Device.Type,
		},
		Observation: obs,
	}
}

func toObservation(O observationJSON) Observation {
	return Observation{
		Timestamp:   O.Timestamp,
		Temperature: O.Temperature,
		Conditions:  O.Conditions,
		Notes:       O.Notes,
		Wind: Wind{
			Speed:     O.Wind.Speed,
			Direction: O.Wind.Direction,
		},
	}
}

func LoadFromJSON(path string) ([]Station, error) {

	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var root rootJSON
	if err := json.Unmarshal(f, &root); err != nil {
		return nil, err
	}
	result := make([]Station, 0)
	for _, s := range root.StationJSON {
		result = append(result, toStation(s))
	}
	return result, nil
}
