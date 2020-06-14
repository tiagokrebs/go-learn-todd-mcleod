package main

import "fmt"

func main() {
	foo()
	defer last()
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func foo() {
	fmt.Println("Foo")
}

func last() {
	fmt.Println("Hey i'm the last one")
}
