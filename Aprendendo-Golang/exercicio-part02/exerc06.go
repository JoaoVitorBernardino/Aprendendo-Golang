package exercicio_part02

import "fmt"

const (
	_ = 2021 + iota
	x1
	x2
	x3
	x4
)

func main() {
	fmt.Printf("%v\n%v\n%v\n%v\n", x1, x2, x3, x4)
}
