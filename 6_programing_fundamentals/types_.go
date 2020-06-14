package main

import "fmt"

var x bool
var y int
var z float64
var v string

const (
	// sem type compilador tem flexibilidade para decisao do type
	t = 42
	// com type
	u string = "Cofee"
)

const (
	// iota autoincrementa int type
	p = iota
	q = iota
	r
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	boolean()
	numeric()
	string_()
	constant()
	iota_()
}

// obtem type de variavel atraves de string formatting
func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// bool example
func boolean() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	a := 33
	b := 35
	fmt.Println(a == b)
	fmt.Println(typeof(a) == typeof(b))
}

// numeric example
func numeric() {
	y = 30
	z = 55.68
	fmt.Println(y, z)
	fmt.Printf("%T %T\n", y, z)
	a := 32
	b := 33.865
	fmt.Println(a, b)
	fmt.Printf("%T %T\n", a, b)
}

// string example
func string_() {
	fmt.Println(v)
	s := "toast and coffe"
	fmt.Println(s)

	// string e uma sequencia de bytees
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println()
	for i, v := range s {
		fmt.Printf("Na posicao %d temos o hexa %#x\n", i, v)
	}
}

// constant example
func constant() {
	fmt.Println(t, u)
	fmt.Printf("%T %T\n", t, u)

	// t = 43
	// cannot assign to t
	// constante tem valor imutavel
}

// iota example
func iota_() {
	fmt.Println(p, q, r)
	fmt.Printf("%T\t%T\t%T\n", p, q, r)

	fmt.Println("binary\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
}
