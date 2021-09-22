package main

import "fmt"

//const x = 10 //diferente do var, o tipo da variável só é atribuído quando for usado
//var y = 20

//declarar várias constantes
const (
	X = 10
	Y = 20
	z = 30
)

//var c int
//var d float64

func main() {
	//d = x
	fmt.Println(X, Y, z)
}
