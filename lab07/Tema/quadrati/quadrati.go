package main
import (
      "fmt"
      //"os"
      "math"
      )

func isSquare(n int ) bool {
  return math.Sqrt(float64(n))/1 == 0.0
}
func main() {
  //arr := os.Args[1:]
  var x int
  fmt.Scan(&x)
  fmt.Println(isSquare(x))
}
