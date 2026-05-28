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

func (a *App) GetStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	st, ok := a.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "station not found")
	} else {
		writeJSON(w, http.StatusOK, st)
	}
}
func (a *App) CreateStation(w http.ResponseWriter, r *http.Request) {
	var st Station
	err := json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid station")
		return
	}
	if a.store.Has(st.ID) {
		writeError(w, http.StatusConflict, "station is already created")
		return
	} else {
		a.store.Put(st)
		writeJSON(w, http.StatusCreated, st)
	}
}

func (a *App) UpdateStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var st Station
	err := json.NewDecoder(r.Body).Decode(&st)
	st.ID = id
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid station")
		return
	}
	if a.store.Has(st.ID) {
		a.store.Put(st)
		writeJSON(w, http.StatusOK, st)
	} else {
		a.store.Put(st)
		writeJSON(w, http.StatusCreated, st)
	}
}
func (a *App) DeleteStation(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	st, ok := a.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "station not found")
		return
	} else {
		a.store.Delete(id)
		writeJSON(w, http.StatusNoContent, st)
	}
}
