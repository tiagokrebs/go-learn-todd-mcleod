package main

import "fmt"

func main() {
	// passando []int
	fmt.Printf("Foo sum: %v\n", foo(1, 2, 3))

	// passando slice de int
	p1 := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("Foo sum: %v\n", foo(p1...))

	// too many arguments in call to bar
	// fmt.Printf("Bar sum: %v\n", bar(1, 2, 3))

	// passando slice de int
	p2 := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("Bar sum: %v\n", bar(p2))
}

// foo recebe N elementos type int
func foo(x ...int) int {
	sum := 0
	for i := range x {
		sum += i
	}
	return sum
}

// recebe int[]
func bar(x []int) int {
	sum := 0
	for i := range x {
		sum += i
	}
	return sum
}
