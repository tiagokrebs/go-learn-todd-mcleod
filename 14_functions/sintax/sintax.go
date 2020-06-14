package main

import "fmt"

func main() {
	// func (r receiver) identifier(parameters) (return(s)) { ... }
	foo()
	bar("hello")
	w := woo("bye")
	fmt.Println(w)
	x, y := dar("Tiago", "Krebs")
	fmt.Println(x, y)
}

func foo() {
	fmt.Println("Foo")
}

func bar(s string) {
	fmt.Printf("Bar %s\n", s)
}

func woo(s string) string {
	return fmt.Sprint("Woo ", s)
}

func dar(s string, t string) (string, bool) {
	a := fmt.Sprint(s, t)
	b := false
	return a, b
}
