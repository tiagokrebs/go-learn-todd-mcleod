package main

import "fmt"

func main() {
	// Channels block goroutine until channel is not ready to rbe readed so this code won't work
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}
