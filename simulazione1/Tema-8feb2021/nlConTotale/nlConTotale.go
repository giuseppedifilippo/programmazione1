package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getnl(f string) int{
	temp, _ := os.ReadFile(f)
	return strings.Count(string(temp), "\n") 
}
func main() {
	in := os.Args[1]
	nl := getnl(in)
	f, err := os.Open(in)
	if err != nil {
		fmt.Println("errore d input")
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	i := 1 
	for scanner.Scan() {
		fmt.Printf("%d/%d:\t%s\n", i , nl, scanner.Text())
		i++
	}
}//funziona come nell esempio ma i file di test non funzionano proprio pace 