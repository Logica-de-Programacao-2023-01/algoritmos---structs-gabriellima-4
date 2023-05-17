// Crie uma struct chamada Viagem com os campos "origem", "destino", "data" e "preço".
// Escreva uma função que receba um slice de Viagens como parâmetro e retorne a viagem mais cara.
package main

import "fmt"

type viagem struct {
	Origem  string
	Destino string
	Data    int
	Preco   float64
}

func ViagemMaisCara(viagens []viagem) viagem {
	var ViagemPraRico viagem
	for _, viagem := range viagens {
		if viagem.Preco > ViagemPraRico.Preco {
			ViagemPraRico = viagem
		}
	}
	return ViagemPraRico
}

func main() {
	BSBpRJ := viagem{
		Origem:  "Brasília",
		Destino: "Rio de Janeiro",
		Data:    03 / 06,
		Preco:   1900,
	}
	BSBpTUR := viagem{
		Origem:  "Brasília",
		Destino: "Istanbul - Turquia",
		Data:    10 / 06,
		Preco:   7100,
	}
	BSBpBA := viagem{
		Origem:  "Brasília",
		Destino: "Salvador",
		Data:    11 / 07,
		Preco:   1500,
	}

	opcoesViagem := []viagem{BSBpRJ, BSBpBA, BSBpTUR}

	fmt.Println("A viagem mais cara é:", ViagemMaisCara(opcoesViagem))
}
