package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func blok(p *int) {
	fmt.Printf("::::::::::::::\ntaglio-%02d\n::::::::::::::\n", *p)
	(*p)++
}
func main() {
	f, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println("file error")
		os.Exit(1)
	}
	defer f.Close()
	n, _ := strconv.Atoi(os.Args[1])
	scanner := bufio.NewScanner(f)
	var count, p int
	blok(&p)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		count++
		if count == n {
			blok(&p)
			count = 0 
		}
	}
}//funziona