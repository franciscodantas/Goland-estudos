package main

import "fmt"

func main()  {
	structanonimo := struct{
		mapa map[string]string
		lista []int
	}{
		mapa: map[string]string{"ana": "anão"},
		lista: []int{1, 2, 3, 4, 5},
	}

	fmt.Println(structanonimo)
}