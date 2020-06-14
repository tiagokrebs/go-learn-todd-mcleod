package main

import "fmt"

type person struct {
	name    string
	address string
}

func main() {
	p1 := person{
		name:    "Tiago",
		address: "POA",
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

// a method set to person struct
func changeMe(p *person) {
	p.name = "Ana"
}
