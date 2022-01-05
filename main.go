package main

import (
	"fmt"
	"net/http"
	"quasar-fire/server"
	utils "quasar-fire/utils"
)

func main() {
	//sat := make([]utils.Satellite, 0)

	fmt.Println("Operación 'Fuergo de Quasar'")

	sat := utils.InitSatellites()

	kenobi := utils.Satellite{
		Name:  "Kenobi",
		Point: utils.Point{X: -500, Y: -200},
	}

	skywalker := utils.Satellite{
		Name:  "Skywalker",
		Point: utils.Point{X: 100, Y: -100},
	}

	sato := utils.Satellite{
		Name:  "Sato",
		Point: utils.Point{X: 500, Y: 100},
	}

	quasarFire := utils.Ship{
		Name:  "Quasar Fire",
		Point: utils.Point{X: -100, Y: 75.5},
	}

	fmt.Println("El valor de prueba es:", utils.Prueba)
	fmt.Println(kenobi)
	fmt.Println(skywalker)
	fmt.Println(sato)
	fmt.Println(quasarFire)

	r1 := quasarFire.Distance(sat[0].Point)
	r2 := quasarFire.Distance(sat[1].Point)
	r3 := quasarFire.Distance(sat[2].Point)

	fmt.Println()
	/*
		fmt.Println("Distancia Quasar Fire a :", r1)
		fmt.Println("Distancia Quasar Fire a Skywalker:", r2)
		fmt.Println("Distancia Quasar Fire a Sato:", r3)
	*/
	fmt.Printf("Distancia Quasar Fire a %v: %v\n", sat[0].Name, r1)
	fmt.Printf("Distancia Quasar Fire a %v: %v\n", sat[1].Name, r2)
	fmt.Printf("Distancia Quasar Fire a %v: %v\n", sat[2].Name, r3)

	fmt.Println()
	fmt.Println("Distancia Quasar Fire a Kenobi 2:", utils.DistancePoints(quasarFire.Point, kenobi.Point))
	fmt.Println("Distancia Quasar Fire a Skywalker 2:", utils.DistancePoints(quasarFire.Point, skywalker.Point))
	fmt.Println("Distancia Quasar Fire a Sato 2:", utils.DistancePoints(quasarFire.Point, sato.Point))

	shipPosition := utils.Trilateration(kenobi.Point, skywalker.Point, sato.Point, r1, r2, r3)
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

	//fmt.Println(sat)

	//Defincion de handlers de URLs
	fmt.Println("..... ..... .....")
	fmt.Println("Iniciando servidor")
	fmt.Println("Es posible visitar en http://localhost:8080/")

	satelliteHandlers := server.NewSatellitesHandlers()
	http.HandleFunc("/satellites", satelliteHandlers.Get)

	http.HandleFunc("/", server.WelcomeHandler)

	//Definición del servidor
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
