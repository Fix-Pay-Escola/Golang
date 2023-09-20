package main

import (
	"fmt"
)

func atvd1() {
	type pessoa struct {
		nome      string
		sobrenome string
		sorvete   []string
	}

	pessoa1 := pessoa{
		nome:      "Lucas",
		sobrenome: "Santos",
		sorvete: []string{
			"chocolate",
			"creme",
		},
	}
	pessoa2 := pessoa{
		nome:      "Charles",
		sobrenome: "Andrade",
		sorvete: []string{
			"brigadeiro",
			"doce de leite",
		},
	}
	for _, v := range pessoa1.sorvete {
		fmt.Println(v)
	}
	for _, v := range pessoa2.sorvete {
		fmt.Println( v)
	}

}
