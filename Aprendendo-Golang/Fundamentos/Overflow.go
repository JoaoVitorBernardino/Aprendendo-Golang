package main

import "fmt"

func main() {
	var i uint16
	i = 65535 //se colocar 65536 vai dar overflow, pois passou do número limite
	fmt.Println(i)
	i++ //o valor quando chega no limite ele zera, voltando para o valor de inicio ou menor valor
	fmt.Println(i)
	i++
	fmt.Println(i)
}
