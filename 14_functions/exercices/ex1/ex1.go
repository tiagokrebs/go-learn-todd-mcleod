package main

import "fmt"

func main() {
	// focus pn waht is important; not upon what is urgent

	fmt.Printf("foo(): %v. bar(): %s\n", foo(), bar())
}

func foo() int {
	return 33
}

func bar() string {
	return "Hey"
}
