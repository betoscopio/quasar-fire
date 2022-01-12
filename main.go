package main

import (
	"fmt"
	"net/http"
	"quasar-fire/server"
	"quasar-fire/utils"
)

func main() {
	fmt.Println("Operación 'Fuergo de Quasar'")

	sats := utils.Sats
	sats.InitSatellites() //Inicia satelites kenobi,skywalker y sato

	quasarFire := utils.Ship{
		Name:  "Quasar Fire",
		Point: utils.Point{X: -100, Y: 75.5},
	}

	fmt.Println(sats.Data[0])
	fmt.Println(sats.Data[1])
	fmt.Println(sats.Data[2])
	fmt.Println(quasarFire)

	r1 := quasarFire.Distance(sats.Data[0].Point)
	r2 := quasarFire.Distance(sats.Data[1].Point)
	r3 := quasarFire.Distance(sats.Data[2].Point)

	fmt.Println(sats)

	fmt.Printf("Distancia Quasar Fire a %v: %v\n", sats.Data[0].Name, r1)
	fmt.Printf("Distancia Quasar Fire a %v: %v\n", sats.Data[1].Name, r2)
	fmt.Printf("Distancia Quasar Fire a %v: %v\n", sats.Data[2].Name, r3)

	shipPosition := utils.Trilateration(sats.Data[0].Point, sats.Data[1].Point, sats.Data[2].Point, r1, r2, r3)
	fmt.Println("Quasar Fire se encuentra en: ", shipPosition)

	x, y := utils.GetLocation(float32(r1), float32(r2), float32(r3))

	fmt.Println("Quasar Fire se encuentra en: ", x, y)

	encode0 := utils.EncodeMessage("este es un mensaje", 0)
	encode1 := utils.EncodeMessage("este es un mensaje", 1)
	encode2 := utils.EncodeMessage("este es un mensaje", 2)

	fmt.Printf("\nMensajes codificados:\n")
	fmt.Println("encode0:", encode0)
	fmt.Println("encode1:", encode1)
	fmt.Println("encode2:", encode2)

	fmt.Printf("\nMensaje decodificado:\n")
	messages := [][]string{encode0, encode1, encode2}
	fmt.Println(utils.GetMessage(messages))

	//Defincion de handlers de URLs
	fmt.Println("..... ..... .....")
	fmt.Println("Iniciando servidor")
	///fmt.Println("Es posible visitar en http://localhost:8080/")

	satelliteHandlers := server.NewSatellitesHandlers()
	//messageHandlers := server.NewSignalHandlers()
	signalHandlers := server.NewSignalHandlers()
	//fmt.Println(*signalHandlers)

	http.HandleFunc("/satellites", satelliteHandlers.Get)
	http.HandleFunc("/topsecret", signalHandlers.Signals)

	http.HandleFunc("/", server.WelcomeHandler)

	//Definición del servidor
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
