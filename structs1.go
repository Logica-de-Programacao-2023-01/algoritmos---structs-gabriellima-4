// Crie uma struct chamada Circulo com o campo "raio".
// Escreva uma função que receba um Circulo como parâmetro e calcule a área do círculo (área = pi * raio * raio).
// Dica: use a constante math.Pi para representar o número pi.
package main

import (
	"fmt"
	"math"
)

type Círculo struct {
	Raio float64
}

func calcCírculo(c Círculo) float64 {
	a := math.Pi * c.Raio * c.Raio
	return a
}

func main() {
	r := calcCírculo(Círculo{10})
	fmt.Println(r)
}
