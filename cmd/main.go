package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/efrei/weather"
)

func main() {

	stations, err := weather.LoadFromJSON("weather_data.json")
	if err != nil {
		log.Fatal(err)
	}
	store := weather.NewStore()
	for _, s := range stations {
		store.Put(s)
	}

	app := weather.NewApp(store)

	log.Printf("bootstrap : %d stations chargées", len(stations))

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	mux.HandleFunc("GET /stations", app.ListStations)
	mux.HandleFunc("GET /stations/{id}", app.GetStation)
	mux.HandleFunc("POST /stations", app.CreateStation)
	http.ListenAndServe(":8080", mux)
}
