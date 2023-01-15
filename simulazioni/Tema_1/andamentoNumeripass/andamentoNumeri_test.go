/*

andamentoNumeri
===============

Scrivere un programma in Go (il file deve chiamarsi 'andamentoNumeri.go') che traccia l'andamento
di una serie di numeri interi letti da standard input, uno per riga.
In particolare il programma legge una sequenza di lunghezza arbitraria, ma maggiore di 1,
di valori interi, terminata da EOF.

A partire dal secondo valore, si vuole indicare con '+', '=' o '-' se l'andamento è stato positivo, costante o negativo rispettivamente, cioè se il nuovo valore letto è maggiore, uguale o minore del valore precedente.

Il programma deve produrre in output una sequenza di '+', '=', '-' che corrisponde all'andamento della serie di numeri letti, a partire dal secondo valore, nel senso descritto qui sopra.

Si può assumere che il programma legga almeno due valori, non sono richiesti controlli in tal senso.

Nota: In fase di esecuzione, per terminare l'input occorre digitare un invio seguito da Ctrl-D.



Esempi
======

1) data in input la sequenza:
34
30
25
34
37
37
37
37

il programma genera la sequenza:
--++===



2) data in input la sequenza:
-938719
0
-73627
10
122
1347
1347
11

il programma genera la sequenza:
+-+++=-

*/

package main

import (
	//"strings"
	//"log"
	"testing"
)

var progname = "./andamentoNumeri"

func TestInput(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(t, progname, "uno.input", "uno.expected")
	LanciaGenericaConFileInOutAtteso(t, progname, "due.input", "due.expected")
	LanciaGenericaConFileInOutAtteso(t, progname, "tre.input", "tre.expected")
	LanciaGenericaConFileInOutAtteso(t, progname, "quattro.input", "quattro.expected")
	LanciaGenericaConFileInOutAtteso(t, progname, "cinque.input", "cinque.expected")
}
