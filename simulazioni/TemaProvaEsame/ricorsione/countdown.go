package main

import "fmt"

func countdown(n int )  {
	if n == 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(n)
	countdown(n-1)
}

func main() {
	n := 10
	countdown(n)
}