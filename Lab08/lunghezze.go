package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
  fmt.Println(mappa)
}