package main

import (
	"fmt"
)

var x int = 10
var y string = "jamesbond"
var z bool = true

func atvd3() {

	s:=fmt.Sprintf("%v,\t%v,\t%v",x, y, z)
	fmt.Println(s)
}
