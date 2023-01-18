package main

import "fmt"

func Abbondante(n int) bool {
	if n <= 0 {
		return false
	}
	abb := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			abb += i
		}
	}
	if n >= abb {
		return false
	} else {
		return true
	}
}

func main() {
	var n, count int
	fmt.Scanf("%d\n", &n)
	for i := 0; count < n; i++ {
		if Abbondante(i) {
			fmt.Println(i)
			count++
		}
	}
}//passato alla prima botta siuuuuuuuu