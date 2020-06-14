package main

import "fmt"

var a int

// golang is all about types
// declaracao do type hotdog (int)
type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// a = b
	// cannot use b (type hotdog) as type int in assignment
	// golan e uma linguagem estatica, um tipo 'hotdog' nao pode ser atribuido a um tipo 'int'

	// conversao de variaveis (nao cast)
	// sabemos que hotdog por baixo e um int entao e possivel converter
	a = int(b)
	fmt.Printf("%T\n", a)
	fmt.Println(b)

}
