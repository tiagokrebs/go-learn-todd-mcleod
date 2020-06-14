package main

// https://www.ardanlabs.com/blog/2015/09/composition-with-go.html

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
	fmt.Println("I am engineer", e.name, e.last)
}

func (e person) speak() {
	fmt.Println("I am person", e.name, e.last)
}

// todo type que tem o metodo speak tambem e do typo human
// um valor pode ter mais de um valor
type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("bar", h.(person).name)
	case engineer:
		fmt.Printf("bar %s engineer\n", h.(engineer).name)
	}
	fmt.Println("bar end")
}

func main() {
	// eng1 e do typo engineer e human ao mesmo tempo
	eng1 := engineer{
		person: person{
			"Tiago",
			"Krebs",
		},
		sre: true,
	}
	fmt.Println(eng1)
	eng1.speak()
	bar(eng1)

	p1 := person{
		"Ana",
		"Giulia",
	}
	fmt.Println(p1)
	p1.speak()
	bar(p1)
}
