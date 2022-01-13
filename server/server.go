package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"quasar-fire/utils"
	"strings"
	"sync"
)

type SatelliteHandlers struct {
	store map[string]utils.Satellite
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func NewSatellitesHandlers() *SatelliteHandlers {

	sats := utils.Sats
	return &SatelliteHandlers{
		store: map[string]utils.Satellite{
			sats.Data[0].Name: sats.Data[0],
			sats.Data[1].Name: sats.Data[1],
			sats.Data[2].Name: sats.Data[2],
		},
	}
}

//Define tipo Signal
type Signal struct {
	Name     string   `json: "name"`
	Distance float64  `json: "distance"`
	Message  []string `json: "message"`
}

//Define estructura para almacenar y acceder a Signals
type SignalHandlers struct {
	sync.Mutex // Mutex embed
	Data       []Signal
}

//Crea nuevo manejador de objetos Signal
func NewSignalHandlers() *SignalHandlers {
	signals := []Signal{
		Signal{
			Name:     "kenobi",
			Distance: 100.0,
			Message:  []string{"este", "", "", "mensaje", ""},
		},
		Signal{
			Name:     "skywalker",
			Distance: 115.5,
			Message:  []string{"", "es", "", "", "secreto"},
		},
		Signal{
			Name:     "sato",
			Distance: 142.7,
			Message:  []string{"este", "", "un", "", ""},
		},
	}

	//fmt.Println(signals)

	return &SignalHandlers{
		Data: signals,
	}

}

//Obtiene los objetos Signal ya ingresados
func (h *SignalHandlers) get(w http.ResponseWriter, r *http.Request) {
	var signals []Signal

	//Bloquea ingreso de nuevos datos para obtener valores actuales
	h.Lock()
	for _, signal := range h.Data {
		//fmt.Println(reflect.TypeOf(signal), signal)
		signals = append(signals, signal)
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(signals)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

//Agrega un objeto Signal al listado
func (h *SignalHandlers) post(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//validates the content-type from the post
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var signal Signal
	err = json.Unmarshal(bodyBytes, &signal)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	h.Lock()
	h.Data = append(h.Data, signal)
	h.Unlock()
}

//Obtiene listado de satelites ingresados
func (h *SignalHandlers) GetSatellites() {
	//TODO: implementar con objeto global o parámetro
	//h.Sats = append(h.Sats, sats)
}

//TODO: Agrega multiples objetos Signal al listado.
func (h *SignalHandlers) postSignals(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//validates the content-type from the post
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	fmt.Println(bodyBytes)
	var signals []Signal
	err = json.Unmarshal(bodyBytes, &signals)
	//fmt.Println(signals) //debug decoded json
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		w.Write([]byte(err.Error()))
		return
	}

	/*
		h.Lock()
		h.Data = append(h.Data, signal)
		h.Unlock()
	*/
}

//Lista o agrega un objeto signal según método HTTP usado: GET, POST
func (h *SignalHandlers) Signals(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método no permitido."))
	}
}

func (h *SignalHandlers) getDistance(w http.ResponseWriter, r *http.Request) {

	//Obtener distancias y mensajes de h si existen

	r1 := 100.0
	r2 := 115.5
	r3 := 142.7

	x, y := utils.GetLocation(float32(r1), float32(r2), float32(r3))

	m1 := []string{"este", "", "", "mensaje", ""}
	m2 := []string{"", "es", "", "", "secreto"}
	m3 := []string{"este", "", "un", "", ""}

	messages := [][]string{m1, m2, m3}
	message := utils.GetMessage(messages)

	//En caso que no existan suficientes puntos
	if len(messages) < 3 {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404: Not Found."))
	}

	pos := utils.SecretPosition{
		Point:   utils.Point{X: float64(x), Y: float64(y)},
		Message: message,
	}

	jsonBytes, err := json.Marshal(pos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

//Interactúa con REST API usando múltiples objetos.
//POST: los guarda en SignalHandlers
//GET: intenta calcular la ubicación de la nava si hay datos suficientes
func (h *SignalHandlers) SignalsMultiple(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.getDistance(w, r)
		return
	case "POST":
		h.postSignals(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método no permitido."))
	}
}

//TODO: Agrega multiples objetos Signal al listado.
func (h *SignalHandlers) SignalSat(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")

	if len(parts) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	sat_name := "kenobi"
	//compara parametro
	if parts[2] == sat_name {
		h.Signals(w, r)
		return
	}

}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<html><h1>Bienvenidos a la Operación Fuego Quasar</h1>
					<p>Encuentre la ubicación de origen del mensaje.</p>
					<p>Para mas detalles consultar en <a href="https://github.com/betoscopio/quasar-fire/">página del proyecto</a>.</p>
					</html>`))
}
