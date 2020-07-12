package main

import (
	"fmt"
	"runtime"
)
// modelo simples de estrutura concorrente utilizando goroutines
func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// a saída de foo() não é exibida porque main() encerra durante a execução da goroutine foo()
	// Wait no package sync é uma forma de garantir a execução de foo()
	// ver wait.go
	go foo()
	bar()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func foo() {
	for i:=0; i<10; i++ {
		fmt.Println("foo:", i)
	}
}

func bar() {
	for i:=0; i<10; i++ {
		fmt.Println("bar:", i)
	}
}