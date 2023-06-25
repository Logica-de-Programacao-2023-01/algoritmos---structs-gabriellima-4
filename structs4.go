// Crie uma struct chamada Playlist com os campos "nome" e "músicas".
// O campo "músicas" deve ser um slice de structs,
// cada uma representando uma música com os campos "título", "artista" e "duração".
// Escreva uma função que receba uma Playlist como parâmetro e imprima o nome da playlist,
// o tempo total da playlist e as informações de cada música.
package main

import "fmt"

type Playlist struct {
	Nome    string
	Musicas []Musica
}

type Musica struct {
	Titulo  string
	Artista string
	Duracao float64
}

func tempoDeDuracao(p Playlist) float64 {
	somaTempo := 0.0

	for _, musica := range p.Musicas {
		somaTempo += musica.Duracao
	}
	return somaTempo
}

func main() {
	playlist := Playlist{
		Nome: "CRISTIANO DE BICICLETA MINHA NOSSA",
		Musicas: []Musica{
			{
				Titulo:  "HIGHEST IN THE ROOM",
				Artista: "Travis Scott",
				Duracao: 2.58,
			},
			{
				Titulo:  "21 Questions",
				Artista: "50 Cent",
				Duracao: 3.44,
			},
			{
				Titulo:  "XSCAPE",
				Artista: "Don Toliver",
				Duracao: 2.36,
			},
			{
				Titulo:  "Hold The Line",
				Artista: "TOTO",
				Duracao: 3.56,
			},
			{
				Titulo:  "If you Leave Me Now",
				Artista: "Chicago",
				Duracao: 3.55,
			},
			{
				Titulo:  "Deixa em Off - Ao Vivo",
				Artista: "Turma do Pagode",
				Duracao: 2.53,
			},
			{
				Titulo:  "Around Me",
				Artista: "Metro Boomin ft. Don Toliver",
				Duracao: 3.12,
			},
			{
				Titulo:  "Push It To The Limit (Scarface)",
				Artista: "Paul Engemann",
				Duracao: 3.44,
			},
		},
	}
	fmt.Println(tempoDeDuracao(playlist))
}
