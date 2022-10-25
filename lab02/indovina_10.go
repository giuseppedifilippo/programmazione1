package main
import "fmt"
func main () {
const x = 4
var n int
fmt.Scan(&n)
if n == x {
  fmt.Println("Hai indovinato!")
} else {
  if n < 1 || n > 10 {
    fmt.Println("Valore non valido")
  } else {
    fmt.Println("Non hai indovinato")
  }
}
}
