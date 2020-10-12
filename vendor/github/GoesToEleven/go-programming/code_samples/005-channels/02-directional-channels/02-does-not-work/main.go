package main

import "fmt"

func main() {
	// send only buffered channel
	c := make(chan<- int, 2)

	c <- 42
	c <- 43

	// pull from this channel won't work because it is a send only type
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
