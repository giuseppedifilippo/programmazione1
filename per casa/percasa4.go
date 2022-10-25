//n4, verificare se tuute le cifre di un numero di al massimo tre cifre sono dispari
package main
import "fmt"
func main () {
  var n int
  fmt.Scan(&n)
  if n < 999 {
    for true {
      if (n%10)%2 == 0 {
        fmt.Println("numero pari individuato")
        break
      } else {
        n/=10
      }
      if n/10 == 0 {
        fmt.Println("tutti le cifre sono dispari")
        break
      }
      n /= 10
    }
  }
}
