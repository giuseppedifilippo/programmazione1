package main

import (
	"fmt"
	"os"
	"strconv"
)

type Rettangolo struct{
	b,h int 
}

func (x Rettangolo) String() string {
	var dis string 
	for i:=0; i<x.h; i++ {
		for j:=0; i<x.b; j++ {
			dis += "B"
		}
		dis += "\n"
	}
	return dis 
}

func main() {
	var r Rettangolo
	in := os.Args[1:3]
	r.b, _ = strconv.Atoi(in[0]) 
	r.h, _ = strconv.Atoi(in[1])
	fmt.Println(r.String())
}