package main
import "fmt"
func main() {
var n,a int
for i:=0; i<10; i++ {
  fmt.Scan(&n)
  if n>a {
    a=n
  }
}
}
