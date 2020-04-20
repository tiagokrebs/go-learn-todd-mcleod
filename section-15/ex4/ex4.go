package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// liga sleak a person
func (p person) speak() {
	fmt.Printf("Hey %s\n", p.first)
}

func main() {
	p1 := person{
		first: "Tiago",
		last:  "Krebs",
		age:   33,
	}

	p1.speak()
}
