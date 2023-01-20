/*
Dato un numero intero >=0 su linea di comando, stampare, uno per riga, i numeri formati da sottosequenze di cifre consecutive non decrescenti. Ad esempio dato il numero 9245182337, il programma deve produrre in output:
9
245
18
2337
*/

package main

import "testing"

var prog = "./sottosequenze"

func TestEsempio(t *testing.T) {
	lanciaGenerica(t, prog, "", "9\n245\n18\n2337\n", "9245182337")
}

func TestAltro(t *testing.T) {
	lanciaGenerica(t, prog, "", "129\n8\n37\n6\n129\n8\n7\n3\n16\n29\n8\n37\n1269\n3\n123\n126\n37\n12\n", "12983761298731629837126931231263712")
}
