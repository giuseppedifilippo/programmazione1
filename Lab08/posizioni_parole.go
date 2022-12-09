package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//rimuovere duplicati
func rmDup(x []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, el := range x {
		if _ , value := keys[el]; !value {
			keys[el] = true
			list = append(list ,el)
		}
	}	
	return list 
}

func main() {
	mappa := make(map[string][]int)
	sl := make([]string, 0)
	scanner:= bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), " ")
		sl = append(sl, a...)
	}
	check := rmDup(sl)
	for _, key := range check {
		mappa[key] = []int{}
		for i, el := range sl {
			if key == el {
				mappa[key] = append(mappa[key], i )
			}
		}
	}
	fmt.Println(mappa)
}