package main
import ("fmt")

func atvd2() {
	slice := []int{1,2,3,4,5,6,7,8,9}
	
	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
	}
}