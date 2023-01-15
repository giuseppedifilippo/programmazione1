/*

Lunghezze
=========

Scrivere un programma 'lunghezze.go' che legge da file,
il cui nome è passato su linea di comando, un testo
e stampa prima le stringhe (separate da 'whitespace', sulla stessa riga)
di lunghezza 1, poi a capo quelle di lunghezza 2, e così via.
Ogni riga inizia con un numero che indica la lunghezza, seguito da " : "
(per il formato dell'output, vedere l'esempio).

Se il programma viene lanciato con un numero di argomenti
diverso da 1, deve terminare stampando "Fornire 1 nome di file".
Se invece il file non esiste,  deve terminare stampando "File non accessibile".

Esempio
-------

$ go run lunghezze.go altro.input
1 : e
2 : a,
3 : qui quo qua
4 : qua.
5 : Pippo Pluto ciao, tutti
6 : questi
8 : Paperino
11 : arrivederci

*/

package main

import (
	"fmt"
	"os/exec"
	"testing"
	//"strings"
	//"log"
)

var progname = "./lunghezze"

func TestMainMultiplo(t *testing.T) {
	//lancia(t, "1 [a a a a a]\n2 [is of an of it to It It in of of of of in by or If to of to be in of on to as on It of of to is or]\n3 [the and has the the and has not but the was the the and are but the you are use you the All the the the the 200 The]\n4 [text been text ever when took type make type only five also leap into with more with like many have some look even need sure tend this true uses over with free from etc.]\n5 [Lorem Ipsum dummy Lorem Ipsum dummy since book. 1960s Lorem Ipsum Aldus Lorem There Lorem Ipsum form, words which don't going Lorem there isn't text. Lorem Ipsum first Latin model Lorem Ipsum which looks Lorem Ipsum words]\n6 [simply 1500s, galley sheets Ipsum. Ipsum, hidden middle repeat chunks making words, always]\n7 [unknown printer release desktop humour, passage handful humour,]\n8 [printing standard specimen survived Letraset recently software versions passages majority suffered injected slightly anything Internet combined sentence generate injected]\n9 [industry. scrambled remaining passages, PageMaker including generator Internet. generated therefore]\n10 [industry's centuries, electronic unchanged. containing publishing variations available, alteration randomised generators predefined necessary, dictionary]\n11 [typesetting essentially popularised believable. structures, reasonable. repetition,]\n12 [typesetting, embarrassing]\n", "lorem.input")

	LanciaGenericaConFileOutAtteso(t, progname, "NIENTE", "altro.expected", "altro.input")
	//lancia(t, "1 [e]\n2 [a,]\n3 [qui quo qua]\n4 [qua.]\n5 [Pippo Pluto ciao, tutti]\n6 [questi]\n", "altro.input")
	LanciaGenericaConFileOutAtteso(t, progname, "NIENTE", "lorem.expected", "lorem.input")

}

func TestFileNonAccessibile(t *testing.T) {
	LanciaGenerica(t, progname, "NIENTE", "File non accessibile\n", "NONESISTENTE")
	//lancia(t, "File non accessibile\n", "nulla")
}

func TestNoArgs(t *testing.T) {
	LanciaGenerica(t, progname, "NIENTE", "Fornire 1 nome di file\n")
	//lancia(t, "Fornire 1 nome di file\n")
}

func lancia(t *testing.T, expected string, in ...string) {
	fmt.Println("### Questo test NON verifica output atteso! (presuppone che il sorgente da testare sia già stato compilato)")

	subproc := exec.Command("./lunghezze", in...)

	stdout, err := subproc.CombinedOutput()

	if err != nil {
		// verificare se uscito per esplicito os.Exit()
		//t.Errorf("Fallito: %s\n", stderr)
		fmt.Printf(">>> Attenzione! Uscito con errore: %s\n>>> (non è un test fallito se l'esercizio richiedeva uscita esplicita con exit code)\n", err)
	}

	fmt.Printf("Actual output:\n%s\n", string(stdout))

	/*
		fmt.Printf("Expected output:\n%s\n", expected)

		if string(stdout) != expected {
			fmt.Println(">>> FAIL!")
			t.Fail()
		}
	*/

	subproc.Process.Kill()
}
