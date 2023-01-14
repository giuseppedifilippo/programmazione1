/*

Rotazioni
========

Si scriva un programma (il file deve chiamarsi 'rotazioni.go') per gestire rotazioni di rettangoli nel piano cartesiano, come spiegato qui di seguito.

Il programma dovrà essere dotato di:

- una struttura Point con campi x, y di tipo float64 (dichiarati in quest'ordine)

- una struttura Rectangle con campi pLL, pUR di tipo Point (dichiarati in quest'ordine),
  dove pLL è il verice in basso a sinistra (LowerLeft) e pUR quello in alto a destra (Upper Right) del rettangolo

- una funzione NewPoint(x, y float64) Point
  che data una coppia di coordinate x,y, restituisce il punto in quella posizione

- una funzione NewRectangle(p1, p2 Point) Rectangle
  che dati due vertici opposti di un rettangolo (P1 e P3 o P2 e P4, indifferentemente),
  restituisce il rettangolo corrispondente (vedi definizione di Rectangle), cioè in cui i due campi
  sono il vertice in basso a sinistra e quello il alto a destra (indicati con P1 e P3 nella figura sotto)

Nota: indipendentemente dai valori delle coordinate di due vertici opposti di un rettangolo (P1 e P3 o P2 e P4), P1 è dato da (min(x1,y1), min(y1,y2) e P3 da (max(x1,y1), max(y1,y2)

              P4 _____________ P3
                |             |
                |             |
                |_____________|
              P1               P2


- una funzione Rotate(r *Rectangle, dir byte)
che ruota il rettangolo r intorno al suo vertice pLL di 90 gradi in senso orario se dir è 'R' e in senso antiorario se dir è 'L'.
Per esempio, ruotando il rettangolo qui sopra in direzione 'L', il rettangolo diventa quello della figura qui sotto (P1 del vecchio rettangolo corrisponde a P2 del nuovo).
Per direzioni diverse da 'L' e 'R' la funzione non fa niente.
      P4 _______ P3
        |       |
        |       |
        |       |
        |       |
        |_______|
      P1         P2

- un ***metodo***  da applicare a Rectangle
  String() string
  che restituisce una riga di descrizione del rettangolo nel seguente formato: ((P1.x,P1.y) (P3.x,P3.y))
  (cioè ad es. "((2,3) (7,13))"

*/

package main

import (
	"testing"
)

func TestNewPoint(t *testing.T) {
	p := NewPoint(4.6, 199.8)
	if !intorno(p.x, 4.6) || !intorno(p.y, 199.8) {
		t.FailNow()
	}
}

func TestNewRectNormale(t *testing.T) {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(4, 6)
	r := NewRectangle(p1, p2)
	if !intorno(r.pLL.x, 0) || !intorno(r.pLL.y, 0) || !intorno(r.pUR.x, 4) || !intorno(r.pUR.y, 6) {
		t.FailNow()
	}
}

func TestNewRectEstremiScambiati(t *testing.T) {
	p1 := NewPoint(0, 6)
	p2 := NewPoint(4, 0)
	r := NewRectangle(p1, p2)
	if !intorno(r.pLL.x, 0) || !intorno(r.pLL.y, 0) || !intorno(r.pUR.x, 4) || !intorno(r.pUR.y, 6) {
		t.FailNow()
	}
}

func TestNewRectDegenere(t *testing.T) {
	p1 := NewPoint(10.2, 6)
	p2 := NewPoint(10.2, 20.4)
	r := NewRectangle(p1, p2)
	if !intorno(r.pLL.x, 10.2) || !intorno(r.pLL.y, 6) || !intorno(r.pUR.x, 10.2) || !intorno(r.pUR.y, 20.4) {
		t.FailNow()
	}
}

func TestNegativo(t *testing.T) {
	p1 := NewPoint(-10.2, 6)
	p2 := NewPoint(10.2, 20.4)
	r := NewRectangle(p1, p2)
	//fmt.Println(r)
	if !intorno(r.pLL.x, -10.2) || !intorno(r.pLL.y, 6) || !intorno(r.pUR.x, 10.2) || !intorno(r.pUR.y, 20.4) {
		t.FailNow()
	}
}

func TestRotateR(t *testing.T) {
	p1 := NewPoint(10.5, 6.9)
	p2 := NewPoint(4, 0)
	r := NewRectangle(p1, p2)
	Rotate(&r, 'R')
	//fmt.Println(r)
	if !intorno(r.pLL.x, 4) || !intorno(r.pLL.y, -6.5) || !intorno(r.pUR.x, 10.9) || !intorno(r.pUR.y, 0) {
		t.FailNow()
	}
}

func TestRotateL(t *testing.T) {
	p1 := NewPoint(10.5, 6.9)
	p2 := NewPoint(4, 0)
	r := NewRectangle(p1, p2)
	Rotate(&r, 'L')
	//fmt.Println(r)
	if !intorno(r.pLL.x, -2.9) || !intorno(r.pLL.y, 0) || !intorno(r.pUR.x, 4) || !intorno(r.pUR.y, 6.5) {
		t.FailNow()
	}
}
