package main

import "fmt"

func main() {
	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}

	x := 1
	for {
		x++
		// se x > 0 para loop
		if x > 9 {
			break
		}
		// se impar seguee para proxima iteracao
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}
