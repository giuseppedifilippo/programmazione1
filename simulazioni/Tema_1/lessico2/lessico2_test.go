/*
Lessico v2
----------

Copiare il programma 'lessico1.go' in un file 'lessico2.go'

Il programma 'lessico2.go' alimenta un "dizionario" con le parole del testo e le righe (i numeri di riga) in cui ciascuna parola compare; inoltre permette di consultare e stampare il "dizionario".

In particolare, si richiede di dotare il programma 'lessico2.go' delle seguenti funzioni, che andranno usate nel main:

- func PrintMenu()
che stampa il menu delle opzioni, cioè il seguente menu:
+ (legge e memorizza)
? (numeri di riga in cui è comparsa la parola data)
m (stampa le lunghezze min e max)
p (stampa le parole memorizzate)

- func PrintDict(m map[string][]int)
che, data una mappa m, stampa in ordine lessicografico, una per riga, le chiavi della mappa, ciascuna seguita, sulla stessa riga, da uno spazio, un ':', uno spazio e dal suo valore (vedi Esempio sotto)

- AddToDict(line string, n int, dict map[string][]int)
(vedi sotto, opzione '+') che riceve la riga letta che inizia con '+' (line) insieme alla sua posizione nel testo (numero di riga n) e gestisce l'aggiornamento del "dizionario".
Si può assumere, senza fare controlli, che la riga abbia il formato atteso, cioè: '+' seguito da uno spazio e poi da un testo

- Stats(dict map[string][]int) (int, int)
che restiluisce la lunghezza della stringa più corta e la lunghezza della stringa più lunga (in quest'ordine), presenti in dict

Aggiungere inoltre al programma 'lessico2.go', per ogni opzione del menu, dopo le istruzioni di stampa della versione 1, le funzionalità specificate qui sotto.

Quando la riga letta inizia con:
- "+" (alimenta dizionario): il programma usa il rimanente della riga e memorizza in un "dizionario" le parole (stringhe separate da white space) che la costituiscono e il numero di riga. La prima riga letta è la numero 1;
- "?" (consulta dizionario): il programma usa il rimanente della riga come stringa per la consultazione del dizionario e stampa i numeri di riga associati a tale stringa;
- "m" (determina min e max): il programma stampa la lunghezza della stringa più corta e la lunghezza della stringa più lunga (in quest'ordine), presenti nel dizionario
- "p" (print): il programma stampa le parole presenti nel  "dizionario", una per riga e in ordine lessicografico, ciascuna seguita dalla lista dei numeri di riga in cui la parola è comparsa (per il formato, vedi Esempio sotto).


Esempio di esecuzione
---------------------

date in input le seguenti righe di testo:

+ aa bbbb ccc
+ aa
q
p
m
? aa
? aa bbbb

il programma produce il seguente output:

+ (legge e memorizza)
? (numeri di riga in cui è comparsa la parola data)
m (stampa le lunghezze min e max)
p (stampa le parole memorizzate)
alimento il dizionario
alimento il dizionario
opzione non valida
stampo il dizionario ordinato
aa : [1 2]
bbbb : [1]
ccc : [1]
lunghezza min e max
2 4
consulto il dizionario
parola: aa
righe [1 2]
consulto il dizionario
parola: aa bbbb
righe []


*/
package main

import (
	//"fmt"
	//"log"
	//"os/exec"

	//"strings"

	"fmt"
	"io/ioutil"
	"os"
	"testing"
	//"bytes"
	//"io"
	//"os"
)

var prog = "./lessico2"

/*
func TestComeDaEsempio(t *testing.T) {
	LanciaGenerica(t, prog, "+ aaa bbb ccc\np\n+ aaa bbb ccc\n+ uno due tre\np\n+ h j k\np", "+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\np (Stampa le parole)\naaa : [1]\nbbb : [1]\nccc : [1]\naaa : [1 2]\nbbb : [1 2]\nccc : [1 2]\ndue : [3]\ntre : [3]\nuno : [3]\naaa : [1 2]\nbbb : [1 2]\nccc : [1 2]\ndue : [3]\nh : [4]\nj : [4]\nk : [4]\ntre : [3]\nuno : [3]\n")
}

func TestSecondo(t *testing.T) {
	LanciaGenerica(t, prog, "+ la befana ha il fazzoletto e la gonna rattoppata\np\n+ ma quest'anno poverina la befana è raffreddata\n? la\n? befana\n? ha il\np", "+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\np (Stampa le parole)\nbefana : [1]\ne : [1]\nfazzoletto : [1]\ngonna : [1]\nha : [1]\nil : [1]\nla : [1 1]\nrattoppata : [1]\nparola: la\nrighe [1 1 2]\nparola: befana\nrighe [1 2]\nparola: ha il\nrighe []\nbefana : [1 2]\ne : [1]\nfazzoletto : [1]\ngonna : [1]\nha : [1]\nil : [1]\nla : [1 1 2]\nma : [2]\npoverina : [2]\nquest'anno : [2]\nraffreddata : [2]\nrattoppata : [1]\nè : [2]\n")
}

*/

func TestPrintDict(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	diz := make(map[string][]int)
	diz["uno"] = append(diz["uno"], 6)
	diz["uno"] = append(diz["uno"], 7)

	diz["due"] = append(diz["due"], 3)
	diz["due"] = append(diz["due"], 9)

	diz["tre"] = append(diz["tre"], 5)

	PrintDict(diz)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	//fmt.Printf("Captured: %s", out)

	if string(out) != "due : [3 9]\ntre : [5]\nuno : [6 7]\n" {
		t.FailNow()
	}
}

func TestUno(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"uno.input",
		"uno.output")
}

func TestPrintMenu(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintMenu()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	//fmt.Printf("Captured: %s", out)

	if string(out) != "+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\nm (stampa le lunghezze min e max)\np (stampa le parole memorizzate)\n" {
		t.FailNow()
	}
}

func TestAdd(t *testing.T) {
	dict := make(map[string][]int)
	AddToDict("+ ciao", 5, dict)
	s := fmt.Sprintln(dict)
	//fmt.Println(s)
	if s != "map[ciao:[5]]\n" {
		t.FailNow()
	}
}

func TestStats(t *testing.T) {
	dict := make(map[string][]int)
	AddToDict("+ ciao", 5, dict)
	AddToDict("+ cip", 13, dict)
	AddToDict("+ topolino", 25, dict)
	AddToDict("+ pippo", 54, dict)

	min, max := Stats(dict)
	fmt.Println(min, max)

	if min != 3 || max != 8 {
		t.FailNow()
	}
}

func TestComeDaEsempio(t *testing.T) {
	//LanciaGenerica(t, prog, "+ aa bbbb ccc\n+ aa\nq\np\nm\n? aa\n? aa bbbb", "+ (legge e memorizza)\n? (numeri di riga in cui è comparsa la parola data)\nm (stampa le lunghezze min e max)\np (stampa le parole memorizzate)\nalimento il dizionario\nalimento il dizionario\nopzione non valida\nstampo il dizionario ordinato\naa : [1 2]\nbbbb : [1]\nccc : [1]\nlunghezza min e max\n2 4\nconsulto il dizionario\nparola: aa\nrighe [1 2]\nconsulto il dizionario\nparola: aa bbbb\nrighe []\n")
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"esempio.input",
		"esempio.output")
}
