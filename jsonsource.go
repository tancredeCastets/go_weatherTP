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
	Wind        windJSON  `json:"wind"`
}
type windJSON struct {
	Speed     float64 `json:"speed_kmh"`     // km/h
	Direction int     `json:"direction_deg"` // Degrees
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

	return nil, nil
}
