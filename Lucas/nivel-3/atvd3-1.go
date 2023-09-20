package main

import (
	"fmt"
)

func atvd1() {
	for i := 65; i <= 90; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
