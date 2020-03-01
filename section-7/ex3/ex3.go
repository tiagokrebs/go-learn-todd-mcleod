package main

import "fmt"

const (
	x        = 33
	y string = "I need more coffe"
)

func main() {
	fmt.Println(x, "\t", y)
	fmt.Printf("%T\t%T\n", x, y)
}
