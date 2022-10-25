package main
import "fmt"
func main () {
var h , m int
fmt.Scan(&h , &m )
orario := float64(h) + (float64(m)/100)
if orario >= 5.30 && orario < 12.30 {
  fmt.Println("mattino")
}
}
