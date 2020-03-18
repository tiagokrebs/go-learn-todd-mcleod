package main

import "fmt"

func main() {
	// don't use array, use slices instead
	var x [5]int
	fmt.Println(x)
	x[3] = 33
	fmt.Println(x)

	fmt.Println(len(x))
}
