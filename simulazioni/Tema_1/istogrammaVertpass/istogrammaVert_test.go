/*
ISTOGRAMMA
----------

Scrivere un programma (il file deve chiamarsi 'istogrammaVert.go') che legge da linea di comando una serie di numeri interi non negativi e ne disegna il relativo istogramma. L'istogramma sarà formato da colonne di asterischi che rappresentano i valori letti.
Non sono richiesti controlli sui dati.

Esempio:

$ ./istogrammaVert 5 0 2 9 4
   *
   *
   *
   *
*  *
*  **
*  **
* ***
* ***

$ ./istogrammaVert 2 3 4 5
   *
  **
 ***
****
****


*/
package main

import (
	"testing"
)

var prog = "./istogrammaVert"

func TestCrescente(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"crescente.output",
		"1", "2", "3", "4")
}

func TestDecrescente(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"decrescente.output",
		"5", "4", "3", "2", "1")
}

func TestEsempio(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"esempio.output",
		"5", "0", "2", "9", "4")
}

func TestTanti(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"tanti.output",
		"5", "0", "2", "9", "4", "5", "0", "12", "9", "14", "5", "0", "29", "2", "4")
}
