package main

import (
	"fmt"
)

func atvd2() {
	x := 10 == 20
	y := 10 != 20
	z := 10 < 20
	a := 10 > 20
	b := 10 <= 20
	c := 10 >= 20

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v", x, y, z, a, b, c)
}
