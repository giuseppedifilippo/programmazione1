package main
import "fmt"
func main () {
var n int
fmt.Scan(&n)
a := n%10==0
b:= n>0
if a && b  {
  fmt.Println("positivo divisibile per 10")
} else if b {
  fmt.Println("positivo")
}else if a {
  fmt.Println("divisibile per 10")
} /*per chi legge cambiare ordine degli if sfancula il programma*/
}
