package main
import (
    "fmt"
    "math/rand"
    "time"
    "strconv"
)
type Carta struct {
    valore, seme string
}


//dato un numero ritorna la carta corrispondente in forma di struct
func carta(n int) (Carta, bool) { // la boolena in questo caso è inutile dato che i numeri che gli passiamo sono già da 0 a 52 ma le specifiche ne richiedono la presenza
  var c Carta
  var s, v string
  if n<0 || n>52 {
    return c , false
  }
  switch n/13 {
    case 0 : s = "cuori"
    case 1 : s = "quadri"
    case 2 : s= "fiori"
    case 3 : s= "picche"
  }
  switch n%13 {
    case 0 : v = "A"
    case 10 : v ="J"
    case 11 : v ="Q"
    case 12 : v="K"
    default : v = strconv.Itoa(n%13 + 1 )
      }
      c.valore = v
      c.seme = s
      return c , true
}
//estrae un numero a caso da 0 a 52
func estraiCarta() Carta {
  c,_:= carta(rand.Intn(52))
  return c
}
func dai4Carte() [4]Carta {
  var a [4]Carta
  for i:=0;i<4;i++ {
    a[i]=estraiCarta()
  }
  return a
}

func main() {
  rand.Seed(time.Now().Unix())
  fmt.Println(estraiCarta()) 
  fmt.Println(dai4Carte())
}
