package main
import (
        "fmt"
        "bufio"
        "os"
        "strings"
        "strconv"
        "math"
)
func main() {
  var pol float64
  scanner:=bufio.NewScanner(os.Stdin)
  scanner.Scan()
  st:= strings.Split(scanner.Text(), " ")
  x,_:=strconv.Atoi(st[len(st)-1])
  st= st[:len(st)-1]
  for i, n:= range st {
    y, err := strconv.Atoi(n)
    if err!=nil {
      fmt.Println("fucked up big time")
      return
    }
    pol+= float64(y) * (math.Pow(float64(x), float64(i)))
  }
  fmt.Println(pol)
}//fatto 
