package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//rimuovere duplicati
func rmDup(x []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, el := range x {
		if _ , value := keys[el]; !value {
			keys[el] = true
			list = append(list ,el)
		}
	}	
	return list 
}

func main() {
	mappa := make(map[string][]int)
	sl := make([]string, 0)
	scanner:= bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		a := strings.Split(strings.Trim(scanner.Text()," "), " ")// questo trim elimina eventuali spazi a inizio o fine riga 
		for _, el := range a { // la funzione di questo ciclo è di rimuovere eventuali spazi aggiuntivi
			el = strings.Trim(el, " ") // tra una parola e l altra
		}
		sl = append(sl, a...)
	}
	check := rmDup(sl)
	for _, key := range check {
		mappa[key] = []int{}
		for i, el := range sl {
			if key == el {
				mappa[key] = append(mappa[key], i )
			}
		}
	}
	fmt.Println(mappa)
}// il programma funziona ma ha un bug per cui se alla fine della riga inserisco uno spazio prima di inseire 
// ENTER il programma conta anche lo spazio come parola(problema dovuto al funzionamento di strings.Split)
// risolto
// p.s so che mancano gli accenti nei commenti ma sto usando una tastiera inglese 