package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
)

func DrawPoint(c byte, k int) string {
	return strings.Repeat(" ", k) + string(c)
}

func DrawSegment(c byte, k, l int) string {
	return strings.Repeat(" ", k) + strings.Repeat(string(c), l)
}

func main() {
	// fmt.Println(DrawSegment('k', 5, 3))

	l,_ := strconv.Atoi(os.Args[1])
	n,_ := strconv.Atoi(os.Args[2])
	c := os.Args[3][0]
	_= n

	if n > 0 && l > 0 {
		for i := 0; i < n; i++ {
			fmt.Println(DrawSegment(c, i*(l-1), l))
			for k := 0; k < l-1; k++ {
				fmt.Println(DrawPoint(c, (l-1)*i+l-1))
			}
		}
	}
}