/**
N.B. leggere il file README.txt per istruzioni di compilazione, test e consegna

=== Cerchi ===

Scrivere un programma 'cerchi.go' dotato di:

- una struttura 'Cerchio' con campi:
	nome string
	x,y,raggio float64
		(dove x e y sono le coordinate del centro)
  dichiarati in quest'ordine

- una funzione 'newCircle(descr string) Cerchio'
	che data una stringa che descrive un cerchio
	(nome, coordinate x e y del centro, raggio , in quest'ordine e separati da spazi)
	restituisce il cerchio corrispondente

- una funzione 'distanzaPunti(x1,y1,x2,y2 float64) float64
	che restituisce la distanza tra i punti di coordinate (x1,y1) e (x2,y2)

- una funzione 'tocca(cerc1, cerc2 Cerchio) bool'
	che restituisce true se i due cerchi sono tangenti
	(internamente o esternamente); false altrimenti.
	Trattandosi di valori float, consideriamo una distanza
	trascurabile se è inferiore a 10^-6 (1e-6)

- una funzione maggiore(cerc1, cerc2 Cerchio) bool
	che restituisce true se cerc1 è più grande di cerc2;
	false altrimenti.

- una funzione main()
	che legge da standard input una sequenza (terminata da ctrl D)
	di righe che descrivono cerchi, del tipo:

uno 10.3 12.7 45.8
due 1.3 2.9 5.8
pippo 7.3 22.5 6.6

	ciascuna contenente nome, coordinate del centro e raggio di un
	cerchio, in quest'ordine.

Il programma crea per ciascuna riga il cerchio corrispondente, lo stampa.
Inoltre stampa se il cerchio, rispetto a quello precedente, è tangente o no e maggiore o no.
Il primo confronto è fatto rispetto al cerchio "zero" ("", 0, 0, 0).
Non sono richiesti controlli sulla correttezza dei dati in input, potete assumere che siano sempre nell'ordine e del tipo previsto.




Esempio
-------
(vengono marcate con '>' le righe digitate da tastiera,
le altre sono l'output del programma)

>uno 0.5 0 2.5
{uno 0.5 0 2.5} non tangente, maggiore
>due 0 0 3
{due 0 0 3} tangente, maggiore
>tre 10.2 -8.4 1.5
{tre 10.2 -8.4 1.5} non tangente, minore o uguale

*/

package main

import (
	"testing"
	//"fmt"
	"math"
)

var prog = "./cerchi"

const E = 1e-6

func TestDistanzaZero(t *testing.T) {
	if distanzaPunti(.0, .0, .0, .0) >= E {
		t.Fail()
	}
}

func TestDistanzaQuasiZero(t *testing.T) {
	if distanzaPunti(E, .0, .0, .0) > E {
		t.Fail()
	}
}

func TestDistanzaQuadro(t *testing.T) {
	d := distanzaPunti(1.0, 0.0, 0.0, 1.0)
	if d <= 1.4142136-E || d >= 1.4142136+E {
		t.Fail()
	}
}

func TestNewCircleSimple(t *testing.T) {
	c := newCircle("t 1.0 11.3 2.6")

	if math.Abs(c.raggio-2.6) >= E {
		t.Fail()
	}
}

func TestNewCircleSimpleConE(t *testing.T) {
	c := newCircle("t 1.0 1e-6 2.6")

	if math.Abs(c.raggio-2.6) >= E {
		t.Fail()
	}
}

func TestNewCircleDistanzaConE(t *testing.T) {
	c := newCircle("t 5.0 1e-6 2.6")

	if distanzaPunti(5.0, 1e-6, c.x, c.y) >= E {
		t.Fail()
	}
}

func TestNewCircleDistanzaGrande(t *testing.T) {
	c := newCircle("t -5.0 -45 2.6")

	if !(distanzaPunti(6.0, 1e-6, c.x, c.y) >= 45) {
		t.Fail()
	}
}

func TestNewCircleNonTocca(t *testing.T) {
	a := newCircle("alpha 1.0 0.0 2.6")
	//b := newCircle("beta 0,0000005 0.0 2.6")
	b := newCircle("beta 0.0000005 0.0 2.6")

	if tocca(a, b) {
		t.Fail()
	}
}

func TestNewCircleToccaEsterno(t *testing.T) {
	a := newCircle("alpha 0.0 0.0 2.6")
	b := newCircle("beta 0.0 3.6 1.0")

	if !tocca(a, b) {
		t.Fail()
	}
}

func TestNewCircleToccaInterno(t *testing.T) {
	a := newCircle("alpha 0.0 0.0 5.6")
	b := newCircle("beta 0.0 3.6 2.0")

	if !tocca(a, b) {
		t.Fail()
	}
}

func TestComeDaEsempio(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"uno 0.5 0 2.5\ndue 0 0 3\ntre 10.2 -8.4 1.5",
		"{uno 0.5 0 2.5} non tangente, maggiore\n{due 0 0 3} tangente, maggiore\n{tre 10.2 -8.4 1.5} non tangente, minore o uguale\n")
}

func TestEsteso(t *testing.T) {
	LanciaGenerica(t,
		prog,
		"uno 0.5 0 2.5\ndue 0 0 3\ntre 10.2 -8.4 1.5\nquattro 10.3 12.7 45.8\ncinque 1.3 2.9 5.8\nsei 7.3 22.5 6.6\n",
		"{uno 0.5 0 2.5} non tangente, maggiore\n{due 0 0 3} tangente, maggiore\n{tre 10.2 -8.4 1.5} non tangente, minore o uguale\n{quattro 10.3 12.7 45.8} non tangente, maggiore\n{cinque 1.3 2.9 5.8} non tangente, minore o uguale\n{sei 7.3 22.5 6.6} non tangente, maggiore\n")
}
