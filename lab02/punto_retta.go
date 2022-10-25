package main
import "fmt"
func main () {
var x,y,m,q float64
fmt.Scan(&x,&y,&m,&q)
yr := m*x + q
if y > yr {
  fmt.Println("sopra")
} else if y==yr {
  fmt.Println("sulla retta")
} else {
  fmt.Println("sotto")
}
}
