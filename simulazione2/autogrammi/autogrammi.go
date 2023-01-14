package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalcQuanteMinMax(frase string) (quante, min, max int) {
	frase = strings.ReplaceAll(strings.ReplaceAll(frase, ":",""), ",", "")
	sl := strings.Split(frase, " ")
	min, max = len(sl[0]), len(sl[0])
	for _, el := range sl {
		_, err := strconv.Atoi(el)
		if err != nil {
			quante++
			if len(el) > max {
				max = len(el) 
			} else if len(el) < min {
				min = len(el)
	}
		}
	}
	return quante, min, max
}

func TrovaNumDopoKeyword(frase, keyword string) (num int) {
	_, frase,_ = strings.Cut(frase, keyword+": ")
	frase, _ , _ = strings.Cut(frase, " ")
	num, _ = strconv.Atoi(strings.Trim(frase, ","))
	return num
}

func Autogramma(frase string) bool {
	quante , min, max := CalcQuanteMinMax(frase) 
	cq, cmin, cmax := TrovaNumDopoKeyword(frase, "contiene"), TrovaNumDopoKeyword(frase, "minima"), TrovaNumDopoKeyword(frase, "massima")
	if quante == cq && min == cmin && max == cmax {
		return true
	} else {
		return false 
	}
}

func StampaAnalisiAutogramma(frase string) {
	quante , min, max := CalcQuanteMinMax(frase) 
	cq, cmin, cmax := TrovaNumDopoKeyword(frase, "contiene"), TrovaNumDopoKeyword(frase, "minima"), TrovaNumDopoKeyword(frase, "massima")
	fmt.Printf("minimo dichiarato %d contro minimo verificato %d\nmassimo dichiarato %d contro massimo verificato %d\nnumero parole dichiarato %d contro numero parole verificato %d\n",cmin, min,cmax,max,cq,quante)
	if Autogramma(frase) {
		fmt.Println("onesto")
	} else {
		fmt.Println("bugiardo")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("file name?")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("file not found")
		os.Exit(2)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f) 
	for scanner.Scan() {
		if len(scanner.Text()) != 0 {
		fmt. Printf("=== %s\n", scanner.Text())
		StampaAnalisiAutogramma(scanner.Text())
	}
}
}//il programma funziona alla perfezione tranne che alla fine della stampa un newline in piu