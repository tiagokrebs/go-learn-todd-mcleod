package main

import "fmt"

type person struct {
	name string
	last string
}

type engineer struct {
	person
	sre bool
}

// receiver (e engineer) liga a funcao ao type engineer
// qualquer valor do type engineer tem acesso ao método speak
func (e engineer) speak() {
	fmt.Println("I am", e.name, e.last)
}

func main() {
	eng1 := engineer{
		person: person{
			"Tiago",
			"Krebs",
		},
		sre: true,
	}

	fmt.Println(eng1)
	eng1.speak()
}
