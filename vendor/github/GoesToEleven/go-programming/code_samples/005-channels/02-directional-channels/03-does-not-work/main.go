package main

import "fmt"

func main() {
	// receive only buffered channel
	c := make(<-chan int, 2)

	// send to this channel wont work because it is a receive only type
	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
