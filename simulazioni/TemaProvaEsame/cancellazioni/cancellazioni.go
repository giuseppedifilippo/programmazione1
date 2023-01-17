package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cancella(lista []string) []string {
	out := []string{}
	for i:= 0; i<len(lista); i++ {
		x, err := strconv.Atoi(lista[i])
		if err == nil {
			i+=x
			continue
		}
		out = append(out, lista[i])
	}
	return out
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Fornire 1 nome di file")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File non accessibile")
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lista := []string{}
	for scanner.Scan() {
		lista = append(lista, strings.Split(scanner.Text(), " ")...)
		
	}
	fmt.Println(lista)
	fmt.Println(cancella(lista))
}