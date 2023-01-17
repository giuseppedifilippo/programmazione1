package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi(os.Args[1])
	j , _ := strconv.Atoi(os.Args[2])
	f, _ := os.Open(os.Args[3])
	defer f.Close()
	x := i
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(scanner.Bytes())  < i {
			fmt.Println()
			continue
		}
		for x= i-1 ;  x <len(scanner.Bytes()) ; x++ {
			if x==j {
				break
			}
			fmt.Print(string(scanner.Bytes()[x]))
		}
		fmt.Println()
		x = i
	}
}//funziona
