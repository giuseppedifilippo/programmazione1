package main
import (
      "fmt"
      "os"
      "math"
      "strconv"
      )

func isSquare(n int ) bool {
x:= math.Sqrt(float64(n))
str := fmt.Sprintf("%2.2f", x )
a, _:= strconv.Atoi(str[len(str)-2:])
return a == 0
}
func main() {
  var vec []int
  vec = make([]int, 0)
  arr := os.Args[1:]
  for i:=0 ; i<len(arr); i++ {
    x, err := strconv.Atoi(arr[i])
    if err == nil && isSquare(x){
    vec = append(vec, x)
  }
  }
  fmt.Println(vec)
}// funziona alla carta ma fallisce il test, probabile che il programma di test voglia un tipo di output ben specifico ma dal comando non riesco a capirlo  
