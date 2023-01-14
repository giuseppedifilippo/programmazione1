/*

Scrivere un programma 'pariDispari.go' che legge una sequenza di valori interi da linea di comando
e controlla che si alternino valori pari e valori dispari.
In questo caso il programma stampa il messaggio "sequenza valida", altrimenti "elemento in posizione i non valido",
dove i è la posizione del primo  elemento (da sinistra) che non rispetta la regola
di alternanza o che non è un valore numerico.

In caso di mancanza di valori, il programma deve stampare "nessun valore in ingresso".

La sequenza può iniziare sia con un valore pari sia con uno dispari.

Si ricorda che lo zero è un numero pari.

Esempi
------

Argomenti:
	3 8 1 12
Output:
	sequenza valida

Argomenti:
	4
Output:
	sequenza valida

Argomenti:
	1 2 3 5
Output:
	elemento in posizione 4 non valido

Argomenti:
	1 2 3eqeqw 5
Output:
	elemento in posizione 3 non valido

Argomenti:

Output:
	nessun valore in ingresso

*/

package main

import (
	"testing"
)

var progname = "./pariDispari"

func TestNoParm(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"nessun valore in ingresso\n")
}

func TestOneParm(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"sequenza valida\n",
		"2")
}

func TestWrong1(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"elemento in posizione 3 non valido\n",
		"3", "8888", "0", "11", "57")
}

func TestWrong2(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"elemento in posizione 4 non valido\n",
		"1", "2", "3", "5")
}

func TestWrong3(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"elemento in posizione 5 non valido\n",
		"112", "1", "2", "3", "5")
}

func TestWrong4(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"elemento in posizione 2 non valido\n",
		"5", "5", "10", "21", "40")
}

func TestWrong5(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"elemento in posizione 3 non valido\n",
		"3", "4", "6", "98", "101")
}

func TestWrong6(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"elemento in posizione 4 non valido\n",
		"40", "3", "6", "80", "67")
}

func TestOk1(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"sequenza valida\n",
		"4", "5", "12", "11", "0", "77")
}

func TestOk2(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"sequenza valida\n",
		"3", "4", "5", "12", "11", "0", "77")
}

func TestOkEsempio(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"sequenza valida\n",
		"3", "8", "1", "12")
}

func TestWrongType1(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"elemento in posizione 3 non valido\n",
		"3", "4728", "5555y555", "-12", "-11", "0", "77")
}

func TestWrongType2(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"elemento in posizione 1 non valido\n",
		"sadas3", "4728", "5555y555", "-12", "-11", "0", "77")
}

func TestWrongType3(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"elemento in posizione 2 non valido\n",
		"3", "ciao", "ciao")
}

func TestWrongType4(t *testing.T) {
	LanciaGenerica(
		t,
		progname,
		"NIENTE",
		"elemento in posizione 2 non valido\n",
		"3", "ciao", "3", "ciao", "3", "ciao")
}
