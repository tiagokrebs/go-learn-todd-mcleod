package main

import "fmt"

func main() {
	name := "Tiago"
	switch {
	case name == "Ana":
		fmt.Println("Hi! my name is Ana")
	case name == "Tiago":
		fmt.Println("Hi! My name is Tiago")
	}
}
