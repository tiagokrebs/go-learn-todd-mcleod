package main

import "fmt"

type person struct {
	nome      string
	sobrenome string
}

type engineer struct {
	person
	skill bool
}

func main() {
	p1 := person{
		nome:      "Ana",
		sobrenome: "Giulia",
	}

	eg := engineer{
		person: person{
			nome:      "Tiago",
			sobrenome: "Krebs",
		},
		skill: true,
	}

	fmt.Println(p1)
	fmt.Println(p1.nome, p1.sobrenome)

	fmt.Println(eg)
	fmt.Println(eg.nome, eg.sobrenome, eg.skill)

	// Anonymous struct doesnt have a name or a type declared
	p2 := struct {
		nome      string
		sobrenome string
	}{
		nome:      "James",
		sobrenome: "Bond",
	}

	fmt.Println(p2)

}
