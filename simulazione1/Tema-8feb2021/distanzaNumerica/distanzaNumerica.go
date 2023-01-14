package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	in := os.Args[1:]
	if len(in) <= 1 {
		fmt.Println("NON OK")
		os.Exit(2)
	}
	numsl := []int{}
	for _, el := range in {
		x, _ := strconv.Atoi(el)
		numsl = append(numsl, x)
	}
	distanza := int(math.Abs(float64(numsl[0]-numsl[1])))
	sum := 0 
	for i:= 0 ; i<= len(numsl)-2; i++ {
		sum += numsl[i]
		if int(math.Abs(float64(numsl[i]-numsl[i+1]))) != distanza {
			fmt.Println("NON OK")
			os.Exit(1)
		}
	}
	sum += numsl[len(numsl)-1]
	fmt.Println(sum)
}//passato