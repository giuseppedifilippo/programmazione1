package main

import (
	"fmt"
	"os"
	"strings"
)

func isPalindrome(s string) bool {
	if len(s)==1 || len(s)==0 {
		return true 
	}
	sl := strings.Split(s,"")
	if sl[0]==sl[len(sl)-1] {
		sl = sl[1:len(s)-1]
		return isPalindrome(strings.Join(sl, ""))
	} else {
		return false 
	}
}

func main() {
	s := os.Args[1]
	fmt.Println(isPalindrome(s))
}