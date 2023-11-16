package main

import "fmt"

const x int = 10
const y = 10

func main()  {
	fmt.Printf("%T%v\n", x, x)
	fmt.Printf("%T%v\n", y, y)
}