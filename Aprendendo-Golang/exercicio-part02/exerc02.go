package exercicio_part02

import "fmt"

func main() {
	a := (20 == 20)
	b := (20 != 20)
	c := (100 <= 20)
	d := (100 < 20)
	e := (50 >= 20)
	f := (50 > 20)
	fmt.Printf("20 == 20: %v\n", a)
	fmt.Printf("20 != 20: %v\n", b)
	fmt.Printf("100 <= 20: %v\n", c)
	fmt.Printf("100 < 20: %v\n", d)
	fmt.Printf("50 >= 20: %v\n", e)
	fmt.Printf("50 > 20: %v\n", f)
}
