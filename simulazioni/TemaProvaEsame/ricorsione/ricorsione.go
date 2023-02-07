package main

import "fmt"

func recurs(c byte, s string) int {
	
	if len(s) == 0 {
		return -1
	}
	if c == s[0] {
		return 1
	}
	return 1 + recurs(c, s[1:]) 
}
func main() {
	stringa := "caopalle"
	c := []byte("z")[0]
	fmt.Println(recurs(c,stringa))
}