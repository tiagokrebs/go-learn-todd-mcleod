package main

import "fmt"

func main() {
	if 2 == 3 {
		fmt.Println("Not print")
	} else if 2 == 4 {
		fmt.Println("Also false")
	} else {
		fmt.Println("All is false so i'm in the final else")
	}
}
