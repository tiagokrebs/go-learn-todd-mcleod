package main

import "fmt"

const (
	x = 33
	y = 27
)

func main() {
	a := (x == y)
	fmt.Println(a)
	b := (x <= y)
	fmt.Println(b)
	c := (x >= y)
	fmt.Println(c)
	d := (x != y)
	fmt.Println(d)
	e := (x < y)
	fmt.Println(e)
	f := (x > y)
	fmt.Println(f)
}
