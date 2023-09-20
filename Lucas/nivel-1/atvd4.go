package main
import("fmt")

func atvd4(){
	type fixpay int
	var x fixpay

	fmt.Print(x)
	fmt.Printf("\n%T\n",x)
	x = 42
}