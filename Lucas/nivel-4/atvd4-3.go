package main

import (
	"fmt"
)

func atvd3() {
	slice := []int{10, 20, 30, 40, 50, 1, 2, 3, 4, 5}

	for i, v := range slice {
		fmt.Print(i, v)
	}
	fmt.Printf("\n%T", slice)
}
