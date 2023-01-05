package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)
type Punto struct {
	etichetta string 
	ascissa, ordinata float64
}

func NuovoTragitto() ( tragitto []Punto) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var p Punto
		sl:= strings.Split(scanner.Text(), ";")
		p.etichetta = sl[0]
		p.ascissa, _ = strconv.ParseFloat(sl[1], 64) 
		p.ordinata , _ = strconv.ParseFloat(sl[2], 64)
		tragitto = append(tragitto, p)
	}
	return tragitto 
}

func Distanza(p1, p2 Punto) float64 {
	Dx := math.Pow(p1.ascissa - p2.ascissa, 2 )
	Dy := math.Pow(p1.ordinata - p2.ordinata, 2 )
	return math.Sqrt(Dx + Dy)
}

func String(p Punto) string {
	return fmt.Sprintf("%s = (%f, %f)", p.etichetta, p.ascissa, p.ordinata)
}

func Lunghezza(tragitto []Punto) float64 {
	var totale float64
	for i:=0 ; i<len(tragitto)-1; i++ {
		totale += Distanza(tragitto[i], tragitto[i+1])
	}
	return totale
}

func main() {
	travel := NuovoTragitto()
	fmt.Println(Lunghezza(travel))
}

