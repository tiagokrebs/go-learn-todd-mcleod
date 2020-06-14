package main

import "fmt"

func main() {
	s := []int{4, 5, 7, 32, 1, 6, 7, 5}
	for _, v := range s {
		fmt.Printf("%v\t", v)
		fmt.Println()
	}
}
