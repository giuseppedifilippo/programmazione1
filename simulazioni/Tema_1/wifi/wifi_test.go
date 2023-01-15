/*

Wifi
----

Si scriva un programma (il file deve chiamarsi 'wifi.go') per gestire l'analisi dei segnali wifi di un ambiente.
Il programma dovrà essere dotato delle seguenti:

- una struttura Wifi con campi (dichiarati in quest'ordine):
	ssid      	string
	channel   	int
	frequency 	int
	signal_dBm  int

- una funzione NewWifi(ssid string, channel int, frequency int, signal_dBm int) Wifi,bool
	che, se i valori passati come parametri rispettano le seguenti condizioni:
	- ssid non vuoto
	- channel
		- tra 1 e 14 (compresi) SE la frequency è tra 2412 e 2484 (compresi)
		- OPPURE tra 7 e 196 SE la frequency è tra 5035 e 5980
	- frequency tra 2412 e 2484 OPPURE tra 5035 e 5980 (compresi estremi)
	- signal_dBm minore di -20 (meno venti)
	crea un'istanza di Wifi opportunamente inizializzata e la restituisce insieme a true, altrimenti restituisce una struttura "zero-value" e false

- una funzione NewWifiDaStringa(line string) Wifi,bool
	che costruisce un'istanza della struttura Wifi a partire da una stringa nel formato CSV (in cui i dati sono separati da virgole, vedere il file 'wifi.csv'), stessi vincoli della funzione definita sopra.
	Il formato è esattamente come nel file, non occorre fare controlli sul formato, ma i dati potrebbero essere non accettabili (ad es. un numero di canale non coerente con la frequenza o l'intestazione del CSV).

- un metodo String da applicare a Wifi
	che restituisca una stringa rappresentativa dei valori della struttura, nella forma:
		{ssid, channel, frequency, signal_dBm, signalW}
	Attenzione che c'è un valore in più rispetto ai dati della struct, va aggiunto un valore calcolato: la potenza del segnale in Watt. Il formato di uscita del valore signalW è quello "naturale" del float64 (formato "%v").

- una funzione ConvertiDBinWatt(signal_dBm int) float64
	che prende come parametro la potenza del segnale in dBm (decibel-milliwatts) e calcola la potenza in Watt, la formula di conversione è:
		Watt = (10^(potenza_in_dBm/10)) / 1000
		Nota: il simbolo ^ indica elevamento a potenza (10^2 è 10 alla seconda)

- una funzione PiuPotente(elenco []Wifi, banda string) int
	che restituisce l'indice della struct che rappresenta il wifi con  segnale (dBm) più potente dell'elenco passato come parametro; in funzione del parametro banda viene restituito il più potente fra i segnali a 2GHz (banda="2GHz"), fra quelli a 5GHz (banda="5GHz") o senza distinzione (banda = qualunque altro valore, compresa la stringa vuota)

- una funzione main()
	che elabora un file (CSV) il cui nome è passato come primo argomento (della linea di comando) e che stampa i dati del segnale più potente sulla base del secondo argomento: se NON presente o se diverso da "5GHz"/"2GHz", trova il più potente senza distinzione di banda; se invece il secondo argomento è presente, stampa i dati secondo la banda richiesta.
	Non occorre fare controlli sulla presenza degli argomenti sulla linea di comando e sul file, potete assumere che il programma venga sempre lanciato correttamente e che il file indicato sia presente e nel formato previsto (vedi sopra).

Esempi di esecuzione
---------------------

$ ./wifi wifi.csv
{at-wind,11,2462,-41,7.94328234724282e-08}

$ ./wifi wifi.csv 5GHz
{at-wind,108,5540,-42,6.30957344480193e-08}

$ ./wifi wifi.csv 2GHz
{at-wind,11,2462,-41,7.94328234724282e-08}

*/
package main

import (
	"fmt"
	"testing"
)

func TestNewDaStringa(t *testing.T) {
	/*
		w, ok := NewWifiDaStringa("")
		fmt.Println("creato:", w)
		if w.ssid != "" || w.channel != 0 || w.frequency != 0 || w.signal_dBm != 0 || ok {
			fmt.Println("errore di creazione da stringa vuota:")
			t.FailNow()
		}
	*/

	w, ok := NewWifiDaStringa("WOW FI - FASTWEB,6,2437,-78")
	fmt.Println("creato:", w)
	if w.ssid != "WOW FI - FASTWEB" || w.channel != 6 || w.frequency != 2437 || w.signal_dBm != -78 || !ok {
		fmt.Println("errore di creazione da stringa corretta 2GHz:")
		t.FailNow()
	}

	w, ok = NewWifiDaStringa("FASTWEB-8WDTsA,36,5180,-87")
	fmt.Println("creato:", w)
	if w.ssid != "FASTWEB-8WDTsA" || w.channel != 36 || w.frequency != 5180 || w.signal_dBm != -87 || !ok {
		fmt.Println("errore di creazione da stringa corretta 5GHz:")
		t.FailNow()
	}

	w, ok = NewWifiDaStringa("FASTWEB-8WDTsA,36,7180,87")
	fmt.Println("creato:", w)
	if w.ssid != "" || w.channel != 0 || w.frequency != 0 || w.signal_dBm != 0 || ok {
		fmt.Println("errore di creazione da stringa sbagliata:")
		t.FailNow()
	}
}

func TestNewCorrette(t *testing.T) {
	w, ok := NewWifi("pippo pluto paperino", 6, 2437, -64)
	fmt.Println("creato:", w)
	if w.ssid != "pippo pluto paperino" || w.channel != 6 || w.frequency != 2437 || w.signal_dBm != -64 || !ok {
		fmt.Println("errore di creazione con dati corretti:")
		t.FailNow()
	}
	w, ok = NewWifi("pippo pluto paperino", 6, 2437, 64)
	fmt.Println("creato:", w)
	if w.ssid != "" || w.channel != 0 || w.frequency != 0 || w.signal_dBm != 0 || ok {
		fmt.Println("errore di creazione con dati NON corretti, signal")
		t.FailNow()
	}
	w, ok = NewWifi("pippo pluto paperino", 6, 1437, -64)
	fmt.Println("creato:", w)
	if w.ssid != "" || w.channel != 0 || w.frequency != 0 || w.signal_dBm != 0 || ok {
		fmt.Println("errore di creazione con dati NON corretti, frequency")
		t.FailNow()
	}
}

var prog = "./wifi"

func TestMain(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"{at-wind,11,2462,-41,7.94328234724282e-08}\n",
		"wifi.csv")

	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"{at-wind,11,2462,-41,7.94328234724282e-08}\n",
		"wifi.csv", "2GHz")

	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"{at-wind,108,5540,-42,6.30957344480193e-08}\n",
		"wifi.csv", "5GHz")
}

func TestStruct(t *testing.T) {
	var w Wifi
	if w.ssid != "" || w.channel != 0 || w.frequency != 0 || w.signal_dBm != 0 {
		fmt.Println("errore nella struttura")
		t.FailNow()
	}
}

func TestString(t *testing.T) {
	w, _ := NewWifi("pippo pluto paperino", 6, 2437, 64)
	fmt.Println("test string:", w)
	/*
		s := w.String()
		if !(strings.Contains(s, "pippo") &&
			strings.Contains(s, "6") &&
			strings.Contains(s, "2437") &&
			strings.Contains(s, "64")) {
			t.FailNow()
		}
	*/
}

//dbm to watt
// -41 => 7.9432823472e-8
// -25 => 0.0000031622776602
func TestConversione(t *testing.T) {
	if !(intorno(ConvertiDBinWatt(-41), 7.9432823472e-8) && intorno(ConvertiDBinWatt(-25), 0.0000031622776602)) {
		fmt.Println("errata conversione")
		t.FailNow()
	}
}

/*
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
*/
