package main

import (
	"fmt"
	"os"
)
func main() {
	in := os.Args[1]
	prev := in[0]
	for _, el := range in {
		if byte(el) < prev {
			fmt.Println()
			fmt.Print(el-48)
			//fmt.Println()
		}else {
			fmt.Print(el-48)
		}
		prev = byte(el)
	}
	fmt.Println()
}//passato