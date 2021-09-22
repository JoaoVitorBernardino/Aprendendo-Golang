package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("decimal: %d\thexadeximal: %#x\tunicode: %v\n", i, i, string(i))

	}
}
