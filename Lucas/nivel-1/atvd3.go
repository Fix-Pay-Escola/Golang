package main

import(
	"fmt"
)
func atvd3(){
	x := 42
	y := "James Bond"
	z := true

	s := fmt.Sprintf("%v,%v,%v",x,y,z)

	fmt.Print(s)
}