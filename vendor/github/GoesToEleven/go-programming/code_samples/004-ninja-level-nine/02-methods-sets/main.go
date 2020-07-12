package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	saySomething(&p1)

	// does not work
	// o type da function importa no uso de interfaces
	// saySomething espera um ponteiro, logo quando usamos um valor o compilador
	// nao tem o metodo no stack para ser utilizado
	// saySomething(p1)

	p1.speak()
}
