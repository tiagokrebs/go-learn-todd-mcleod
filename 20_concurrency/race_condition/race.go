package main

import (
	"fmt"
	"runtime"
	"sync"
)

// $ run race.go
// $ go run -race race.go -> identifica race condition e retorna
func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// cada goroutine tem v na sua pilha e compartilha counter com outras goroutines
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	// nao pare o programa antes que todas as goroutines estejam done
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
