package main
import "fmt"
func main () {
var num1, den1, num2,den2 int
fmt.Scan(&num1, &den1, &num2, &den2)
if num1*den2 == num2*den1 {
  fmt.Println("equivalenti")
} else {
  fmt.Println("non equivalenti")
}
}
