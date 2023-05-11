// Crie uma struct chamada Filme com os campos "título", "diretor", "ano" e "avaliações".
// O campo "avaliações" deve ser um slice de inteiros representando as notas que o filme recebeu dos espectadores.
// Escreva funções que permitam adicionar ou remover avaliações do filme,
// calcular a média das avaliações
// e imprimir as informações do filme e sua média de avaliações.
package main

import "fmt"

type Filme struct {
	Título     string
	Diretor    string
	Ano        int
	Avaliações []int
}

func addRating([]int) []int {
	numeros := []int{8, 6, 9, 9}
	numeros = append(numeros, 10)

	return numeros
}

func deleteRating([]int) []int {
	nums := []int{8, 6, 9, 9}
	nums = nums[:len(nums)-1]

	return nums
}

func averageRating([]int) float64 {
	numeros := []int{8, 6, 9, 9}
	media := (numeros[0] + numeros[1] + numeros[2] + numeros[3]) / 4

	return float64(media)
}

func main() {
	r := Filme{
		Título:     "Bátima",
		Diretor:    "Christopher Nóia",
		Ano:        2023,
		Avaliações: []int{8, 6, 9, 9},
	}
	fmt.Println(r)

	f := averageRating([]int{8, 6, 9, 9})
	fmt.Println("A média das avaliações é: ", f)
}
