/*
Cisterna
========


Scrivere un programma 'cisterna.go' dotato di:

- una struttura 'Cisterna' con 3 campi float64 per le tre dimensioni (in cm) (NOTA SERVIRA' SPECIFICARE NOME E ORDINE DEI CAMPI?) e un ulteriore campo float64 che indica il livello (in cm) a cui arriva l'acqua nella cisterna

- una funzione

    variazione(cisterna *Cisterna, lt int) error

  che, se lt è positivo e c'è abbastanza spazio nella cisterna, versa lt litri (***) nella cisterna e restituisce errore nil, altrimenti non versa niente nella cisterna e restituisce un errore (vedi esempio) segnalando la capienza in litri ancora disponibile; se invece lt è negativo, se ci sono abbastanza litri, prende lt litri dalla cisterna e restituisce errore nil, altrimenti non prende niente dalla cisterna e restituisce un errore (vedi esempio) segnalando la disponibilità (in litri).

- una funzione
	func contenuto(cisterna Cisterna) (litri int)
  che restituisce il numero di litri contenuti nella cisterna

- un metodo String() per Cisterna che restituisce una descrizione della cisterna (vedi esempio sotto)

- una funzione main() che legge da linea di comando tre numeri (float64) che rappresantano larghezza, lunghezza e altezza (in quest'ordine) di una cisterna a base rettangolare.
Se i dati sono tutti > 0, crea una cisterna vuota (livello = 0) con quelle dimensioni e la stampa (vedi es.); altrimenti segnala l'errore "tutti gli argomenti devono essere >0" (se anche uno solo è minore o uguale a 0) o "tutti gli argomenti devono essere numerici" (se anche uno solo non è numerico) e termina.
Poi legge da standard input una sequenza di interi, che rappresentano litri, terminata da 9999: quando il numero letto è positivo, viene messo nella cisterna un numero di litri corrispondenti, viceversa quando il numero è negativo, i litri vengono tolti, sempre che ciò sia possibile, altrimenti viene segnalato un errore (senza però terminare).
Dopo ogni variazione effettiva,  il programma stampa lo stato della cisterna (vedi es.).
Se invece la variazione della cisterna non va a buon fine, il programma non ristampa il suo stato,

*** Nota: un litro corrisponde a un decimetro cubo

Suggerimenti:
1. per creare un errore, utilizzare la funzione New del package 'errors' (vedere la documentazione relativa).

Esempio
=======

(nota: per distinguere input da output, l'input è preceduto da >>>, che però non andrà digitato)

$ ./cisterna 50 100 150
cisterna 50.00 cm x 100.00 cm x 150.00 cm
livello attuale: 0.00 cm, litri 0
>>>1000
puoi mettere max 750 litri
>>>500
cisterna 50.00 cm x 100.00 cm x 150.00 cm
livello attuale: 100.00 cm, litri 500
>>>-600
puoi prendere max 500 litri
>>>-300
cisterna 50.00 cm x 100.00 cm x 150.00 cm
livello attuale: 40.00 cm, litri 200
>>>9999
*/

package main

import (

	//"io/ioutil"

	"fmt"
	"testing"
	//"log"
)

var prog = "./cisterna"

func TestNoArgs(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "fornire larghezza, lunghezza, altezza\n")
}
func TestUnArg(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "fornire larghezza, lunghezza, altezza\n", "qualcosa")
}

func Test9999(t *testing.T) {
	lanciaGenerica(t, prog, "9999", "cisterna 10.00 cm x 10.00 cm x 10.00 cm\nlivello attuale: 0.00 cm, litri 0\n","10","10","10")
}

func TestRiempi(t *testing.T) {
	var cisterna = Cisterna{100, 200, 100, 0}
	fmt.Println(cisterna)
	variazione(&cisterna, 200)
	fmt.Println(cisterna)

	if contenuto(cisterna) != 200 {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}

func TestRiempiFloat(t *testing.T) {
	var cisterna = Cisterna{120.5, 200.7, 10.9, 7.8}
	fmt.Println(cisterna)
	variazione(&cisterna, 10)
	fmt.Println(cisterna)

	if contenuto(cisterna) != 198 {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}

func TestRiempiTroppo(t *testing.T) {
	var cisterna = Cisterna{10, 10, 110, 100}
	fmt.Println(cisterna)
	variazione(&cisterna, 200)
	fmt.Println(cisterna)

	if contenuto(cisterna) != 10 {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}

func TestSvuota(t *testing.T) {
	var cisterna = Cisterna{120, 80, 100, 50}
	fmt.Println(cisterna)
	variazione(&cisterna, -50)
	fmt.Println(cisterna)

	if contenuto(cisterna) != 430 {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}

func TestSvuotaTroppo(t *testing.T) {
	var cisterna = Cisterna{120, 80, 100, 50}
	fmt.Println(cisterna)
	variazione(&cisterna, -500)
	fmt.Println(cisterna)

	if contenuto(cisterna) != 480 {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}
}

func TestCubica(t *testing.T) {
	lanciaGenericaConFileOutAtteso(t, prog, "40\n34\n-78\n-23\n51\n500\n500\n398\n-1000\n9999\n", "cubica.expected", "100", "100", "100")
}

func TestArgAZero(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "tutti gli argomenti devono essere >0\n", "0", "0", "0")
}

func TestUnArgAZero(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "tutti gli argomenti devono essere >0\n", "50", "0", "120")
}

func TestArgNonNum(t *testing.T) {
	lanciaGenerica(t, prog, "nil", "tutti gli argomenti devono essere numerici\n", "50", "pippo", "120")
}

func TestString(t *testing.T) {
	var cisterna = Cisterna{120, 80, 100, 50}
	fmt.Println(cisterna)
	
	if cisterna.String() != "cisterna 120.00 cm x 80.00 cm x 100.00 cm\nlivello attuale: 50.00 cm, litri 480" {
		t.Fail()
	}
}
