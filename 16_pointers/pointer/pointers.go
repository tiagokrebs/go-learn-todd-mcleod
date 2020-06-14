package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	// & representa o endereço na memoria
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	// * representa o valor no endereço
	fmt.Println(*b)
	fmt.Printf("%T\n", b)

	// nice!
	fmt.Println(*&a)

	// define 43 no endereço de a
	*b = 43
	fmt.Println(a)
}
