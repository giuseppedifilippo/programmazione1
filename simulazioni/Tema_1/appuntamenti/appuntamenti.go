package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Appuntamento struct {
	giorno, oraInizio, oraFine int
}

func NuovoAppuntamento(gg, h1, h2 int) (Appuntamento, bool) {
	a := gg >= 0 && gg <= 366
	b := h1 >= 0 && h1 <= 24
	c := h2 >= 0 && h2 <= 24
	if a && b && c && h2 > h1 {
		return Appuntamento{giorno: gg, oraInizio: h1, oraFine: h2}, true
	}
	return Appuntamento{}, false
}

func CheckSovrapposizioneAppuntamenti(app1, app2 Appuntamento) bool {
	if app1.giorno != app2.giorno {
		return false
	} else if app1.oraInizio == app2.oraFine || app2.oraInizio == app1.oraFine {
		return false
	}
	for i := app1.oraInizio; i <= app1.oraFine; i++ {
		if i == app2.oraInizio || i == app2.oraFine {
			return true
		}
	}
	return false
}

func AggiungiAppuntamento(app Appuntamento, agenda *[]Appuntamento) bool {
	for _, el := range *agenda {
		if CheckSovrapposizioneAppuntamenti(el, app) {
			return false
		}
	}
	*agenda = append(*agenda, app)
	return true
}

func AppuntamentiGiornata(gg int, agenda []Appuntamento) []Appuntamento {
	registro := []Appuntamento{}
	for _, el := range agenda {
		if el.giorno == gg {
			registro = append(registro, el)
		}
	}
	return registro
}

func main() {
	agenda := []Appuntamento{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		temp := []int{}
		for _, el := range strings.Split(scanner.Text(), " ") {
			x, _ := strconv.Atoi(el)
			temp = append(temp, x)
		} 
		app, ok := NuovoAppuntamento(temp[0], temp[1], temp[2])
		if ok {
			AggiungiAppuntamento(app, &agenda)
		}
	}
	fmt.Println(agenda)
}//passato