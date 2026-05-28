package weather

import (
	"encoding/json"
	"net/http"
)

type App struct{ store *Store }

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func (a *App) ListStations(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, a.store.All())
}
func NewApp(store *Store) *App {
	return &App{store: store}
}
