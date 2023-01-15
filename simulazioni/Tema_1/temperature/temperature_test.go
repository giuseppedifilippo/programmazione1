/*
Temperature
===========

Scrivere un programma 'temperature.go' che legge da standard input
una sequenza (terminata da EOF) di almeno una temperatura e stampa
la temperatura minima e il numero di volte che è stata registrata,
la temperatura massima e il numero di volte che è stata registrata,
le temperature che si sono verificate con la massima frequenza.

Esempio
-------
su input
2 3 4 -2 2 6 12 8 3 1 -3
0 -2 5 2 4 1 -1 0 7 11 8 3

produce in output

max : 12, 1 volte; min : -3, 1 volte
-3:1 -2:2 -1:1 0:2 1:2 2:3 3:3 4:2 5:1 6:1 7:1 8:2 11:1 12:1
temperature [2 3] con frequenza 3, la massima frequenza
*/

package main

import "testing"

var prog = "./temperature"

func TestEsempio(t *testing.T) {
	lanciaGenerica(t, prog, "2\n3\n4\n-2\n2\n6\n12\n8\n3\n1\n-3\n0\n-2\n5\n2\n4\n1\n-1\n0\n7\n11\n8\n3\n", "max : 12, 1 volte; min : -3, 1 volte\n-3:1 -2:2 -1:1 0:2 1:2 2:3 3:3 4:2 5:1 6:1 7:1 8:2 11:1 12:1\ntemperature [2 3] con frequenza 3, la massima frequenza\n")
}

func TestAltro(t *testing.T) {
	lanciaGenerica(t, prog, "1\n1\n1\n1\n12\n11\n10\n1\n1\n-1\n", "max : 12, 1 volte; min : -1, 1 volte\n-1:1 1:6 10:1 11:1 12:1\ntemperature [1] con frequenza 6, la massima frequenza\n")
}
