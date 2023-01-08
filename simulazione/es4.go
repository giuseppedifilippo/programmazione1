package main

import (
	"os"
	"strconv"
  "fmt"
	"strings"
)

func main() {
	n := os.Args[1]
	d, _ := strconv.Atoi(os.Args[2])
	strsl := strings.Split(n, "")
	sl := []int{}
	//creo una slice che contiene le cifre di n in tipp intero
	for _, el := range strsl {
		x ,_ := strconv.Atoi(el)
		sl = append(sl, x)
	}

	for i:=0; i< d ; i++ {
		var num, index int  
		for j , el := range sl {
			if el > num {
				num, index  = el ,j
			}
		}
		copy(sl[index:],sl[index+1:])
		sl[len(sl)-1] = 0
		sl = sl[:len(sl)-1]
		
	}
  out := ""
  for _, el := range sl {
    out += strconv.Itoa(el)
  }
	fmt.Printf("numero migliore: %s", out)
}