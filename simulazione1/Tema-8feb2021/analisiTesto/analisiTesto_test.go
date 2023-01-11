/*

Analisi di testo (NOTA: CASE SENSITIVE O NO?)
================

Scrivere un programma in Go (il file deve chiamarsi 'analisiTesto.go') dotato di

- un tipo Vocabolario  in cui salvare le stringhe (sequenze di caratteri separate da white spaces) di un testo e, per ciascuna, il numero di volte che essa appare nel testo

- un metodo String per Vocabolario che stampa il Vocabolario in ordine lessicografico, un lemma per riga seguito da spazio,':', spazio e dalla sua frequenza. Ad esempio:
Lorem : 3
ever : 1

- una funzione

		statistiche(voc Vocabolario) (int, int, float64)

  che restituisce la frequenza più bassa, la frequenza più alta e la media delle frequenze delle stringhe del Vocabolario

- una funzione

		selezione(voc Vocabolario, char rune) []string

	che restituisce la lista delle stringhe contenute nel Vocabolario che iniziano con il carattere (runa) char

- una funzione

		main()

  che legge da linea di comando il nome di un file e un carattere (runa);
  legge il testo contenuto nel file e crea un Vocabolario delle stringhe contenute nel testo, ciascuna con la frequenza con cui appare nel testo;
  stampa il Vocabolario creato, i risultati restituiti dalla funzione statistiche e la lista di stringhe restituita dalla funzione selezione (vedi esempio).
  Se non viene passato un argomento per il nome del file e/o un carattere, il programma deve stampare il messaggio "indicare nome file e iniziale".
  Se si verifica un errore nell'apertura del file, il programma deve stampare l'errore restituito per la mancata apertura del file.

Esempio
-------
(sul file uno.input messo a disposizione)

$ ./analisiTesto uno.input s
Bëatrice : 1
Quant’ : 1
al : 1
bene : 1
che : 2
ch’era : 1
color, : 1
convenia : 1
da : 1
dentro : 1
di : 1
dov’ : 1
entra’mi, : 1
esser : 1
in : 1
io : 1
lucente : 1
lume : 1
l’atto : 1
ma : 1
meglio, : 1
non : 2
parvente! : 1
per : 3
quel : 1
quella : 1
scorge : 1
si : 1
sol : 1
sporge. : 1
subitamente : 1
suo : 1
sé : 1
sì : 2
tempo : 1
È : 1
1 3 1.1388888888888888
[scorge si sol sporge. subitamente suo sé sì]
*/

package main

import (
	"fmt"
	"math"
	"testing"
)

var prog = "./analisiTesto"

func TestComeDaEsempio(t *testing.T) {
	lanciaGenericaConFileOutAtteso(t, prog, "nil", "uno.expected", "uno.input", "s")
}

func TestBreve(t *testing.T) {
	lanciaGenericaConFileOutAtteso(t, prog, "nil", "lorem.expected", "lorem.input", "i")
}

func TestRepeated(t *testing.T) {
	lanciaGenericaConFileOutAtteso(t, prog, "nil", "repeated.expected", "repeated.input", "U")
}

func TestTutteDiverse(t *testing.T) {
	lanciaGenericaConFileOutAtteso(t, prog, "nil", "uniques.expected", "uniques.input", "R")
}

func TestWhitespaces(t *testing.T) {
	lanciaGenericaConFileOutAtteso(t, prog, "nil", "whitespaces.expected", "whitespaces.input", "I")
}

func TestNoArgs(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "indicare nome file e iniziale\n")
}

func TestUnArg(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "indicare nome file e iniziale\n", "qualcosa")
}

func TestFileNonesistente(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "open sicuroCheNonCe: no such file or directory\n", "sicuroCheNonCe", "o")
}

func TestEsistenzaTipo(t *testing.T) {
	var voc Vocabolario
	voc = make(map[string]int)
	fmt.Println(voc)
	//fallisce solo per non compilazione
}

func TestEsistenzaStat(t *testing.T) {
	var voc Vocabolario
	voc = make(map[string]int)
	_, _, _ = statistiche(voc)
}

func TestStatUnaParola(t *testing.T) {
	var voc Vocabolario
	voc = make(map[string]int)
	voc["pippo"]++

	min, max, med := statistiche(voc)

	if min != 1 || med != 1 || max != 1 {
		t.Fail()
	}
}

func TestStatParoleFreq1(t *testing.T) {
	var voc Vocabolario
	voc = make(map[string]int)

	voc["pippo"]++
	voc["pluto"]++
	voc["paperino"]++

	min, max, med := statistiche(voc)

	if min != 1 || med != 1 || max != 1 {
		t.Fail()
	}
}

func TestTestoNormaleMediaIntera(t *testing.T) {
	var voc Vocabolario
	voc = make(map[string]int)

	voc["pippo"] = 5
	voc["pluto"] = 15

	min, max, med := statistiche(voc)

	if min != 5 || math.Abs(med-10) > .01 || max != 15 {
		t.Fail()
	}
}

func TestTestoNormaleMediaFloat(t *testing.T) {
	var voc Vocabolario
	voc = make(map[string]int)

	voc["pippo"] = 5
	voc["pluto"] = 12

	min, max, med := statistiche(voc)

	if min != 5 || math.Abs(med-8.5) > .01 || max != 12 {
		t.Fail()
	}
}

func TestEsistenzaSelezione(t *testing.T) {
	var voc Vocabolario
	voc = make(map[string]int)
	_ = selezione(voc, 'R')
}

func TestSelezioneVocVuoto(t *testing.T) {
	var voc Vocabolario
	voc = make(map[string]int)

	rest := selezione(voc, 'R')

	if len(rest) != 0 {
		t.Fail()
	}
}

func TestSelezioneCaseSbagliato(t *testing.T) {
	var voc Vocabolario
	voc = make(map[string]int)

	voc["pippo"] = 5
	voc["pluto"] = 15

	rest := selezione(voc, 'P')

	if len(rest) != 0 {
		t.Fail()
	}
}

func TestSelezioneNormale(t *testing.T) {
	var voc Vocabolario
	voc = make(map[string]int)

	voc["pippo"] = 5
	voc["Pluto"] = 15
	voc["foo"] = 50
	voc["bar"] = 154
	voc["pluto"] = 3

	rest := selezione(voc, 'p')

	if len(rest) != 2 {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	var voc Vocabolario
	voc = make(map[string]int)

	voc["pippo"] = 5
	voc["Pluto"] = 15
	voc["foo"] = 50
	voc["bar"] = 154
	voc["pluto"] = 3

	fmt.Println(voc)

	if voc.String() != "Pluto : 15\nbar : 154\nfoo : 50\npippo : 5\npluto : 3" {
		t.Fail()
	}
}

/*
// TODO
func TestString(t *testing.T)

func TestStatistiche(t *testing.T)

func TestSelezione(t *testing.T)

func TestNoArgs(t *testing.T)

func TestNoFile(t *testing.T)

*/
