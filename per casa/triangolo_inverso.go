package main
import "fmt"
func main() {
var n int
fmt.Scan(&n)
for r:=0;r<n;r++ {
  for c:=0;c<n;c++ {
    if c<n-r {
      fmt.Print(" ")
    } else {
      fmt.Print("*")
    }
  }
  fmt.Println()
}
}
