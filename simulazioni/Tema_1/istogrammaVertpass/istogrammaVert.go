package main

import (
	"fmt"
	"os"
	"strconv"
)
func Max(sl []int) int {
	var max int 
	for i := 0 ; i <len(sl); i++ {
		if sl[i] > max {
			max = sl[i]
		}
	}
	return max
}
func main() {
	in:= os.Args[1:]
	sl := []int{}
	for _, el := range in {
		x,_ := strconv.Atoi(el)
		sl = append(sl,x)
	}
	max := Max(sl)
	for i:=max; i>0;i--{
		for _, el := range sl {
			if el < i {
				fmt.Print(" ")
			}else{
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}//passa tutti i test e prima botta