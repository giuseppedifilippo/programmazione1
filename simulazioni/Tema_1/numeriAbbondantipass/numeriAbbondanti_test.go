/*
NUMERI ABBONDANTI
-----------------

Scrivere un programma (il file deve chiamarsi 'numeriAbbondanti.go') che legge un numero intero positivo n da standard input e stampa i primi n "numeri abbondanti", uno per riga.

Un numero abbondante è un numero naturale minore della somma dei suoi divisori interi (1 compreso, numero stesso ovviamente escluso).

Per esempio, 12 è un numero abbondante poiché inferiore alla somma dei suoi divisori: 1+2+3+4+6=16
mentre invece 15 non lo è in quanto 1+3+5 = 9

Il programma deve essere dotato di una funzione
	Abbondante(n int) bool
che, dato un intero positivo, restituisce true se n è un numero abbondante, false altrimenti.
Se il numero passato come parametro è minore o uguale a 0, la funzione restituisce false.

Esempio di esecuzione
---------------------

./numeriAbbondanti
3
12
18
20

NOTA BENE: il 3 alla prima riga è l'input, i numeri successivi sono l'output

*/

package main

import (
	"fmt"
	"testing"
)

var prog = "./numeriAbbondanti"

func TestMain(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"6",
		"12\n18\n20\n24\n30\n36\n")
}

func TestMain2(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"20",
		"12\n18\n20\n24\n30\n36\n40\n42\n48\n54\n56\n60\n66\n70\n72\n78\n80\n84\n88\n90\n")
}

func TestAbbondante(t *testing.T) {
	if !Abbondante(12) {
		fmt.Println("12 è abbondante!")
		t.FailNow()
	}
	if !Abbondante(48) {
		fmt.Println("48 è abbondante!")
		t.FailNow()
	}
	if Abbondante(11) {
		fmt.Println("11 non è abbondante!")
		t.FailNow()
	}
	if Abbondante(65) {
		fmt.Println("65 non è abbondante!")
		t.FailNow()
	}

	if Abbondante(-50) {
		fmt.Println("-50 è negativo!")
		t.FailNow()
	}
	//fmt.Println(Abbondante(11312312))
	if !Abbondante(11312312) {
		fmt.Println("11312312 non è abbondante!")
		t.FailNow()
	}
}
