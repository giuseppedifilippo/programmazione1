package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	sl := os.Args[1:]
	for i:=0; i<len(sl)-1; i+=2 {
		x,_ := strconv.Atoi(sl[i+1])
		for j:=0; j<x; j++{
			fmt.Print(sl[i])
		}
	}
	fmt.Println()
}