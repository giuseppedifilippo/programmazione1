/*

Dispensa
--------

Si scriva un programma (il file deve chiamarsi 'dispensa.go') per gestire una dispensa di un ristorante.

Il programma leggerà da un file CSV (Comma Separated Values) un "diario" degli acquisti e consumi per poi offrire all'utente la possibilità di conteggiare le rimanenze.

Il programma dovrà essere dotato delle seguenti:

- una funzione AggiornaDispensa(dispensa map[string]float64, filename string) bool
	che aggiorna una dispensa prendendo i dati da un file CSV il cui nome è passato come parametro.
	In particolare il file conterrà gli acquisti nel formato "approv,<prodotto>,<peso>" e i consumi nel formato "uso,<prodotto>,<peso>", dove le quantità (pesi) sono sempre espresse in Kg. (Non vanno fatti controlli sul formato dei dati, si può assumere che siano sempre esattamente come nel file di esempio).
	Se il file non esiste o non è leggibile, la funzione deve terminare subito restituendo 'false'.
	Nel caso in cui un prodotto dovesse "andare in negativo"  durante l'aggiornamento della dispensa, la funzione dovrà terminare subito e restituire 'false', mentre restituirà 'true' in caso contrario.

- una funzione Rimanenza(dispensa map[string]float64, alimento string) float64
	che, data una dispensa opportunamente popolata di provviste, restituisce la rimanenza in Kg dell'alimento passato come parametro; restituisce 0 se il prodotto non è presente in lista.

- una funzione main()
	che legge ed elabora un file in formato CSV che contiene i dati relativi ad acquisti e consumi di una dispensa.
	Il nome del file è passato come primo argomento sulla linea di comando.
	In caso di presenza di argomenti oltre al nome del file, il programma stampa poi la rimanenza degli alimenti passati come argomenti successivi.
	In caso di mancanza di argomenti oltre al nome del file, il programma deve stampare solo "file corretto" se nessun prodotto è andato in negativo, e "file non corretto" in caso contrario.


Esempi di esecuzione
---------------------

$ ./dispensa dispensa.csv zucchero miele peperoncino
zucchero 14.177
miele 28.046
peperoncino 9.452

$ ./dispensa dispensaErrata.csv
file non corretto

$ ./dispensa dispensa.csv
file corretto


*/
package main

import (
	"fmt"
	"testing"
)

var prog = "./dispensa"

func TestMain(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"zucchero 14.177\nmiele 28.046\npeperoncino 9.452\n",
		"dispensa.csv", "zucchero", "miele", "peperoncino")
}

func TestMainCorretto(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"file corretto\n",
		"dispensa.csv")
}

func TestMainNonCorretto(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"file non corretto\n",
		"dispensaErrata.csv")
}

func TestAggiorna(t *testing.T) {
	dispensa := make(map[string]float64)
	successo := AggiornaDispensa(dispensa, "dispensa.csv")
	if !successo {
		fmt.Println("dispensa erroneamente rilevata in negativo!")
		t.FailNow()
	}
	if !(len(dispensa) > 0) {
		fmt.Println("dispensa erroneamente vuota!")
		t.FailNow()
	}
	if dispensa["zucchero"] < 0 {
		fmt.Println("errore rimanenza!")
		t.FailNow()
	}
}

func TestAggiornaErrato(t *testing.T) {
	dispensa := make(map[string]float64)
	successo := AggiornaDispensa(dispensa, "dispensaErrata.csv")
	if successo {
		fmt.Println("dispensa in negativo non rilevata!")
		t.FailNow()
	}
}

func TestAggiornaFileMancante(t *testing.T) {
	dispensa := make(map[string]float64)
	successo := AggiornaDispensa(dispensa, "disicurononce")
	if successo {
		fmt.Println("file non disponibile non rilevato!")
		t.FailNow()
	}
}

func TestRimanenza(t *testing.T) {
	dispensa := make(map[string]float64)
	successo := AggiornaDispensa(dispensa, "dispensa.csv")
	r := Rimanenza(dispensa, "zucchero")
	if !successo || !intorno(r, 14.177) {
		fmt.Println("errore rimanenza!", r)
		t.FailNow()
	}
}
