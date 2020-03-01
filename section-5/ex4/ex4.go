package main

import "fmt"

// declaracao de novo tipo mytype com type int subjacente
type mytype int

// declaracao de variavel com type mytype
var x mytype

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
