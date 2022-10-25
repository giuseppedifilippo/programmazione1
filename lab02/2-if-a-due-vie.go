package main

import "fmt"

func main() {
	/*
	 scrivere un programma che preso un numero inter oda input
   ritorna se sia pari o dispari
	*/

	var n int
	fmt.Print("numero: ")
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println(n, "è pari")
	} else {
		fmt.Println(n, "è dispari")
	}
}
