// Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço".
// O campo "endereço" deve ser uma outra struct com os campos "rua", "número", "cidade" e "estado".
// Escreva uma função que receba uma Pessoa como parâmetro e imprima seu endereço completo.
package main

import "fmt"

type Endereco struct {
	rua    string
	cidade string
	estado string
}
type Pessoa struct {
	nome  string
	idade int
	Endereco
}

func main() {
	p := Pessoa{
		nome:  "Gabriel",
		idade: 18,
		Endereco: Endereco{
			rua:    "567",
			cidade: "Brasilia",
			estado: "DF",
		},
	}
	fmt.Println(p.rua, p.cidade, p.estado)
}
