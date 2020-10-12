package main

import "fmt"

func main() {
	// buffered bufer allow up to 2 values befor unlock so this code wil work as expected
	c := make(chan int, 2)

	c <- 42
	c <- 43

	// note this action bellow is like a "pull of" from a queue, so first we have 42 and after 43
	fmt.Println(<-c)
	fmt.Println(<-c)
}
