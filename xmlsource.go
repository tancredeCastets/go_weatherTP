package weather

import (
	"encoding/xml"
	"os"
	"time"
)

type rootXML struct {
	StationXML []stationXML `xml:"station"`
}

type stationXML struct {
	XMLName  xml.Name `xml:"station"`
	ID       string   `xml:"id,attr"`
	Name     string   `xml:"name"`
	Country  string   `xml:"country,attr"`
	Location struct {
		Latitude  float64 `xml:"latitude,attr"`
		Longitude float64 `xml:"longitude,attr"`
	}
	Altitude    int              `xml:"altitude,attr"`
	Device      deviceXML        `xml:"device"`
	Observation []observationXML `xml:"observations"`
}

type deviceXML struct {
	Type string `xml:"type,attr"`
}

type observationXML struct {
	Timestamp   time.Time `xml:"timestamp"`
	Temperature float64   `xml:"temperature_celsius,attr"` // celsius
	Conditions  string    `xml:"conditions,attr"`
	Wind        windXML   `xml:"wind"`
}
type windXML struct {
	Speed     float64 `xml:"speed_kmh,attr"`     // km/h
	Direction int     `xml:"direction_deg,attr"` // Degrees
}

func toStationXML(S stationXML) Station {
	obs := make([]Observation, 0)
	for _, o := range S.Observation {
		obs = append(obs, toObservationXML(o))
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
func toObservationXML(O observationXML) Observation {
	return Observation{
		Timestamp:   O.Timestamp,
		Temperature: O.Temperature,
		Conditions:  O.Conditions,
		Wind: Wind{
			Speed:     O.Wind.Speed,
			Direction: O.Wind.Direction,
		},
	}
}

func LoadFromXML(path string) ([]Station, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var root rootXML
	if err := xml.Unmarshal(f, &root); err != nil {
		return nil, err
	}
	result := make([]Station, 0)
	for _, s := range root.StationXML {
		result = append(result, toStationXML(s))
	}
	return result, nil
}
