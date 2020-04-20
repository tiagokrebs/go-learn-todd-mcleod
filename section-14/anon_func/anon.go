package main

import "fmt"

func main() {

	// anonymous function
	func() {
		fmt.Println("Anonymous func here")
	}()

	// with parameters
	func(x int) {
		fmt.Println("Anonymous func here:", x)
	}(42)
}
