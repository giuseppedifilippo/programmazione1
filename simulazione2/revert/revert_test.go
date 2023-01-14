/*

Revert
------

Realizzare un programma Go (nome file 'revert.go') che legge da standard input **stringhe** costituite da soli caratteri **ASCII standard (a 7 bit)** (separate da un numero arbitrario di white space) e stampa solo quelle di lunghezza pari, al contrario (il primo carattere per ultimo e l'ultimo per primo), una per riga.
Il programma termina quando legge la stringa "stop", che non va trattata.

*/

package main

import "testing"

var prog = "./revert"

func TestEsempio(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"esempio.input",
		"esempio.output")
}
