package main
import (
        "fmt"
        "strings"
        "strconv"
        )
const VOCALI = "aeiouAEIOU"
const CIFRE = "0123456789"
const S1 = "s"
const S2 = "sa"

//1 controlla se la stringa letta contiene S2
func isS2(x string)   {
  if strings.Contains(x, S2) {
    fmt.Printf("%c contiene %c\n",x,S2 )
  } else {
    fmt.Printf("%c  non contiene %c\n",x,S2 )
  }
}

//2 se la stringa letta contine almeno una vocale ERRORE
func isVoc(x string)  {
  if strings.ContainsAny(x, VOCALI) {
    fmt.Printf("%c contiene vocali\n",x )
    return b
  } else {
    fmt.Printf("%c  non contiene vocali\n",x)
    re
  }
}

//3 quante occorrenze di S1 ha la stringa letta
func isS1(x string) {
  c:= strings.Count(x, S1)
  fmt.Printf("%c contine %d %c\n", x,c,S1)
}

//in che posizione si trova la prima vocale della stringa letta
func firstVoc(x string) int {
  return strings.IndexAny(x, VOCALI))
}
func lastVoc(x string) int {

}
