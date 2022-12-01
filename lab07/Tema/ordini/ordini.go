package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Conto(s string, x *float64 ) {
	c := 1.0
	sl:= strings.Split(s, "#")
	for _, str := range sl {
		a ,_ := strconv.ParseFloat(str, 64) 
		c  *= a 
	}
	*x += c 
	return
}

func main() {
	var c float64
	var NewS string 
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		NewS = scanner.Text()
		Conto(NewS, &c )
	}
	fmt.Println(c)
}