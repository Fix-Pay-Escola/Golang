package main
import ("fmt")

func atvd5() {
	slice := []int {42,43,44,45,46,47,48,49,50,51}
	slice = append(slice, 52,53,54,55)

	fmt.Print(slice,"\n")

	y := []int {56,57,58,59,60}

	slice = append(slice,y...)

	fmt.Print(slice)
}