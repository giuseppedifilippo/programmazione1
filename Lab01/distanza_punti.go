package main
import "fmt"
import "math"
func main () {
  var x1,x2,y1,y2,d1,d2 float64
  fmt.Scan(&x1, &x2, &y1, &y2)
  d1 = x2 - x1
  d2 = y2 - y1
  fmt.Println(math.Sqrt(math.Pow(d1, 2 ) + math.Pow(d2, 2 )))

}
