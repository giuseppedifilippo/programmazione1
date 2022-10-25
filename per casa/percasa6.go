// dat idue numeri trovare MCD ricorsivamente
package main
import "fmt"
func main () {
  var n1,n2, i int
  fmt.Scan(&n1, &n2)
  if n2 < n1 {
    n1 = n1 + n2
    n2 = n1 - n2
    n1 = n1 - n2
  }
  i = n1
  for true {
    if n1%i== 0  && n2%i==0 {
      fmt.Println(i)
      break
    }
    i--
  }
}// fatto
