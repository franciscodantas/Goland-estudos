package main

import "fmt"

var x = 43
var y = "James Bond"
var z = true

func main(){
	s := fmt.Sprintf("%v\n%v\n%v",x,y,z)
	fmt.Println(s)
}