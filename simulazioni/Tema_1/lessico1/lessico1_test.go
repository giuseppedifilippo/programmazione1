/*
Lessico v1
----------

Scrivere un programma 'lessico1.go' che si comporta come segue.

1) Stampa il seguente menu di opzioni:
  + (legge e memorizza)
  ? (numeri di riga in cui è comparsa la parola data)
  m (stampa le lunghezze min e max)
  p (stampa le parole memorizzate)

2) Legge da standard input un testo di un numero arbitrario di righe e termina quando riceve un "end of file" (cioè EOF, da tastiera pressione di 'CTRL-d' dopo aver dato <invio>).

Quando la riga letta **inizia** con:
- "+" : il programma stampa "alimento il dizionario"
- "?" : il programma stampa "consulto il dizionario"
- "m" : il programma stampa "lunghezza min e max"
- "p" : il programma stampa "stampo il dizionario ordinato"
- per qualsiasi altro carattere il programma stampa "opzione non valida"

Esempio di esecuzione
---------------------

date in input le seguenti righe di testo:

+ la befana ha il fazzoletto e la gonna rattoppata
? ha
n
+ ma quest'anno poverina
+ la befana è raffreddata!
? ha il
m
p

il programma produce il seguente output:

+ (legge e memorizza)
? (numeri di riga in cui è comparsa la parola data)
m (stampa le lunghezze min e max)
p (stampa le parole memorizzate)
alimento il dizionario
consulto il dizionario
opzione non valida
alimento il dizionario
alimento il dizionario
consulto il dizionario
lunghezza min e max
stampo il dizionario ordinato

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

var prog = "./lessico1"

func TestComeDaEsempio(t *testing.T) {
	//LanciaGenerica(t, prog, "+ la befana ha il fazzoletto e la gonna rattoppata\n? ha\nn\n+ ma quest'anno poverina\n+ la befana è raffreddata!\n? ha il\nm\np", "+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\nm (stampa le lunghezze min e max)\np (stampa le parole memorizzate)\nalimento il dizionario\nconsulto il dizionario\nopzione non valida\nalimento il dizionario\nalimento il dizionario\nconsulto il dizionario\nlunghezza min e max\nstampo il dizionario ordinato\n")
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"esempio.input",
		"esempio.output")
}
func TestDue(t *testing.T) {
	//LanciaGenerica(t, prog, "+ aaa bbb ccc\np\nm\nq\n? aaa", "+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\nm (stampa le lunghezze min e max)\np (stampa le parole memorizzate)\nalimento il dizionario\nstampo il dizionario ordinato\nlunghezza min e max\nopzione non valida\nconsulto il dizionario\n")
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"due.input",
		"due.output")
}
