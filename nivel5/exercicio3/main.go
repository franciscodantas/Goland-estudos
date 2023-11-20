package main

import "fmt"

type veiculo struct{
	portas int
	cor string
}

type caminhonete struct{
	veiculo
	quatroRodas bool
}

type sedan struct{
	veiculo
	modeloLuxo bool
}

func main()  {
	caminhonete1 := caminhonete{veiculo{2, "preto"}, true}
	sedan1 := sedan{veiculo{4, "branco"}, false}

	fmt.Println(caminhonete1)
	fmt.Println(sedan1)
	fmt.Println(sedan1.modeloLuxo)
	fmt.Println(caminhonete1.quatroRodas)
}