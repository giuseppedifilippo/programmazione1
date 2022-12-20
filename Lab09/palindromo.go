package main

import (
	"fmt"
	"os"
	//"strings"
)

func isPalindrome(s string) bool {
	if len(s)==1 || len(s)==0 {
		return true 
	}
	//sl := strings.Split(s,"")
	if s[0]==s[len(s)-1] {
		s = s[1:len(s)-1]
		return isPalindrome(s[1:len(s)-1])
	} else {
		return false 
	}
}

func main() {
	s := os.Args[1]
	fmt.Println(isPalindrome(s))
}  