package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func reverse(s string) string {
	temp := []byte(s)
	for i:=0; i<len(temp)/2;i++ {
		temp[i], temp[len(temp)-1-i] = temp[len(temp)-1-i] , temp[i]
	}
	return string(temp)
}

func main() {
scanner := bufio.NewScanner(os.Stdin)
myloop:
for scanner.Scan() {
	sl := strings.Split(strings.Trim(scanner.Text(),"\t"), " ")
	for _, el := range sl {
		if el == "stop" {
			break myloop
		}
		if len(el)%2==0 && len(el) != 0{
			fmt.Println(reverse(el))
		}
	}
}
}//passato