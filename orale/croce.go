// dato da n dispari dalinea di comando stampare una croce lunga n asterischi
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	n, _ := strconv.Atoi(os.Args[1])
	for i:= 0 ; i < n ; i++ {
		if i == n /2 {
			fmt.Println(strings.Repeat("*", n ))
			continue
		}
		for j := 0 ; j < n ; j++ {
			if j == n /2 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
