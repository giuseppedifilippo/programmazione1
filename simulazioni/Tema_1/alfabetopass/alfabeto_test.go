/*

Alfabeto
--------

Scrivere un programma (il cui file si deve chiamare 'alfabeto.go') che fa la compitazione, cioè converte le parole in sequenze di stringhe corrispondenti ai nomi delle lettere di cui sono composte (a -> "a", b -> "bi", c -> "ci", d -> "di", e -> "e", f -> "effe", ..., z -> "zeta")

Il programma legge una parola (una stringa composta di sole lettere) da linea di comando, e per ogni lettera incontrata per la prima volta chiede all'utente di immettere il nome corrispondente (alimentando un "dizionario").
Alla fine stampa la sequenza delle stringhe che compongono la compitazione della parola passata per argomento, separate da " - ".

L'unico controllo richiesto sull'input è che ci sia almeno un argomento (parola).
Si assuma che le stringhe passate come argomenti contangano solo lettere e non altri caratteri.
Gli eventuali argomenti supplementari devono essere ignorati.
Il programma non distingue tra lettere maiuscole e minuscole.


Ad esempio, per la parola "Zappatrice", il programma stamperà:
zeta - a - pi - pi - a - ti - erre - i - ci - e

Esempio di esecuzione
---------------------

$ go run alfabeto.go sasso
s? esse
a? a
o? o
esse - a - esse - esse - o


*/

package main

import (
	//"fmt"
	//"log"
	//"os/exec"

	//"strings"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

var prog = "./alfabeto"

func TestComeDaEsempio(t *testing.T) {
	lanciaGenerica(
		t,
		prog,
		"esse\na\no\n",
		"s? a? o? esse - a - esse - esse - o\n",
		"sasso")
}

func TestConTantiArg(t *testing.T) {
	lanciaGenerica(
		t,
		prog,
		"pi\ni\no\n",
		"p? i? o? pi - i - pi - pi - o\n",
		"pippo", "pluto", "paperino", "qui", "quo", "qua")
}

func TestPippo2(t *testing.T) {
	lanciaGenerica(
		t,
		prog,
		"pi\ni\no\n",
		"p? i? o? pi - i - pi - pi - o - pi - i - pi - pi - o\n",
		"pippoPIPPO")
}

func TestZuzzurellone(t *testing.T) {
	lanciaGenerica(
		t,
		prog,
		"zeta\nu\nerre\ne\nelle\no\nenne\n",
		"z? u? r? e? l? o? n? zeta - u - zeta - zeta - u - erre - e - elle - elle - o - enne - e\n",
		"ZuzZurelLone")
}
