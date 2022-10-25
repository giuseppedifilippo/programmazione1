package main
import "fmt"
func sw(c,n,r,a int ) {
  v1:= c<r
  v2:=c<n-r
  switch true {
  case v1 && a==0 :fmt.Print("*")
  case !v1 && a==0 : fmt.Print(" ")
  case v1 && a == 180 : fmt.Print(" ")
  case !v1 && a == 180 :fmt.Print("*")
  case v2 && a ==90: fmt.Print("*")
  case !v2 && a == 90 : fmt.Print(" ")
  case v2 && a ==270 : fmt.Print(" ")
  case !v2 && a == 270  : fmt.Print("*")
}
}
func main () {
var n,a int
fmt.Println("inserire n")
fmt.Scan(&n)
fmt.Println("inserire angolo di rotazione")
fmt.Scan(&a)
  for r:=0;r<n;r++ {
    for c:=0;c<n;c++ {
      sw(c,n,r,a)
    }
    fmt.Println()
  }
}
