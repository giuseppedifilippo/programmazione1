package main

import (
	"fmt"
	"strconv"
)

func findcouple(n int ) int {
	s  := strconv.Itoa(n)
	if len(s) < 2 {
		return 0 
	}
	if s[0] == s[1] {
		x, _ := strconv.Atoi(s[1:])
		return 1 + findcouple(x)
	} else {
		x, _ := strconv.Atoi(s[1:])
		return 0 + findcouple(x) 
	}
}
func main() {
	num := 5122
	fmt.Println(findcouple(num))
}