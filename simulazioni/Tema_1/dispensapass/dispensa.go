package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func AggiornaDispensa(dispensa map[string]float64, filename string) bool {
	f, err := os.Open(filename)
	if err != nil {
		return false
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sl := strings.Split(scanner.Text(), ",")
		x, _ := strconv.ParseFloat(sl[2], 64)
		if sl[0] == "approv" {
			dispensa[sl[1]] += x
		} else {
			dispensa[sl[1]] -= x
		}
		if dispensa[sl[1]] < 0 {
			return false
		}
	}
	return true
}

func Rimanenza(dispensa map[string]float64, alimento string) float64 {
	return dispensa[alimento]
}

func main() {
	file := os.Args[1]
	dispensa := make(map[string]float64)
	ok := AggiornaDispensa(dispensa, file)
	if ok && len(os.Args) < 3 {
		fmt.Println("file corretto")
	} else if !ok && len(os.Args) < 3 {
		fmt.Println("file non corretto")
	} else {
		for _, el := range os.Args[2:] {
			fmt.Printf("%s %g\n", el , Rimanenza(dispensa, el))
		}
	}
	
}//passato