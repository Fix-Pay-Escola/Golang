package main

import (
	"fmt"
)

func atvd5() {
	const (
		ano = 2020
		a   = ano + iota
		b
		c
		d
	)
	fmt.Printf("ano atual:%v\nproximos anos:\n%v\n%v\n%v", ano, a, b, c)
}
