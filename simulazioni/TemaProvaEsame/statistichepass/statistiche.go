package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func moda(serie []float64) float64 {
	mappa := map[float64]int{}
	lista := []float64{}
	max := 0
	for _, el := range serie {
		mappa[el]++
		if v, _ := mappa[el]; v > max {
			max = v 
		}
	}
	for k,v := range mappa {
		if v == max {
			lista = append(lista, k)
		}
	}
	sort.Float64s(lista)
	return lista[0]
} 

func mediana(serie []float64) float64{
	sort.Float64s(serie)
	if len(serie)%2!=0 {
		return serie[len(serie)/2]
	} else {
	return (serie[len(serie)/2] + serie[(len(serie)/2)-1] )/ 2.0
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	serie := []float64{}
	for scanner.Scan() {
		x, _ := strconv.ParseFloat(scanner.Text(), 64)
		serie = append(serie, x)
	}
	fmt.Println("moda:",moda(serie))
	fmt.Println("mediana:",mediana(serie))
}//passa tutti i test