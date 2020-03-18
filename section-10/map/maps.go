package main

import "fmt"

func main() {
	// composite literal
	m := map[string]int{
		"Tiago": 33,
		"Ana":   27,
	}
	fmt.Println(m)
	fmt.Println(m["Tiago"])

	// a key Jarbas nao existe entao o zero value do type do map e retornado
	// o map tem tipo int (zero value = 0)
	fmt.Println(m["Jarbas"])

	// teste ,ok para a key Jarbas
	v, ok := m["Jarbas"]
	fmt.Println(v)
	fmt.Println(ok)
	// com if
	if v, ok := m["Ana"]; ok {
		fmt.Println("Valor em if para key Ana", v)
	}

	// add to map
	m["Rex"] = 1
	fmt.Println(m)

	// for range on map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// delete from map
	delete(m, "Tiago")
	fmt.Println(m)
	// deletar key nao existente nao gera excessao
	delete(m, "Jarbas")
	fmt.Println(m)
}
