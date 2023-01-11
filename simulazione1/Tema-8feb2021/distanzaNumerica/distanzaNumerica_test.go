/*
Scrivere un programma 'distanzaNumerica.go' che legge da linea di comando
una serie di numeri interi e verifica che la distanza numerica (differenza in valore assoluto)
fra due elementi consecutivi della sequenza sia sempre corrispondente
al numero di argomenti passati.
Se lo è, stampa la somma dei numeri; altrimenti stampa NON OK.
Anche nel caso la sequenza contenga nessun o un solo elemento, il programma stampa NON OK.

Esempi
------

$ go run distanzaNumerica.go 1 4 7
12

$ go run distanzaNumerica.go 10 7 7 7 8
NON OK

$ go run distanzaNumerica.go 10 5 0 -5 0
10

*/

package main

import (
	"testing"
	"fmt"
)

func TestIntestazione(t *testing.T) {
	fmt.Print(
		"\n",
		"****************************************\n",
		"*****                              *****\n",
		"*****    TEST DISTANZA NUMERICA    *****\n",
		"*****                              *****\n",
		"****************************************\n\n")
}

func TestOK(t *testing.T) {
	lanciaGenerica(t, "./distanzaNumerica", "", "12\n", "1", "4", "7")
	lanciaGenerica(t, "./distanzaNumerica", "", "10\n", "10", "5", "0", "-5", "0")
	lanciaGenerica(t, "./distanzaNumerica", "", "40\n", "10", "15", "10", "5", "0")
}

func TestNonOK(t *testing.T) {
	lanciaGenerica(t, "./distanzaNumerica", "", "NON OK\n", "10", "7", "12", "17", "12")
	lanciaGenerica(t, "./distanzaNumerica", "", "NON OK\n", "-10", "-15", "-20", "-25", "-40")
	lanciaGenerica(t, "./distanzaNumerica", "", "NON OK\n", "2", "-2", "3", "-1")
}

func Test1Arg(t *testing.T) {
	lanciaGenerica(t, "./distanzaNumerica", "", "NON OK\n", "10")
}

func Test0Arg(t *testing.T) {
	lanciaGenerica(t, "./distanzaNumerica", "", "NON OK\n")
}
