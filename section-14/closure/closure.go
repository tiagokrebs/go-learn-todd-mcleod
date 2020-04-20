package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	// x faz parte do scopo de incrementor assim como a funcao
	// anonima abaixo
	// dessa forma a funcao anonima tem acesso a x
	return func() int {
		x++
		return x
	}
}
