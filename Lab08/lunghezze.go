package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sort"
)    

func aggiornaConteggio(m map[int]int, riga string) {
	sl:= strings.Split(riga, " ")
  for _, el := range sl {
  	m[len(el)]++
  }
}
func main() {
  mappa := make(map[int]int)
	scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
  	aggiornaConteggio(mappa, scanner.Text())
  }
  list := []int{}
  for  k := range mappa {
	list = append(list, k )
  }  
  sort.Ints(list)
  for _, el := range list {
	fmt.Printf("ci sono %d parole di lunghezza %d", mappa[el], el)
  }

}