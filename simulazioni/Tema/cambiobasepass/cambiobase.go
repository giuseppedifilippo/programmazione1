package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	basi := "0123456789ABCDEF"
	num,_  := strconv.Atoi(os.Args[1])
	base,_ := strconv.Atoi(os.Args[2])
	out := ""
	for  num != 0  {
		resto := num%base
		for i, char := range basi {
			if resto ==i {
				out += string(char)
			}
		}
		num = num/base
	}
	for i := len(out)-1; i >= 0; i-- {
		fmt.Print(string(out[i]))
	}
	fmt.Println()
}// il programma va per sbaglio ho fatto tutto nel main al posto di creare una funzione a parte e quindi i test ora non vanno 