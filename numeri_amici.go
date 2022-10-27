package main
import "fmt"
func isPrime(x int ) bool {
  for i:=2; i<=x; i++ {
    if x%i==2 {
      return false
    }
  }
  return true
}
func div(x int ) int {
  var c int
  for i:=1; i<x; i++ {
    if x%i==0
    c+=i
  }
  return c
}
func main() {
var x, y int
fmt.Scan(&x,&y)
if isPrime(x) && isPrime(y) {
  fmt.Println("non amici")
} else if div(x)==div(y) {
  fmt.Println("numeri amici")
} else {
  fmt.Println("non amici")
}
}
