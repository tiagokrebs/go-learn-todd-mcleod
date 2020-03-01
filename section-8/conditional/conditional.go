package main

import "fmt"

func main() {
	if true {
		fmt.Println("Hey i'm true")
	}

	if x := 2; x == 2 {
		fmt.Println("two statements")
	}

	x := 33
	if x == 34 {
		fmt.Println("Value is 34")
	} else if x == 33 {
		fmt.Println("Value is 33")
	} else {
		fmt.Println("Value is not 33 or 34")
	}

	switch {
	case false:
		fmt.Println("this not print")
	case (2 == 2):
		fmt.Println("this should print")
	case (4 == 4):
		fmt.Println("this print only if you use 'fallthrough'")
		fallthrough
	case (3 == 4):
		fmt.Println("do nothing")
	default:
		fmt.Println("you know how this work")
	}

	y := 33
	switch y {
	case 9:
		fmt.Println("Nop")
	case 33:
		fmt.Println("Look at that!")
	case 5, 88, 34:
		fmt.Println("Hello again, no")
	}
}
