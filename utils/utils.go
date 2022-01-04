package utils

import (
	"math"
	"strings"
)

var Prueba int = 10

type Point struct {
	X float64
	Y float64
}

type Satellite struct {
	Name string
	Point
}

type Ship struct {
	Name string
	Point
}

//Calcula la distacia ente los puntos p y p2 usando la formula de distancia entre dos puntos
func (p Point) Distance(p2 Point) float64 {
	diffX := math.Pow(p2.X-p.X, 2)
	diffY := math.Pow(p2.Y-p.Y, 2)
	return math.Sqrt(diffX + diffY)
}

//Calcula distancia de dos puntos sin necesitar objeto
func DistancePoints(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
}

// Inicializa coordenadas de satelites Kenobi, Skywalker y Sato
func getSatellites() (Point, Point, Point) {
	kenobi := Point{X: -500, Y: -200}
	skywalker := Point{X: 100, Y: -100}
	sato := Point{X: 500, Y: 100}

	return kenobi, skywalker, sato
}

// Obtiene valor de norma  vectorial de un punto (vector 2D)
// https://en.wikipedia.org/wiki/Norm_(mathematics)
func norm(p Point) float64 {
	return math.Sqrt(math.Pow(p.X, 2) + math.Pow(p.Y, 2))
}

// Obtiene punto con el método trilateration
// https://en.wikipedia.org/wiki/True-range_multilateration
// https://es.wikipedia.org/wiki/Trilateraci%C3%B3n
// https://stackoverflow.com/questions/29656921/trilateration-2d-algorithm-implementation
// https://stackoverflow.com/questions/9747227/2d-trilateration
//
func Trilateration(p1 Point, p2 Point, p3 Point, r1 float64, r2 float64, r3 float64) Point {

	p2p1diff := p1.Distance(p2)
	ex := Point{X: (p2.X - p1.X) / p2p1diff, Y: (p2.Y - p1.Y) / p2p1diff}
	aux := Point{X: p3.X - p1.X, Y: p3.Y - p1.Y}
	//signed magnitude of the x component
	i := ex.X*aux.X + ex.Y*aux.Y
	//the unit vector in the y direction.
	aux2 := Point{X: p3.X - p1.X - i*ex.X, Y: p3.Y - p1.Y - i*ex.Y}
	ey := Point{X: aux2.X / norm(aux2), Y: aux2.Y / norm(aux2)}
	//the signed magnitude of the y component
	j := ey.X*aux.X + ex.Y*aux.Y
	//Obtención de coordenadas
	x := (math.Pow(r1, 2) - math.Pow(r2, 2) + math.Pow(p2p1diff, 2)) / (2 * p2p1diff)
	y := (math.Pow(r1, 2) - math.Pow(r3, 2) + math.Pow(i, 2) + math.Pow(j, 2)/(2*j) - (i*x)/j)

	finalX := p1.X + x*ex.X + y*ey.X
	finalY := p1.Y + x*ex.Y + y*ey.Y

	return Point{X: finalX, Y: finalY}
}

//Obtiene ubicación basandose en coordenadas
func GetLocation(r1 float32, r2 float32, r3 float32) (float32, float32) {
	//Obtiene coordenadas de los satelites
	kenobi, skywalker, sato := getSatellites()
	point := Trilateration(kenobi, skywalker, sato, float64(r1), float64(r2), float64(r3))
	return float32(point.X), float32(point.Y)
}

// Simula la generación de un mensaje con "ruido"
func EncodeMessage(s string, noise int) []string {
	//determinar cuanto se pierde segun distancia, reemplazar palabras por "", hacerlo con rangos
	splited := strings.Split(s, " ")
	n := len(splited)
	//fmt.Println("noise:", noise)
	//fmt.Println("n:", n)
	for i := range splited {
		if noise == 0 {
			break
		} else if noise == 1 {
			splited[1] = ""
			break
		} else if noise == 2 {
			splited[2] = ""
			break
		} else if n <= noise {
			//Definir lógica para cuando el ruido es muy grande
			//inserta caracteres vacíos al inicio
			//diff := noise - n - 1
		}
		splited[i] = ""
		noise--
	}
	return splited
}

// input: el mensaje tal cual es recibido en cada satélite
// output: el mensaje tal cual lo genera el emisor del mensaje
//func GetMessage(messages ...[]string) (msg string)
//Obtiene el mensaje descifrado usando una lista de strings
func GetMessage(messages [][]string) string {
	//decodificar mensaje sumando valores
	msg := "Este es el mensaje"
	return msg
}