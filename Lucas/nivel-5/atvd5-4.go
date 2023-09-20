package main

import (
	"fmt"
)

func atvd4() {

	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "agua",
		sabor:   "nada",
		ondetem: []string{"parangaba", "Shopping", "na minha casa"},
		vaibemcom: map[string]string{
			"Café da manhã": "pão",
			"Almoço":        "lasanha"},
	}

	fmt.Println(x)
}
