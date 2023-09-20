package main

import (
	"fmt"
)

func atvd8() {
	mapa := map[string][]string{
		"Santos_Lucas": []string{
			"comer", "dormir", "acordar cedo",
		},
		"Andrade_Charles": []string{
			"gritar", "beber", "gritar mais",
		},
		"Angelica_Maria": []string{
			"comer", "dormir", "Gritar",
		},
	}
	for k, v := range mapa {
		fmt.Println(k,"\nhobbys:")
		for i, h := range v {
			fmt.Println("\t", i+1, " -- ", h)
		}
	}
}
