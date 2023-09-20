package main
import ("fmt")

func atvd4() {
	slice := []int{10, 20, 30, 40, 50, 1, 2, 3, 4, 5}

	fmt.Print(slice[:3],"\n")
	fmt.Print(slice[4:],"\n")
	fmt.Print(slice[1:7],"\n")
	fmt.Print(slice[2:9],"\n")
}