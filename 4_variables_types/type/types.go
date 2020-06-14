package main

import "fmt"

var x = 33
var y = "Cerveja e fandangos"

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// y = 40
	// cannot use 40 (type untyped int) as type string in assignment
	// y foi declarada como tipo string, o tipo em goland e imutavel
	// golang e uma linguagem de programacao estatica
}
