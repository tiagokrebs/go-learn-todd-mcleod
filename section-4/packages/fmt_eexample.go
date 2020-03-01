package main

import (
	"fmt"
)

func main() {
	const name, age = "Kim", 22
	n, err := fmt.Println(name, "is", age, "years old.")
	fmt.Println(n)
	fmt.Println(err)

	// It is conventional not to worry about any
	// error returned by Println.

}
