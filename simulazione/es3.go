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
	return fmt.Sprintf("%s = (%.2f, %.2f)", p.etichetta, p.ascissa, p.ordinata)
}

func Lunghezza(tragitto []Punto) float64 {
	var totale float64
	for i:=0 ; i<len(tragitto)-1; i++ {
		totale += Distanza(tragitto[i], tragitto[i+1])
	}
	return totale
}
func NearToHalf(tragitto []Punto) (out Punto){
	half := Lunghezza(tragitto)/2.0
	var track float64
	var i int 
	for i=1; track<half; i++{
		track += Distanza(tragitto[i-1], tragitto[i])
	}
	return tragitto[i-1]
}
func main() {
	travel := NuovoTragitto()
	fmt.Printf("Lunghezza percorso: %.2f\n", Lunghezza(travel))
	fmt.Printf("Punto oltre metà: %s\n", String(NearToHalf(travel)))
}

