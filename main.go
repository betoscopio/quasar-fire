package main

import (
	"fmt"
	utils "quasar-fire/utils"
)

func main() {
	fmt.Println("Operación 'Fuergo de Quasar'")

	kenobi := utils.Satellite{
		Name: "Kenobi",
		X:    -500,
		Y:    -200,
	}

	skywalker := utils.Satellite{
		Name: "Skywalker",
		X:    100,
		Y:    -100,
	}

	sato := utils.Satellite{
		Name: "Sato",
		X:    500,
		Y:    100,
	}

	fmt.Println("El valor de prueba es:", utils.Prueba)
	fmt.Println(kenobi)
	fmt.Println(skywalker)
	fmt.Println(sato)
}
