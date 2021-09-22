package main

import "fmt"

func main() {
	//s := "Hello, playground"
	//sb := []byte(s) //slice of byte
	//para ficar da formatação que quiser
	/*s := `Hello,
	playground
					asdjg`*/
	//fmt.Printf("%v\n%T", sb, sb) //%v dá o valor da variável e %T dá o tipo da variável
	s := "ASCII é\u00F8â \u9999"
	for _, v := range s { //usando range ele vai dar caracter por caracter
		fmt.Printf("%b - %T - %#U - %#x\n", v, v, v, v) //%v = valor numérico
		//%T = tipo
		//%#U = valor unicode
		//%#x = valor hexadecimal
		//%b = valor binário
	}
	fmt.Println()
	for i := 0; i < len(s); i++ { //vai dar por byte
		fmt.Printf("%b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
}
