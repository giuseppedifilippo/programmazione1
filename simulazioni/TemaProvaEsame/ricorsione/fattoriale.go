package main

import "fmt"

func fact(n int) int {
	if n == 0{
		return 1 
	}
	return n * fact(n-1)
}

func main() {
	num := 10 
	fmt.Println(fact(num))
}