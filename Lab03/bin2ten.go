package main
import (
"fmt"
"math"
)
func digit(x int) int {
  var c int
  for x>0 {
    x/=10
    c++
  }
  return c
}
func isBin(x int) bool {
for x>0 {
  if !(x%10==0 || x%10==1) {
    return false
  }
  x/=10
}
return true
}

func main() {
var n,c int
fmt.Scan(&n)
if isBin(n)==false  {
  fmt.Println("input errato inserire un numero binario")
} else {
  x:=n
  for i:=0; i<digit(x); i++ {
    a:= n%10
    c=c + a * (int(math.Pow(2.0 , float64(i))))
    n/=10
}
fmt.Println(c)
}
}
//fatto
