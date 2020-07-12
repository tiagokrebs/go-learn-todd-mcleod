package main

import (
	"fmt"
	"runtime"
	"sync"
)

// waitgroup para controle de fluxo de main()
var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// adiciona delta ao waitgrpup
	wg.Add(1)
	// dispara goroutine
	go foo()
	bar()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// mantem main() ex execução até que o waigroup receba .Done em foo()
	wg.Wait()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func foo() {
	for i:=0; i<10; i++ {
		fmt.Println("foo:", i)
	}
	// define done no waitgroup
	wg.Done()
}

func bar() {
	for i:=0; i<10; i++ {
		fmt.Println("bar:", i)
	}
}