package main

import "fmt"

func main() {
	/*
	 Scrivi un programma che inserita una variabile intera ritorna se
   il voto non sia valido ( minore di 0  o maggiore di 30 )
	*/

	var voto int

	fmt.Print("voto: ")
	fmt.Scan(&voto)

	if voto < 0 || voto > 30 {
		fmt.Println("voto non valido")
	}
}
