package main

import "fmt"

var y = 20
var x = 10

func main()  {
	igual := x == y
	diferente := x != y
	menor := x < y
	menor_igual := x <= y
	maior := x > y
	maior_igual := x >= y
	fmt.Println(igual, diferente, menor, menor_igual, maior, maior_igual)
}