package main

import (
	"fmt"
	"os"
)
type condizione struct {
	timestamp string 
	temp float64
	hum float64
}
func parseFile(filename string) (sl []condizione) {
	f, err := os.Open(filename) 
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "errore di apertura del file %s\n", filename)
	}
	scanner := bufio.
}