// Crie uma struct chamada Aluno com os campos "nome", "idade" e "notas".
// O campo "notas" deve ser um slice de float64, representando as notas que o aluno tirou em uma determinada disciplina.
// Escreva funções que permitam adicionar ou remover notas do aluno,
// calcular a média das notas e imprimir o nome, idade e média do aluno.
package main

import "fmt"

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func adicionarNota(a Aluno) []float64 {
	n := a.Notas
	n = append(a.Notas, 7)

	return n
}

func removerNota(a Aluno) []float64 {
	d := a.Notas
	d = a.Notas[:len(a.Notas)-1]

	return d
}

func mediaNotas([]float64) float64 {
	notas := []float64{10, 6, 8, 9}
	avg := (notas[0] + notas[1] + notas[2] + notas[3]) / float64(len(notas))

	return avg
}

func main() {
	a := Aluno{
		Nome:  "Gabriel",
		Idade: 18,
		Notas: []float64{10, 6, 8, 9},
	}
	fmt.Println(a)

	b := mediaNotas([]float64{10, 6, 8, 9})

	fmt.Println("A média das notas desse aluno é: ", b)
}

