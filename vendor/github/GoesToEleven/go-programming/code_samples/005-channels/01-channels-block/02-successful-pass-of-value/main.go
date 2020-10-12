package main

import "fmt"

func main() {
	c := make(chan int)

	// concurrent goroutine will read channel as eexpected
	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
