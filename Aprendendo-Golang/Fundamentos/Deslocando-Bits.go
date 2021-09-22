package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB                    //= 1 << (iota * 10) // 1 << (2 * 10)
	GB                    //= 1 << (iota * 10) // 1 << (3 * 10)
	TB                    //= 1 << (iota * 10) // 1 << (4 * 10)
	//só precisava colocar um iota que ele conseguiria entender a sequência
)

func main() {
	fmt.Println("binary\t\t\t\t\t\tdecimal")
	fmt.Printf("%b\t\t\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)

	/*
		x := 24
		y := x << 2 // para usar o deslocamento, usa-se o << para deslocar para esquerda e >> para direita
		z := x >> 2

		fmt.Printf("%b\n", x)
		fmt.Printf("%b\n", y)
		fmt.Printf("%b\n", z)
	*/
}
