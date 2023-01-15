/*
Random walk
============
Scrivere un programma in Go (il file deve chiamarsi 'randomwalk.go') dotato di:

- un tipo Location che rappresenta una posizione in una griglia bidimensionale senza confini (infinita nelle quattro direzioni)

- una funzione
	func Newloc(x, y int) Location
 che crea una posizione con coordinate x,y

- una funzione
	func Move(loc Location, direction rune) (newloc Location, err error)
che "sposta" una posizione di un passo unitario corrispondente ad una delle
quattro possibili direzioni: 'U' o 'u' (up), 'D' o 'd' (down), 'R' o 'r' (right), 'L' o 'l' (left).
Restituisce la nuova posizione e nil, oppure la posizione di default (0,0) ed un errore con associato
il messaggio "unknown direction", se la direzione è sconosciuta


- un metodo String() per Location
che restituisce la stringa "(x,y)" corrispondente alla sua posizione


- una funzione
	func RandomWalk(start Location, length uint) []Location
che restituisce un cammino di lunghezza n generato in modo casuale partendo da
una posizione iniziale (che è il primo elemento del cammino)
se la lunghezza specificata è 0, il cammino contiene solo la posizione di partenza


- una funzione
	func ContainsLoop(walk []Location) (found bool)
che verifica se un cammino contiene un ciclo


- una funzione main ()
che crea una posizione di partenza ed (eventualmente) stampa un "random walk" che parte da essa.
Le coordinate della posizione di partenza e la lunghezza del cammino
sono passati, nell'ordine, come primi tre argomenti da linea di comando;
un quarto POSSIBILE argomento è l'opzione "-v" (verbose).
La funzione controlla UNICAMENTE che il numero di argomenti passati sia almeno tre,
altrimenti stampa su standard error "too few arguments".
Stampa sullo standard output, su righe separate, nell'ordine:
- lunghezza del cammino
- "loop" oppure "no loops" a seconda che il cammino contenga o meno un ciclo
- il cammino, ma solo se l'opzione "-v" è presente
*/

package main

import (
	"os"
	"strings"
	"testing"
)

//funzioni di test
func TestNewloc(t *testing.T) {
	loc1 := Newloc(-1, 1)
	loc2 := Location{-1, 1}
	if loc1 != loc2 {
		t.Errorf("%v and %v shoud be equal", loc1, loc2)
	}
	loc2 = Newloc(-1, 0)
	if loc1 == loc2 {
		t.Errorf("%v and %v are NOT equal", loc1, loc2)
	}
}

func TestMove(t *testing.T) {
	loc1 := Newloc(0, 0)
	loc2, err := Move(loc1, 'R')
	if loc2 != Newloc(1, 0) {
		t.Errorf("moved to %v, instead of %v", loc2, Newloc(1, 0))
	}
	if err != nil {
		t.Errorf("the move is valid: should return nil")
	}
	loc2, _ = Move(loc1, 'l')
	if loc2 != Newloc(-1, 0) {
		t.Errorf("moved to %v, instead of %v", loc2, Newloc(-1, 0))
	}
	loc2, _ = Move(loc1, 'U')
	if loc2 != Newloc(0, 1) {
		t.Errorf("moved to %v, instead of %v", loc2, Newloc(0, 1))
	}
	loc2, _ = Move(loc1, 'd')
	if loc2 != Newloc(0, -1) {
		t.Errorf("moved to %v, instead of %v", loc2, Newloc(0, -1))
	}
	_, err = Move(loc1, 'A')
	if err == nil {
		t.Errorf("invalid move: should NOT return nil")
	}

}

func TestString(t *testing.T) {
	loc1 := Newloc(0, -1)
	got := loc1.String()
	expected := "(0,-1)"
	if strings.Compare(got, expected) != 0 {
		t.Errorf("the representation of %v should be %s instead of %s", loc1, expected, got)
	}

}

func TestRandomWalk(t *testing.T) {
	start := Newloc(0, 1)
	w := RandomWalk(start, 0)
	if len(w) != 1 {
		t.Errorf("the length of the slice should be 1 instead of %d", len(w))
	}
	if w[0] != start {
		t.Errorf("the path at position 0 should hold %v instead of %v", start, w[0])
	}

	w = RandomWalk(start, 5)
	if len(w) != 6 {
		t.Errorf("the length of the slice should be 6 instead of %d", len(w))
	}
	if w[0] != start {
		t.Errorf("the path at position 0 should hold %v instead of %v", start, w[0])
	}

	if !checkpath(w) {
		t.Errorf("two consecutive elements in %v are not at distance one", w)
	}
}

func TestRandomness(t *testing.T) {
	start := Newloc(0, 1)
	w1 := RandomWalk(start, 7)
	w2 := RandomWalk(start, 7)
	i := 1
	for ; i < len(w1); i++ {
		if w1[i] != w2[i] {
			break
		}
	}
	if i == len(w1) {
		t.Errorf("got the same paths: are you using a random generator or not?")
	}
	//fmt.Println("1st:", w1, "2nd:", w2) //uncomment for debugging

}

func TestContainsLoop(t *testing.T) {
	w := []Location{Newloc(0, 0), Newloc(0, -1), Newloc(-1, -1), Newloc(-1, 0), Newloc(0, 0), Newloc(0, 1)}
	if !checkpath(w) {
		os.Exit(1) //controllo "interno"
	}
	if !ContainsLoop(w) {
		t.Errorf("the path %v contains a loop", w)
	}
	if ContainsLoop(w[:4]) {
		t.Errorf("the path %v contains NO loops", w[:4])
	}
}

//Funzione tolta
/*
func TestContains(t *testing.T) {
	w := []Location{Newloc(0, 0), Newloc(0, -1), Newloc(-1, -1), Newloc(-1, 0), Newloc(0, 0), Newloc(0, 1)}
	w2 := []Location{Newloc(-1, -1), Newloc(-1, 0), Newloc(0, 0)}

	if !checkpath(w) {
		os.Exit(1) //controllo "interno"
	}
	if !Contains(w, w2) {
		t.Errorf("the path %v contains %v", w, w2)
	}
	w2[2] = Newloc(-2, 0)
	if Contains(w, w2) {
		t.Errorf("the path %v does NOT contain %v", w, w2)
	}
	w2 = w2[1:2]
	if !Contains(w, w2) {
		t.Errorf("the path %v contains %v", w, w2)
	}
}
*/

//esegue un test del main
func ExampleMain() {
	os.Args = []string{"randomwalk", "0", "0", "100"}
	main()
	// Output:
	// 100
	// loop
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n

}
func checkdistance(l1, l2 Location) bool {
	return abs(l1.x-l2.x)+abs(l1.y-l2.y) == 1
}

func checkpath(p []Location) bool {
	if len(p) > 1 {
		for i := 1; i < len(p); i++ {
			if !checkdistance(p[i], p[i-1]) {
				return false
			}
		}
	}
	return true
}
