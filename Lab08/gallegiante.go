package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SPLITTA LA SLICE DI STRINGHE IN UNA SLICE  DI SLICE DI STRINGHE
func reSlice(x []string) (a [][]string) {
  for _, el := range x {
    a = append(a,strings.Split(el, " "))
  }
  return a
}



// RIORGANIZZA UNA SINGOLA SLICE DI STRINGHE METTENDO PRIMA LE LETTERE E POI GLI ASTRERISCHI
func reorg(x [][]string) {
	if len(x)<=2 {
		return
	}
	for i:=len(x)-1; i>=1; i-- {
	  for c:=0;c<len(x[0]);c++ {
			if x[i][c]=="*" && x[i-1][c]!="*"{
		  		x[i][c], x[i-1][c]= x[i-1][c], x[i][c]
	  		}
	 		reorg(x[1:])
		}
	}		
}


//RIORGANIZZA LA SLICE DI SLICE (POSSIBILMENTE RIORGANIZZATA) IN UNA STAMPA PIù ORDINATA 
func printer(x [][]string) {
  for i:=0;i<len(x);i++ {
    fmt.Printf("%s\n", strings.Join(x[i]," "))
  }
}



func main() {
  var s []string
  s = make([]string,0)
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    s = append(s, scanner.Text())
  }
  n := reSlice(s)
  reorg(n)
  printer(n)
}//FATTO 
