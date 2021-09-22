package main

import "fmt"

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Print("Hora: ", horas, "\n")
		for i := 0; i < 60; i++ {
			fmt.Print(" ", i)
		}
		fmt.Println("\n")
	}
}
