package main

import (
	"fmt"
)

func exemplo1() {
	x := 10
	y := "Olá, Bom dia! "

	//forma de executar com o  println:
	//fmt.Println(x,y)

	//Printf:
	fmt.Printf("x: %v,%T\n", x, x) /* Este número é Int (numero inteiro) */
	fmt.Printf("y: %v,%T\n", y, y) /* Este texto texto é string */

	//mudando o valor do x e adicionando outra variavel
	x,z := 20, 30
	fmt.Printf("x: %v,%T\n", x, x) /* vale lembrar que o "x" no final so pega na ordem que as coisas , no caso o primeiro x pega o "%v"", e o segundo x pega o "%T""Ï*/
	fmt.Printf("z: %v,%T\n", z, z)
} 

//Outro codigo sobre true ou false
func calculo(){
	l:= 10==7
	fmt.Println(l)
}

//Exemplo de criar funçao para chamar na tela

func principal() { /*main*/
	a:= 10
	qualquercoisa(a)
}

func qualquercoisa (x int) {
	fmt.Println(x)
}
//leleo
