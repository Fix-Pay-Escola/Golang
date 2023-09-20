package main

import (
	"fmt"
)

func atvd5() {
	type fixpay int
	var x fixpay
	var y int

	fmt.Print(x)
	fmt.Printf("\n%T\n", x)
	x = 42
	fmt.Print(x)
	y = int(x)
	fmt.Print("\n", y)

}
