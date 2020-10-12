package main

import "fmt"

func main() {
	c := make(chan int)

	// detach send to channel
	go foo(c)

	// block until channel receive something
	bar(c)

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
