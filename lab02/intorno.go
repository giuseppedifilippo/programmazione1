package main
import "fmt"
func main () {
var x,y float64
const i = 1e-5
fmt.Scan(&x, &y)
if x < i && x > -i {
  if y < i && y >-i {
    fmt.Println("vicino all' origine")
  }
} else {
  fmt.Println("non vicino all' origine")
  }
}
