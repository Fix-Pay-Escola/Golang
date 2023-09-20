package main
import ("fmt")

func atvd9() {
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
	mapa["Maria_Livia"] = []string{
		"chorar",
		"gritar",
		"comer",
	}
	for k, v := range mapa{
		fmt.Println(k)
		for i,h := range v{
			fmt.Println(i," -- ",h)
		}
	}
}