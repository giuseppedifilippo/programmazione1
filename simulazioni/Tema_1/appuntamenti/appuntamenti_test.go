/*
Appuntamenti
------------

Scrivere un programma (il cui file deve chiamarsi 'appuntamenti.go') dotato di:

- un tipo Appuntamento con tre campi int (giorno, oraInizio, oraFine, dichiarati in quest'ordine) che rappresentano un giorno dell'anno, e l'ora di inizio e di fine  dell'appuntamento.
Si considerano per semplicità solo appuntamenti che iniziano e finiscono nella stessa giornata e a ore precise (intere).

- una funzione
	NuovoAppuntamento(gg, h1, h2 int) (Appuntamento, bool)
che, se i parametri sono validi (giorno tra 1 e 366, ora inizio precedente o uguale a ora di fine, entrambe nell'intervallo [0,24])
crea un appuntamento corrispondente a questi dati e lo restituisce insieme al valore 'true',
altrimenti restituisce un Appuntamento "zero-value" e 'false'.

- una funzione
    CheckSovrapposizioneAppuntamenti(app1, app2 Appuntamento) bool
che, dati due appuntamenti che si assume siano validi, restituisce 'true' se si sovrappongono (anche parzialmente), 'false' altrimenti.

- una funzione
	AggiungiAppuntamento(app Appuntamento, agenda *[]Appuntamento) bool
che, se l'appuntamento app non si sovrappone con NESSUNO degli altri già presenti in agenda, lo aggiunge alla stessa restituendo 'true', altrimenti non fa nulla e restituisce 'false'

- una funzione
    AppuntamentiGiornata(gg int, agenda []Appuntamento) []Appuntamento
che, dati un giorno e una lista di appuntamenti, restituisce gli appuntamenti del giorno

- una funzione main() che crea un'agenda vuota, legge da standard input
delle righe che contengono ognuna tre interi che rappresentano, nell'ordine,

giorno ora-inizio ora-fine

(come separatore si usano degli spazi, eventualmente ripetuti)

e per ciascuna aggiunge in agenda il corrispondente appuntamento, se valido e non in sovrapposizione con altri.
Una volta raggiunto EOF (che su standard input, su sistemi Linux, si ottiene premendo CTRL D), il programma visualizza su standard output gli appuntamenti in agenda, nell'ordine in cui sono stati inseriti.

Si assuma che da standard input siano inseriti solo numeri interi (non occorre fare controlli).

*/
package main

import (
	//"strings"
	//"log"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

func TestAppuntamentoNorm(t *testing.T) {
	_, correct := NuovoAppuntamento(300, 11, 12)
	if !correct {
		t.Fail()
	}
}

func TestAppuntamentoWrongDay(t *testing.T) {
	_, correct := NuovoAppuntamento(400, 11, 12)
	if correct {
		t.Fail()
	}
}

func TestAppuntamentoWrongTime(t *testing.T) {
	_, correct := NuovoAppuntamento(300, 11, 32)
	if correct {
		t.Fail()
	}
}

func TestSovrAppuntamentoPost(t *testing.T) {
	uno, _ := NuovoAppuntamento(300, 11, 13)
	due, _ := NuovoAppuntamento(300, 12, 14)
	if !CheckSovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
}

func TestSovrAppuntamentoFake(t *testing.T) {
	uno, _ := NuovoAppuntamento(302, 11, 13)
	due, _ := NuovoAppuntamento(300, 11, 13)
	if CheckSovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
}

func TestSovrAppuntamentoIdentici(t *testing.T) {
	uno, _ := NuovoAppuntamento(300, 11, 13)
	due, _ := NuovoAppuntamento(300, 11, 13)
	if !CheckSovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
}

func TestSovrAppuntamentoPre(t *testing.T) {
	uno, _ := NuovoAppuntamento(300, 11, 13)
	due, _ := NuovoAppuntamento(300, 10, 12)
	if !CheckSovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
}

func TestSovrAppuntamentoPreSameEnd(t *testing.T) {
	uno, _ := NuovoAppuntamento(300, 11, 13)
	due, _ := NuovoAppuntamento(300, 10, 13)
	if !CheckSovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
}

func TestSovrAppuntamentoFakeContigui(t *testing.T) {
	uno, _ := NuovoAppuntamento(300, 11, 13)
	due, _ := NuovoAppuntamento(300, 13, 15)
	if CheckSovrapposizioneAppuntamenti(uno, due) {
		t.Fail()
	}
}

func TestAppuntamentoReversedTime(t *testing.T) {
	_, correct := NuovoAppuntamento(300, 12, 10)
	if correct {
		t.Fail()
	}
}

func TestAgendaNorm(t *testing.T) {
	agenda := make([]Appuntamento, 0)
	for index := 300; index < 400; index += 20 {
		a, ok := NuovoAppuntamento(index, 11, 13)
		if ok {
			AggiungiAppuntamento(a, &agenda)
			success := AggiungiAppuntamento(a, &agenda)

			if success {
				t.Fail()
			}
		}
	}

	if len(agenda) != 4 {
		t.Fail()
	}

	if len(AppuntamentiGiornata(300, agenda)) != 1 {
		t.Fail()
	}
}

func TestMain1(t *testing.T) {
	subproc := exec.Command("./appuntamenti") // presuppone che sia già stato compilato
	input := "200 11 12\n201 11 12\n201 11 13\n201 13 12\n"
	subproc.Stdin = strings.NewReader(input)
	out, err := subproc.CombinedOutput()

	if err != nil {
		// verificare se uscito per esplicito os.Exit()
		//t.Errorf("Fallito: %s\n", stderr)
		fmt.Printf(">>> Attenzione! Uscito con errore: %s\n>>> (non è un test fallito se l'esercizio richiedeva uscita esplicita con exit code)\n", err)
	}

	fmt.Printf("Input:\n%s\n", input)
	fmt.Printf("Output:\n%s\n", string(out))
	if string(out) != "[{200 11 12} {201 11 12}]\n" {
		t.Fail()
	}
	subproc.Process.Kill()
}
