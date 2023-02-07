package main

import "fmt"

func somma(start , end int) int {
	if start == end {
		return end 
	}
	return start + somma(start + 1 , end )
}

func main() {
	start ,end := 2 , 6 
	fmt.Println(somma(start , end))
}