/*
Scrivere un programma in Go (il file deve chiamarsi 'leastDigit.go') che, data una serie di numeri interi non negativi da standard input, terminata da -1, conta quanti di questi numeri (escluso il -1) hanno come cifra meno significativa (cifra delle unità) la cifra 0, quanti la cifra 1, quanti la cifra 2, ...., quanti la cifra 9.
Il programma ignora, senza terminare, eventuali dati (tranne -1 che è il terminatore dell'input) che non siano interi non negativi.
Il programma stampa infine quanto calcolato (vedi esempio sotto).

Il programma deve essere dotato di:

- una funzione
	checkInput(s string) int
che restituisce -2 se la stringa s non rappresenta un intero >= -1 (se quindi s rappresenta un int < -1 o un non int), e altrimenti restituisce l'intero corrispondente a s.

- una funzione
	leastDigit(n int) int
che restituisce la cifra meno significativa (quella più a destra, la cifra delle unità) del numero n.


Esempio
-------
dato il seguente input

123
456 789 ciao uno due
-35
66 88 99
-2
-1

il programma stampa
0 - 0
1 - 0
2 - 0
3 - 1
4 - 0
5 - 0
6 - 2
7 - 0
8 - 1
9 - 2

*/

package main

import (
	//"fmt"
	//"log"
	//"os/exec"

	//"strings"
	"fmt"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

var prog = "./leastDigit"

func TestComeDaEsempio(t *testing.T) {
	lanciaGenerica(
		t,
		prog,
		"123\n456 789 ciao uno due\n-35\n66 88 99\n-2\n-1\n",
		"0 - 0\n1 - 0\n2 - 0\n3 - 1\n4 - 0\n5 - 0\n6 - 2\n7 - 0\n8 - 1\n9 - 2\n")
}

func TestFuncLeastDigit1(t *testing.T) {
	if leastDigit(2183712) != 2 {
		fmt.Println(">>> funzione leastDigit FAIL!")
		t.Fail()
	}
}

func TestFuncCheckInputNonValue(t *testing.T) {
	if checkInput("iqwueyqwiu") != -2 || checkInput("-767") != -2 {
		fmt.Println(">>> funzione checkInput (non value) FAIL!")
		t.Fail()
	}
}

func TestFuncCheckInputNormValue(t *testing.T) {
	if checkInput("8768") != 8768 {
		fmt.Println(">>> funzione checkInput (norm value) FAIL!")
		t.Fail()
	}
}
