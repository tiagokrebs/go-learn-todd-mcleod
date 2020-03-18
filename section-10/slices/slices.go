package main

import "fmt"

var y = []int{22, 35, 6, 9, 77}

func main() {
	// composite literal
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)

	//slice permite agrupar valores do mesmo type

	for_range()
	slice_slice()
	slice_append()
	del_from_slice()
	slice_array_make()
	multid_slice()
}

func for_range() {
	for i, v := range y {
		fmt.Println(i, v)
	}
}

func slice_slice() {
	fmt.Println(y[1:])

	// em slices o inicio e incluso enquanto o final nao
	fmt.Println(y[1:3])
}

func slice_append() {
	y = append(y, 1, 2, 3)
	fmt.Println(y)

	z := []int{9, 8, 7}
	// z... e um variatic parameter
	// todos os valors int de z estao em z...
	y = append(y, z...)
	fmt.Println(y)
}

func del_from_slice() {
	// removing positions from slice
	x := append(y[:2], y[5:]...)
	fmt.Println(x)
}

func slice_array_make() {
	// slice tem tamanho dinamico/variavel porem e construido em cima de um array
	// que possui um tamanho fixo, quando o tamanho do slice excede o array o compilador
	// precisa instanciar um novo array, issuo usa poder computacional

	// se o tamanho maximo do array e previstro podemos usar a funcao builtin 'make'
	// informando o tamanho do splice e a capacidade maxima do array

	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[3] = 9
	fmt.Println(x)

	// isso nao funciona porque a posicao 42 nao existe no slice
	// x[42] = 33

	// o append extende o tamaho do slice
	x = append(x, 999)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// quando a capacidade do array e ultrapassada um novo array e criado pelo compilador
	// com a mesma capacidade informada no comando make
	// o append abaixo excede o tamanho de 12 posicoes do array entao um novo array de 12 e
	// criado fazendo com que a capacidade total do nosso array aumente para 24
	x = append(x, 998, 997, 996)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func multid_slice() {
	x := []string{"Tiago", "Krebs", "Pizza"}
	z := []string{"Ana", "Giulia", "Lasanha"}
	xz := [][]string{x, z}
	fmt.Println(xz)
}
