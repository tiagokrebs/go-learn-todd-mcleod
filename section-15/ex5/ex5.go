package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

// liga funcao area em struct circle
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//liga funcao area em struct square
func (s square) area() float64 {
	return s.length * s.length
}

// cria iterface e proporciona uso de polimorfismo
type shape interface {
	// cria interface para todas as funcoes area() disponiveis
	area() float64
}

func info(s) {
	fmt.Println(s.area())
}

func main() {
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
	}

	info(circ)
	info(squa)
}
