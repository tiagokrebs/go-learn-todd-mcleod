package main

import "fmt"

func main() {
	m := map[string][]string{
		"Tiago": []string{"Pizza", "Sorvete", "Bicicleta"},
		"Ana":   []string{"Lasanha", "Chocolate", "Netflix"},
	}
	fmt.Println(m)

	for i, v := range m {
		fmt.Printf("Key: %v\n", i)
		for j, k := range v {
			fmt.Printf("\tValue in position %v is %s \n", j, k)
		}
	}

	m["Sem graca"] = []string{"Brocolis", "Bringela", "Yoga"}

	for i, v := range m {
		fmt.Printf("Key: %v\n", i)
		for j, k := range v {
			fmt.Printf("\tValue in position %v is %s \n", j, k)
		}
	}

}
