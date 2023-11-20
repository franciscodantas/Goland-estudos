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

	for i, v := range pessoa1.sabores_favoritos{
		fmt.Println(i, v)
	}

	for i, v := range pessoa2.sabores_favoritos{
		fmt.Println(i, v)
	}
}