package server

import (
	"encoding/json"
	"net/http"
	"quasar-fire/utils"
)

type SatelliteHandlers struct {
	store map[string]utils.Satellite
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<html><h1>Bienvenidos a la Operación Fuego Quasar</h1></html>"))
}

func (h *SatelliteHandlers) Get(w http.ResponseWriter, r *http.Request) {
	sats := make([]utils.Satellite, len(h.store))

	i := 0
	for _, sat := range h.store {
		sats[i] = sat
		i++
	}

	jsonBytes, err := json.Marshal(sats)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func NewSatellitesHandlers() *SatelliteHandlers {

	sats := utils.InitSatellites()
	return &SatelliteHandlers{
		store: map[string]utils.Satellite{
			sats[0].Name: sats[0],
			sats[1].Name: sats[1],
			sats[2].Name: sats[2],
		},
	}
}
