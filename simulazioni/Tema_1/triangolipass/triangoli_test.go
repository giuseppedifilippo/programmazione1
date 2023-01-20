/**

Esercizio "triangoli"
=====================

Scrivere un programma Go (il file deve chiamarsi 'triangoli.go') dotato di:

- una struttura Punto(x,y), dove x e y sono float64

- una struttura Triangolo(P1,P2,P3), dove P1, P2 e P3 sono di tipo Punto

- una funzione

	 lunghezzeLati(A, B, C Punto) [3]float64

  che restituisce una array di tre elementi corrispondenti, nell'ordine, alle distanze
 A-B, B-C e C-A

- una funzione

	newTriangolo(A, B, C Punto) (Triangolo, error)

  che restituisce il triangolo A, B, C se i punti A, B, C determinano
  effettivamente un triangolo (*), altrimenti restituisce un errore

- una funzione

	tipoTriangolo(triangolo Triangolo) int

  che, passato un triangolo come parametro, restituisce il numero di
  lati di uguale (**) lunghezza, cioè:

	- 0 se il triangolo è scaleno

	- 2 se il triangolo è isoscele

	- 3 se il triangolo è equilatero

- una funzione

	main()

  che legge le coordinate x,y  (float64) di tre punti A, B, C (un punto per riga) da
  standard input ed emette su standard output le distanze A-B, B-C e C-A.
  Inoltre, se i punti non formano un triangolo, emette il messaggio di
  errore "non è un triangolo", altrimenti emette "triangolo" e i suoi dati
  e "tipo" e il suo tipo (n. di lati uguali).
  Per il formato dell'output, vedi gli esempi sotto.

  (*) Disuguaglianza triangolare: In un triangolo la misura di ciascun lato è sempre
  strettamente minore della somma delle misure degli altri due lati.

  (**) Avvertenza: per verificare se due valori double a,b sono "uguali" si suggerisce
  di utilizzare la seguente approssimazione:
  a e b sono uguali se e solo se |a - b| < precisione
  dove precisione sia fissata a 1e-6 (notazione Go per "10 alla -6")

Esempi
------

1) dato come input:

0 -2
-2 0
0 0

l'output è:
[2.8284271247461903 2 2]
triangolo {{0 2} {2 0} {0 0}}
tipo 2

2) dato come input:

1 1
2 2
3 3

l'output è:
[1.4142135623730951 1.4142135623730951 2.8284271247461903]
non è un triangolo

*/

package main

import (
	"fmt"

	//"log"
	"testing"
)

var programma = "./triangoli"

func TestMain(t *testing.T) {
	LanciaGenerica(t, programma, "1 1\n2 2\n3 3", "[1.4142135623730951 1.4142135623730951 2.8284271247461903]\nnon è un triangolo\n")                 //, "non triangolo")
	LanciaGenerica(t, programma, "0 -2\n-2 0\n0 0", "[2.8284271247461903 2 2]\ntriangolo {{0 -2} {-2 0} {0 0}}\ntipo 2\n")                            //, "isoscele")
	LanciaGenerica(t, programma, "0 0\n10 0\n0 10", "[10 14.142135623730951 10]\ntriangolo {{0 0} {10 0} {0 10}}\ntipo 2\n")                          //, "isoscele")
	LanciaGenerica(t, programma, "0 0\n100 0\n0 10", "[100 100.4987562112089 10]\ntriangolo {{0 0} {100 0} {0 10}}\ntipo 0\n")                        //, "scaleno")
	LanciaGenerica(t, programma, "1 1\n100 1\n1 10", "[99 99.40824915468535 9]\ntriangolo {{1 1} {100 1} {1 10}}\ntipo 0\n")                          //, "scaleno")
	LanciaGenerica(t, programma, "-1 0\n0 1.73205\n1 0", "[1.9999993006248777 1.9999993006248777 2]\ntriangolo {{-1 0} {0 1.73205} {1 0}}\ntipo 3\n") //, "equilatero")
}

func TestFunzioniTriOK(t *testing.T) {
	fmt.Println("------chiama funzioni-----------------------------")
	var P1, P2, P3 Punto

	P2.x = 5
	P2.y = 0
	P3.x = 0
	P3.y = 10

	triangolo, err := newTriangolo(P1, P2, P3)

	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	}

	fmt.Println(triangolo)

	lun := lunghezzeLati(P1, P2, P3)
	fmt.Println(lun)

	if lun[0] != 5 {
		t.Fail()
	}
}

func TestFunzioniNONTri(t *testing.T) {
	fmt.Println("------chiama funzioni, deve dare errore-----------------------------")
	var P1, P2, P3 Punto

	P1.y = 10
	P2.x = 50
	P2.y = 10
	P3.x = 100
	P3.y = 10

	triangolo, err := newTriangolo(P1, P2, P3)

	if err == nil {
		t.Errorf("Fallito: %s\n", err)
	}

	fmt.Println(triangolo)
	fmt.Println(err)
}

func TestFunzioneTipo0(t *testing.T) {
	fmt.Println("------chiama funzioni-----------------------------")
	var P1, P2, P3 Punto

	P1.x = -3
	P1.y = -4

	P2.x = 5
	P2.y = 2

	P3.x = 3
	P3.y = -5

	triangolo, err := newTriangolo(P1, P2, P3)
	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	} else {
		fmt.Println(triangolo)
		if tipoTriangolo(triangolo) != 0 {
			t.Fail()
		}
	}
}

func TestFunzioneTipo2(t *testing.T) {
	fmt.Println("------chiama funzioni-----------------------------")
	var P1, P2, P3 Punto

	P1.y = 40

	P2.x = 5
	P2.y = 40

	P3.x = 0
	P3.y = 45

	triangolo, err := newTriangolo(P1, P2, P3)
	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	} else {
		fmt.Println(triangolo)
		if tipoTriangolo(triangolo) != 2 {
			t.Fail()
		}
	}
}

func TestFunzioneTipo3(t *testing.T) {
	fmt.Println("------chiama funzioni-----------------------------")
	var P1, P2, P3 Punto

	P1.x = 5
	P1.y = 0

	P2.x = 9
	P2.y = 0

	P3.x = 7
	P3.y = 3.464101615

	triangolo, err := newTriangolo(P1, P2, P3)
	if err != nil {
		t.Errorf("Fallito: %s\n", err)
	} else {
		fmt.Println(triangolo)

		if tipoTriangolo(triangolo) != 3 {
			t.Fail()
		}
	}
}
