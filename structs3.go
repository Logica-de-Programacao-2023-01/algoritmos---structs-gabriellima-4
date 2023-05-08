// Crie uma struct chamada Triângulo com os campos "base" e "altura".
// Escreva uma função que receba um Triângulo como parâmetro e calcule a área do triângulo (área = base * altura / 2).
package main

import "fmt"

type Triângulo struct {
	Base   int
	Altura int
}

func calcTriangulo(t Triângulo) int {
	a := t.Base * t.Altura / 2
	return a
}

func main() {
	r := calcTriangulo(Triângulo{
		Base:   8,
		Altura: 3,
	})
	fmt.Println(r)
}
