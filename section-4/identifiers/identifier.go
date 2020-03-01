package main

import "fmt"

// declaracao fora de void {} usa var
// o escopo de y e o package
var y = 43

// declara variavez "z" do tipo inteiro inicializada com valor 0 (int default)
var z int

func main() {
	// declarando e atribuindo valor (short declaration)
	// o escopo de x é tudo a partir da sua declaracao
	x := 33
	fmt.Println(x)

	// atrinbuindo valor
	x = 55
	fmt.Println(x)

	fmt.Println(y)

	fmt.Println(z)
}
