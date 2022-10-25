package main
import "fmt"
func main () {
  const con =   9/5
  const b = 32
  var c int
  fmt.Println("inserire temperatura centigrada")
  fmt.Scan(&c)
  fmt.Println(c * con + b )

}
