package main
import (
  "fmt"
  "strconv"
)
func giorniInMese(mese  int)  (g int) {
    switch mese  {
    case 2 : g=28
    case 4 , 6 , 9 , 11 : g=30
    default : g = 31
    }
    return
  }
func main() {
var s string
fmt.Scan(&s)
mese, _:= strconv.Atoi(s[3:5])
fmt.Printf("il mese %d ha %d giorni\n", mese, giorniInMese(mese) )
}
