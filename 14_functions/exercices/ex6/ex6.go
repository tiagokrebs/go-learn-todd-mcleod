package main

import (
	"fmt"
)

func main() {

	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	// funcoes anonimas podem apresentar casos de uso comum
	// com a chamada `go func() {...}`
	// dessa forma fazendo com que a funcao seja eexecutada em
	// paralelo com a rotina principal (multi CPU)

	fmt.Println("done")
}
