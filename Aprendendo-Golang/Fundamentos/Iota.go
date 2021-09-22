package main

import "fmt"

const (
	a = iota * 10 //pode declarar só um e tirar a declaração dos outros
	_             //= iota
	c             //= iota
	d             //= iota
	_             //= iota
	f             //= iota
)

func main() {
	fmt.Println(a, c, d, f)
}
