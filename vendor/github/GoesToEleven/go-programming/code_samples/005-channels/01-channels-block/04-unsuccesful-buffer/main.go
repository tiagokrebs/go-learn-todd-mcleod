package main

import "fmt"

func main() {
	// buffered channel allow 1 value before unlock
	c := make(chan int, 1)

	c <- 42
	// this second value are not allowed on the channel so the block will occour and the code will stop with deadlock
	c <- 43

	fmt.Println(<-c)
}
