package main

import("fmt"
)
 
func bytes() {
	//utf codes
	a:= "e"
	b:= "é"
	c:= "Ħ"
	fmt.Printf("%v,%v,%v\n",a, b, c)
	
	//vamos converter os valores de a,b e c para saber o numero de bytes desses caracteres
	d:= [] byte(a)
	e:= [] byte(b) 
	f:= [] byte(c)
	fmt.Printf("%v,%v,%v",d, e, f)

}


 // o resultado será : [101] 9 primerio caractere usa 2 bytes,[195 169],[196 166]


 