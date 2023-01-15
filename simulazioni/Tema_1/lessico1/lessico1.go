package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	opts := []byte{}
	fmt.Println("+ (legge e memorizza)")
	fmt.Println("? (numeri di riga in cui è comparsa la parola data)")
	fmt.Println("m (stampa le lunghezze min e max)")
	fmt.Println("p (stampa le parole memorizzate)")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		opts = append(opts, scanner.Text()[0])
	}

	for _, v := range opts {
		switch(v) {
		case '+':
			fmt.Println("alimento il dizionario")
			break
		case '?':
			fmt.Println("consulto il dizionario")
			break
		case 'm':
			fmt.Println("lunghezza min e max")
			break
		case 'p':
			fmt.Println("stampo il dizionario ordinato")
			break
		default:
			fmt.Println("opzione non valida")
			break
		}
	}
}