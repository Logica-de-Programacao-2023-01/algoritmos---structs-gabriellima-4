// Crie uma struct chamada Funcionário com os campos "nome", "salário" e "idade".
// Escreva funções que permitam aumentar ou diminuir o salário do funcionário em uma determinada porcentagem
// e uma função que calcule o tempo de serviço do funcionário (considerando que ele começou a trabalhar aos 18 anos).
package main

import "fmt"

type Funcionário struct {
	Nome    string
	Salario float64
	Idade   int
}

func salaryChange(f Funcionário) float64 {

	newSalary := f.Salario + f.Salario*0.10

	return newSalary
}

func timeService(f Funcionário) int {
	tempoS := f.Idade - 18
	if f.Idade < 18 {
		return 0
	}
	return tempoS
}

func main() {
	s := salaryChange(Funcionário{
		Nome:    "Juvisnélio",
		Salario: 10000,
		Idade:   45,
	})
	fmt.Println("O novo salário é: ", s)

	t := timeService(Funcionário{
		Nome:  "Juvisnélio",
		Idade: 45,
	})
	fmt.Println("O tempo de serviço do funcionário é: ", t, "anos")
}
