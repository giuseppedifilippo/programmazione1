package main

import (
	"fmt"
	"os"
	"strconv"
)

func NestedSquares(n int) string {
	out := ""
	for i := 0; i < n; i++ {
		out += "*"
	}
	out += "\n"
	for i := 0; i < n-2; i++ {
		for j := 0; j < n; j++ {
			if j == 0 || j == n-1 {
				out += "*"
			} else if i>= n/2 && i <= n/2 +(n-4) && j>= n/2 && j <= n/2 +(n-4)   {
				out += "%"
			} else {
				out += " "
			}

		}
		out += "\n"
	}
	for i := 0; i < n; i++ {
		out += "*"
	}
	return out
}

func main() {
	n, _ := strconv.Atoi(os.Args[1])
	fmt.Println(NestedSquares(n))
}
