package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkInput(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil || x < -1 {
		return -2
	}
	return x
}

func leastDigit(n int) int {
	return n%10
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var mappa map[int]int
	mappa = make(map[int]int)
	myloop :
	for scanner.Scan() {
		sl := strings.Split(scanner.Text()," ")
		for _, el := range sl {
			if el =="-1" {
				break myloop
			} else if checkInput(el) != -2{
				mappa[leastDigit(checkInput(el))]++
			}
		}
	}
	/*lista := []int{}
	for k := range mappa {
		lista = append(lista, k)
	}
	sort.Ints(lista)*/
	for i:=0; i<=9;i++{
		fmt.Printf("%d - %d\n", i, mappa[i])
	}
}//passato a prima botta siummo
