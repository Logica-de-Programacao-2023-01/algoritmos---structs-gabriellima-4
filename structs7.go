// Crie uma struct chamada Animal com os campos "nome", "espécie", "idade" e "som".
// Escreva funções que permitam modificar o som que o animal faz e
// uma função que imprima as informações do animal e o som que ele faz.
package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	Especie string
	Idade   int
	Som     string
}

func soundChange(a Animal) string {
	som := a.Som
	novoSom := strings.ReplaceAll(a.Som, som, "AAAAAAAAAAAAA")

	return novoSom
}

func main() {
	l := Animal{
		Especie: "Canis Lupus",
		Idade:   5,
	}

	an := soundChange(Animal{
		Especie: "Canis Lupus",
		Idade:   5,
		Som:     "woof",
	})
	fmt.Println(l, an)

}
