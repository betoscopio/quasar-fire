package main

import (
	"fmt"
	utils "quasar-fire/utils"
)

func main() {
	fmt.Println("Operación 'Fuergo de Quasar'")

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
		Point: utils.Point{X: 0, Y: 0},
	}

	fmt.Println("El valor de prueba es:", utils.Prueba)
	fmt.Println(kenobi)
	fmt.Println(skywalker)
	fmt.Println(sato)
	fmt.Println(quasarFire)

	r1 := quasarFire.Distance(kenobi.Point)
	r2 := quasarFire.Distance(skywalker.Point)
	r3 := quasarFire.Distance(sato.Point)
	//d := kenobi.Distance(skywalker.Point)

	fmt.Println()
	fmt.Println("Distancia Quasar Fire a Kenobi:", r1)
	fmt.Println("Distancia Quasar Fire a Skywalker:", r2)
	fmt.Println("Distancia Quasar Fire a Sato:", r3)

	shipPosition := utils.Trilateration(kenobi.Point, skywalker.Point, sato.Point, r1, r2, r3)
	fmt.Println("Quasar Fire se encuentra en: ", shipPosition)
}
