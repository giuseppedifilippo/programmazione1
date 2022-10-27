package main
import "fmt"
import "math"
func main() {
var n,a int
d:= 20
const target = 17
for i:=0;i<5;i++ {
  fmt.Scan(&n)
  a=n
  n = int(math.Abs(float64(n-17)))
  fmt.Println(n)
  if n<d {
    d=a
  }
}
  fmt.Println("il piu vicino è",a)
}
