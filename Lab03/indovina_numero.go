package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n, i int
	const MAX = 200
	rand.Seed(MAX)
	m := rand.Intn(MAX)
	for i = 0; i < MAX/2; i++ {
		fmt.Scan(&n)
		if n < 1 || n > 200 {
			fmt.Println("fuori intervallo!")
		} else if n == m {
			fmt.Println("hai vinto!, in", i+1, "tentativi")
			break
		}
	}
	if i == MAX/2 {
		fmt.Println("hai perso, il numero era", m)
	}
}
