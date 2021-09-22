package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 4512
	fmt.Println(slice[3])
	//slice[20] = 1 //ocorre erro, pois não existe o índice 20 nesse slice
	//fmt.Println(slice[20])
}
