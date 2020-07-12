package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex

	// corrige race condition do exercicio anterior com o uso de Mutex
	for i := 0; i < gs; i++ {
		go func() {
			// faz lock entre as goroutines
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			// faz unlock para a psroxima goroutine (que vai realizar outro lock)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
