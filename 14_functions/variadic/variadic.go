package main

import "fmt"

func main() {
	foo(2, 3, 4, 5, 7, 8, 9)
}

// ...int -> um número ilimitado de parametros do type int
func foo(x ...int) {
	fmt.Printf("%T\n", x) // []int
	fmt.Println(x)
}
