package exercicio_part02

import "fmt"

func main() {
	x := 10
	fmt.Println("decimal\tbinário\thexadecimal")
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)
	y := x << 1
	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}
