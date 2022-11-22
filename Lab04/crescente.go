package main
import "fmt"
func main() {
  var x string
  y := ""
  fmt.Scan(&x)
  for i, ch:=range x {
    if x[i]<x[i+1] {
      y = y + string(ch) + "-"
    } else {
      y= y + string(ch)
    }
  }
  fmt.Println(y)
}
