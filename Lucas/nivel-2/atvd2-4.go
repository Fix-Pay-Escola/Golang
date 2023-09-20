package main
import ("fmt")
func atvd4(x int){
	fmt.Printf("%d,%b,%#x\n",x,x,x)
	y := x << 1
	fmt.Printf("%d,%b,%#x",y,y,y)
}