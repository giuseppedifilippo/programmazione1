package main

import (
	"fmt"
	"os"
	"strconv"
)

func DrawPoint(c byte, k int) string {
	var out string
	for i := 0; i < k; i++ {
		out += " "
	}
	return out + string(c)
}

func DrawSegment(c byte, k, l int) string {
	out := DrawPoint(c, k)
	for i := 1; i < l; i++ {
		out += string(c)
	}
	return out
}

func main() {
	in := os.Args[1:]
	c := []byte(in[2])[0]
	l, _ := strconv.Atoi(in[0])
	n, _ := strconv.Atoi(in[1])
	if l > 0 && n > 0 {
		for i := 0; i < n; i++ {
			fmt.Println(DrawSegment(c, i*(l-1), l))
			for j := 0; j < l-1; j++ {
				fmt.Println(DrawPoint(c, i*(l-1)+(l-1)))
			}
		}
	}
}//passato
