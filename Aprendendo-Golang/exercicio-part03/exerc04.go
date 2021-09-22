package main

func main() {
	anoNasci := 2000
	anoAtual := 2021
	for {
		if anoNasci <= anoAtual {
			println(anoNasci)
			anoNasci++
		} else {
			break
		}
	}
}
