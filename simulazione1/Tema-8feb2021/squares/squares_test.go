/*

Scrivere un programma 'squares.go' dotato di:

- una funzione
	func NestedSquares(n int) string
che costruisce e restituisce una stringa corrispondente ad una figura
formata da un quadrato di '*' di dimensione n (con n > 4), che racchiude
un quadrato di '*' di dimensione 'n - 4', separato da quello esterno da una
"cornice" di spazi di spessore unitario (si vedano gli esempi).
La stringa restituita NON deve terminare con '\n'.
Se n < 5, la funzione restituisce la stringa vuota

- una funzione main() che legge da linea di comando un intero
e produce su standard output la figura di quadrati della misura fornita.
Se non ci sono dati sulla linea di comando, il programma stampa il messaggio
"too few arguments" e termina.
Se il valore passato da linea di comando non è un intero, il programma stampa il messaggio
"required integer value" e termina.


Esempi
------
per n = 5
*****
*   *
* * *
*   *
*****

per n = 6
******
*    *
* ** *
* ** *
*    *
******

per n = 7
*******
*     *
* *** *
* * * *
* *** *
*     *
*******

per n = 8
********
*      *
* **** *
* *  * *
* *  * *
* **** *
*      *
********

per n = 9
*********
*       *
* ***** *
* *   * *
* *   * *
* *   * *
* ***** *
*       *
*********
etc..
*/

package main

import (
	"strings"
	"testing"
)

var prog = "./squares"

func TestSquares(t *testing.T) {
	picture := NestedSquares(4)
	if len(picture) != 0 {
		t.Errorf("if size < 5 should return the empty string")
	}
	expected := []string{
		"*****\n*   *\n* * *\n*   *\n*****",
		"******\n*    *\n* ** *\n* ** *\n*    *\n******",
		"*******\n*     *\n* *** *\n* * * *\n* *** *\n*     *\n*******",
		"********\n*      *\n* **** *\n* *  * *\n* *  * *\n* **** *\n*      *\n********",
		"*********\n*       *\n* ***** *\n* *   * *\n* *   * *\n* *   * *\n* ***** *\n*       *\n*********",
	}
	for i := 0; i < len(expected); i++ {
		got := NestedSquares(5 + i)
		if strings.Compare(expected[i], got) != 0 {
			t.Errorf("expected:\n%s\ngot:\n%v", expected[i], got)
		}
	}

}


func TestNoArgs(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "too few arguments\n")
}

func TestWrongArg(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "required integer value\n","aldkhaslk")
}


func TestWrongArgFloat(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "required integer value\n","7.7")
}

func TestWrongArgFloatComma(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "required integer value\n","7,7")
}


func TestMain(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "******\n*    *\n* ** *\n* ** *\n*    *\n******\n","6")
}
