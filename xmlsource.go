package weather

import (
	"encoding/xml"
	"os"
	"strconv"
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
		Latitude  float64 `xml:"lat,attr"`
		Longitude float64 `xml:"lon,attr"`
		Altitude  int     `xml:"altitude,attr"`
	} `xml:"coordinates"`
	Device struct {
		Type string `xml:"model,attr"`
	} `xml:"hardware"`
	Observations struct {
		Items []observationXML `xml:"observation"`
	} `xml:"observations"`
}
type observationXML struct {
	XMLName     xml.Name     `xml:"observation"`
	Timestamp   time.Time    `xml:"at,attr"`
	Measures    []measureXML `xml:"measure"`
	Temperature float64      `xml:"temperature_celsius,attr"` // celsius
	Conditions  string       `xml:"sky,attr"`
	Note        *string      `xml:"note"`
	Wind        struct {
		XMLName   xml.Name `xml:"wind"`
		Speed     float64  `xml:"speed,attr"`     // km/h
		Direction int      `xml:"direction,attr"` // Degrees
	}
}

type measureXML struct {
	Type  string `xml:"type,attr"`
	Value string `xml:",chardata"`
}

func toStationXML(S stationXML) Station {
	obs := make([]Observation, 0)
	for _, o := range S.Observations.Items {
		obs = append(obs, toObservationXML(o))
	}
	return Station{
		ID:      S.ID,
		Name:    S.Name,
		Country: S.Country,
		Location: Location{
			Latitude:  S.Location.Latitude,
			Longitude: S.Location.Longitude,
		},
		Altitude: S.Location.Altitude,
		Device: Device{
			Type: S.Device.Type,
		},
		Observation: obs,
	}

}
func toObservationXML(O observationXML) Observation {
	var temp float64
	for _, m := range O.Measures {
		switch m.Type {
		case "temperature":
			temp, _ = strconv.ParseFloat(m.Value, 64)
		}
	}
	return Observation{
		Timestamp:   O.Timestamp,
		Temperature: temp,
		Conditions:  O.Conditions,
		Notes:       O.Note,
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
