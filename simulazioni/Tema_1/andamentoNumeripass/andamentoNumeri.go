package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func sw(a,b int) {
	switch {
	case a < b : fmt.Print("+")
	case a > b : fmt.Print("-")
	case a == b : fmt.Print("=")
	}
}
func main() {
	sl := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n , _ := strconv.Atoi(scanner.Text())
		sl = append(sl, n)
	}
	for i:= 0; i< len(sl)-1; i++ {
		sw(sl[i], sl[i+1])
	}
	fmt.Println()
}//passa i test 