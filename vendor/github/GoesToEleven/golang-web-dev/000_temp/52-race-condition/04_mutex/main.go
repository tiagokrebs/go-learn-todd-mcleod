package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	// Mutex permite o uso de Lock e Unlock durante a execucao das goroutines
	// essa e uma das formas de controle de acesso a memória compartilhada entre elas,
	// nesse caso a variavel counter
	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			// Lock bloqueia acesso a counter para todas as outras goroutines em execução
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			// Unlock libera acesso para a próxima goroutine
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
