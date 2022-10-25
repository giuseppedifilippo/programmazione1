package main
import "fmt"
func main () {
var g1,i1,f1,g2,i2,f2 int
fmt.Scan(&g1,&i1,&f1,&g2,&i2,&f2)
if g1 == g2 {
  if i2 >= i1 && i2<=f1 {
    fmt.Println("si sovrappongono")
  } else if f2 >= i1 && f2 <= f1 {
    fmt.Println("si sovrappongono")
  } else {
    fmt.Println("non si sovrappongono")
  }
} else {
  fmt.Println("non si sovrappongono")

}
}
