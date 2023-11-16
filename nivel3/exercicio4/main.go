package main

import "fmt"

func main()  {
	ano := 2003
	for {
		if ano > 2023 {
			break
		}
		fmt.Println(ano)
		ano++
	}
}