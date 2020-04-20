package main

import "fmt"

func main() {
	// nothing magic or strange in here
	s1 := foo()
	fmt.Println(s1)

	// x recebe uma anonymous function
	x := bar()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
	fmt.Println(bar()())

}

func foo() string {
	s := "hey"
	return s
}

// bar retorna uma anonymous function
func bar() func() int {
	// anon function retorna int 33
	return func() int {
		return 33
	}
}
