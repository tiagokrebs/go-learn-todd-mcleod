package main

import "fmt"

func main() {

	// funcao em Go e um type logo pode definida a uma variavel
	// chamamos isso de function expression
	f := func() {
		fmt.Println("Function expression")
	}

	f()

	g := func(x int) {
		fmt.Println("Function expression:", x)
	}

	g(33)

}
