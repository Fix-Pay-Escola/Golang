package main

import (
	"fmt"
)

func atvd3() {
	x := 2006
	for {
		if x <= 2029 {
			fmt.Print(x, "\n")
			x++
		} else {
			break
		}
	}
}
