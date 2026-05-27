package weather

import "time"

type Root struct {
	Stations []Station
}
type Station struct {
	ID          string
	Name        string
	Country     string
	Location    Location
	Altitude    int
	Device      Device
	Observation []Observation
}

type Location struct {
	Latitude  float64
	Longitude float64
}

type Device struct {
	Type string
}

type Observation struct {
	Timestamp   time.Time
	Temperature float64 // celcius
	Conditions  string
	Wind        Wind
}
type Wind struct {
	Speed     float64 // km/h
	Direction int     // Degrees
}
