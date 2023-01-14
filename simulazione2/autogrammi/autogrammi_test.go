/*

Autogrammi
----------

Un *autogramma* è una frase che in qualche modo descrive (a volte mentendo) se stessa, ad esempio questi che seguono sono autogrammi onesti:
"questa frase contiene cinque parole"
oppure
"questa frase contiene: 10 parole, lunghezza massima: 9, lunghezza minima: 5, esclusi numeri"
o ancora
"questa lunghissima frase invece contiene: 14 parole, lunghezza massima: 11, lunghezza minima: 2, esclusi numeri si intende".

La frase "questa frase contiene tre parole" invece ovviamente è un autogramma bugiardo.

Dovete scrivere un programma (il cui file deve chiamarsi 'autogrammi.go') per verificare se frasi nella forma specificata qui sotto sono autogrammi onesti o meno.

Le frasi avranno tutte le seguenti componenti, nell'ordine:
- un preambolo arbitrario ma che termina sempre con la stringa "contiene:"
- l'indicazione del numero di parole della frase nella forma "contiene: <numero> "
- l'informazione sulla lunghezza della parola più lunga nella forma "massima: <numero>,"
- l'informazione sulla lunghezza della parola più corta nella forma "minima: <numero>,"
- e eventuale altro testo di lunghezza arbitraria

Si noti anche che:
- l'unico divisore di parola è lo spazio ' ' singolo
- la punteggiatura (caratteri possibili solo ',' e ':') è sempre attaccata alla fine delle parole, mai da sola

In particolare il programma dovrà definire le seguenti funzioni:

- CalcQuanteMinMax(frase string) (quante, min, max int)
	- conta le stringhe (separate da white space) presenti nella frase, escludendo i numeri (sequenze di sole cifre e eventuale punteggiatura attaccata alla fine)
	- calcola la lunghezza della parola più corta nella frase, **non considerando** l'eventuale punteggiatura come parte delle parole da misurare ("minima:" sarà considerata di lunghezza 6 e non 7), ed escludendo i numeri nella frase
	- calcola la lunghezza della parola più lunga nella frase, **non considerando** la punteggiatura come parte delle parole da misurare, ed escludendo i numeri nella frase
	- nota bene: per 'lunghezza' si intende il numero di byte
	- nota bene: in caso di stringa vuota, la funzione deve restituire (0,0,0)

- func TrovaNumDopoKeyword(frase, keyWord string) (num int)
	estrae da una frase il valore numerico indicato dopo una keyword e lo restituisce.
	Ad esempio, con la frase "le parole in questa frase in italiano sono == 8" e keyWord "==", restituisce 8; con la frase "questa contiene: 9" e keyword "contiene:", restituisce 9.
	Se la keyword non è presente, la funzione restituisce 0.

- func Autogramma(frase string) bool
	esamina la frase e verifica che essa sia un autogramma onesto o no; se lo è, restituisce true, false altrimenti

- func StampaAnalisiAutogramma(frase string)
	esamina la frase e stampa l'analisi effettuata secondo il formato:

		===  <frase>
		minimo dichiarato <min_dichiarato> contro minimo verificato <min_verificato>
		massimo dichiarato <max_dichiarato> contro massimo verificato <max_verificato>
		numero parole dichiarato <num_parole_dichiarato> contro numero parole verificato <num_parole_verificato>
		<onesto|bugiardo>

	senza indentazione, si veda anche esempio sottostante


- func main()
	 legge un file di testo il cui nome viene specificato a linea di comando;
	 il file contiene frasi (una per riga) che **potrebbero** essere autogrammi onesti;
	 bisogna analizzare ogni frase e stampare in output la frase letta, le informazioni ottenute dall'analisi, e "onesto" o "bugiardo" in corrispondenza delle analisi effettuate.
	 Le righe vuote vanno ignorate.

	 Se il nome del file non viene fornito, il programma stampa il messaggio "file name?" e termina.
	 Se l'apertura del file dà errore, il programma stampa il messaggio "file not found" e termina.


Esempio di esecuzione
---------------------

Se il file in input contiene le frasi:

questa frase contiene: 13 parole, lunghezza massima: 8, lunghezza minima: 5, esclusi numeri
questa più corta contiene: 11 parole, lunghezza massima: 9, lunghezza minima: 4, esclusi numeri

il programma stamperà:

===  questa frase contiene: 13 parole, lunghezza massima: 8, lunghezza minima: 5, esclusi numeri
minimo dichiarato 5 contro minimo verificato 5
massimo dichiarato 8 contro massimo verificato 9
numero parole dichiarato 13 contro numero parole verificato 10
bugiardo
===  questa più corta contiene: 11 parole, lunghezza massima: 9, lunghezza minima: 4, esclusi numeri
minimo dichiarato 4 contro minimo verificato 4
massimo dichiarato 9 contro massimo verificato 9
numero parole dichiarato 11 contro numero parole verificato 11
onesto

*/
package main

import (
	//"strings"
	//"log"
	//"fmt"
	//"os/exec"
	//"strings"

	"testing"
)

func TestCalcNorm(t *testing.T) {
	q, min, max := CalcQuanteMinMax("uno, due, tre")
	if q != 3 || min != 3 || max != 3 {
		t.FailNow()
	}
	q, min, max = CalcQuanteMinMax("un, due, tre, quattro")
	if q != 4 || min != 2 || max != 7 {
		t.FailNow()
	}
	q, min, max = CalcQuanteMinMax("conta le parole presenti nella stringa, la punteggiatura non fa parte della parola da contare")
	if q != 15 || min != 2 || max != 13 {
		t.FailNow()
	}

	q, min, max = CalcQuanteMinMax("punteggiatura non fa parte della parola da contare")
	if q != 8 || min != 2 || max != 13 {
		t.FailNow()
	}

}

func TestCalcLimite(t *testing.T) {
	q, min, max := CalcQuanteMinMax("")
	//fmt.Println(q, min, max)
	if q != 0 || min != 0 || max != 0 {
		t.FailNow()
	}
	q, min, max = CalcQuanteMinMax("uno")
	if q != 1 || min != 3 || max != 3 {
		t.FailNow()
	}
}

func TestTrovaNumParoleInFrase(t *testing.T) {
	if TrovaNumDopoKeyword("questa lunghissima frase invece contiene: 17 parole, lunghezza massima: 11, lunghezza minima: 2, esclusi numeri si intende", "contiene:") != 17 {
		t.FailNow()
	}
	if TrovaNumDopoKeyword("questa contiene: 5 parole, ...", "contiene:") != 5 {
		t.FailNow()
	}
}

func TestTrovaLunMaxInFrase(t *testing.T) {
	if TrovaNumDopoKeyword("questa lunghissima frase invece contiene: 17 parole, lunghezza massima: 11, lunghezza minima: 2, esclusi numeri si intende", "massima:") != 11 {
		t.FailNow()
	}
	if TrovaNumDopoKeyword("questa ... lunghezza massima: 21, lunghezza minima: ...", "massima:") != 21 {
		t.FailNow()
	}
}

func TestTrovaLunMinInFrase(t *testing.T) {
	if TrovaNumDopoKeyword("questa lunghissima frase invece contiene: 17 parole, lunghezza massima: 11, lunghezza minima: 2, esclusi numeri si intende", "minima:") != 2 {
		t.FailNow()
	}
	if TrovaNumDopoKeyword("questa lunghissima frase invece contiene: ... lunghezza minima: 6, esclusi numeri si intende", "minima:") != 6 {
		t.FailNow()
	}
}
func TestAutogramma(t *testing.T) {
	if !Autogramma("questa lunghissima frase invece contiene: 14 parole, lunghezza massima: 11, lunghezza minima: 2, esclusi numeri si intende") {
		t.FailNow()
	}
	if Autogramma("questa lunghissima frase invece contiene: 16 parole, lunghezza massima: 11, lunghezza minima: 2, esclusi numeri si intende") {
		t.FailNow()
	}

	if Autogramma("questa frase contiene: 13 parole, lunghezza massima: 8, lunghezza minima: 5, esclusi numeri") {
		t.FailNow()
	}
	if Autogramma("questa frase contiene 19 parole, lunghezza massima 8, lunghezza minima 5, esclusi numeri") {
		t.FailNow()
	}
}

var prog = "./autogrammi"

func TestMain(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"autogrammi.output",
		"autogrammi.input")
}

func TestEsempio(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"esempio.output",
		"esempio.input")
}
func TestMainNoArgs(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"file name?\n")
}

func TestMainNoFile(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"file not found\n",
		"disicurononce")
}
