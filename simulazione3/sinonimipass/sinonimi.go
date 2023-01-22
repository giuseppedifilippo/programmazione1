package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
type Vocabolario map[string][]string 
func ParseSin(filename string) Vocabolario {
	//var vocabolario Vocabolario
	vocabolario := make(Vocabolario)
	sin, err1 := os.Open(filename)
	if err1 != nil {
		fmt.Println("file 1 not found")
		os.Exit(2)
	}
	defer sin.Close()
	scanner := bufio.NewScanner(sin)
	for scanner.Scan() {
		key , sinonimi,_ := strings.Cut(scanner.Text(), ":")
		sl := strings.Split(strings.Trim(sinonimi," "),", ")
		vocabolario[key]= sl
	}
	return vocabolario
}


func main() {
	if len(os.Args) < 3 {
		fmt.Println("2 file names, please")
		os.Exit(1)
	}
	text,err2 := os.Open(os.Args[2])
	if err2 != nil {
		fmt.Println("file 2 not found")
		os.Exit(3)
	}
	defer text.Close()
	vocabolario := ParseSin(os.Args[1])
	scanner := bufio.NewScanner(text)
	for scanner.Scan() {
		var index int 
		riga := scanner.Text()
		sl:= strings.Split(riga, " ")
		for _, el := range sl {
			index += len(el)
			if val, ok := vocabolario[el]; ok {
				riga = fmt.Sprintf("%s%s%s", riga[:index], val,riga[index:])
			}
			index++
		}
		fmt.Println(riga)
	}
}