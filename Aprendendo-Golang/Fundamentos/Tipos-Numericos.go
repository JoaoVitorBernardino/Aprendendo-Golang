package main

import (
	"fmt"
	"runtime"
)

var x1 = 10
var y = 10.0

func main() {
	a := "e"
	b := "é"
	c := "\u9999"
	fmt.Printf("%v, %v, %v\n", a, b, c)
	d := []byte(a)
	e := []byte(b)
	f := []byte(c)
	fmt.Printf("%v, %v, %v\n", d, e, f)

	//x := 10
	//y := 10.0
	//x = 20.5 vai dar erro por ser tipos diferentes, precisa de conversão
	fmt.Printf("%v, %T\n", x1, x1)
	fmt.Printf("%v, %T\n", y, y)

	fmt.Println(runtime.GOOS)   //ver os SO
	fmt.Println(runtime.GOARCH) //ver a arquitetura
}
