package main

import (
	"fmt"
)

func atvd7() {
	ss := [][]string{
		{
			"miu",
			"milton",
			"encher o saco",
		},
		[]string{
			"mimi",
			"martha",
			"pedir comida",
		},
		[]string{
			"meus alunos queridos",
			"que estudam bastante",
			"fazer os exercícios ninja",
		},
	}

	for _, v := range ss {
		fmt.Println(v)
	}

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}

	}

}
