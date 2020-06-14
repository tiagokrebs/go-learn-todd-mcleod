package main

import "fmt"

func main() {
	// atrasa a funcao foo() até o final de main()
	defer foo()
	defer bar()
	lag()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func lag() {
	fmt.Println("lag")
}
