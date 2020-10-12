package main

import "fmt"

func main() {
	// buffered channel will alow 1 value on it befora unlock so this code will work as expected
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
