package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	sabores_favoritos []string
}

func main()  {
	pessoa1 := pessoa{
		nome: "ana",
		sobrenome: "silva",
		sabores_favoritos: []string{"morango", "chocolate"},
	}
	pessoa2 := pessoa{
		nome: "kauã",
		sobrenome: "abreu",
		sabores_favoritos: []string{"maca verde"},
	}

	mapa := make(map[string]pessoa)

	mapa[pessoa1.sobrenome] = pessoa1
	mapa[pessoa2.sobrenome] = pessoa2

	for _, v := range mapa {
		fmt.Println(v)
	}

	for _, v := range mapa {
		fmt.Println(v.nome)
		for _, sabor := range v.sabores_favoritos{
			fmt.Printf("\t %s\n", sabor)
		}
	}
}