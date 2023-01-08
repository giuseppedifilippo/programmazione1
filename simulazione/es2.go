package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	in := os.Args[1]
	mappa := map[string]int{}
  fmt.Println(len(in))
	for i := 0; i < len(in)-2; i++ {
		for j := i+1; j < len(in); j++ {
			if in[i] == in[j] && len(in[i:j]) >=3 {
				mappa[in[i:j+1]]++
			}
		}
	}
  lista := []string{}
  for  k := range mappa {
    lista = append(lista, k )
  }
  sort.Strings(lista)
  for _, el := range lista {
    fmt.Printf("%s -> Occorrenze: %d\n", el, mappa[el] )
  }
}
