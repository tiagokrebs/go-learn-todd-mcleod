package main

import "fmt"

func main() {
	x := 33
	fmt.Printf("%d\t%b\t%x\n", x, x, x)
	// hexa com zero a esquerda
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
}
