package main

import "fmt"

func main()  {
	mapa := map[string][]string{
		"ana": {"costurar", "conversar"},
		"paulo": {"beber"},
	}
	mapa["francisco"] = []string{"jogar"}
	
	for k, v := range mapa {
		fmt.Printf("%s: %s\n", k, v[:])
	}
}