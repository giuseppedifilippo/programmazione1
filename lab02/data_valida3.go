package main
import "fmt"
func isBi(y int) bool {
    return y%4==0 && (y%100!=0 || y%400==0)
  func main() {
    var gg,mm ,aa int
    mesi:=[]int{1,31,2,28,3,31,4,30,5,31,6,30,7,31,8,31,9,30,10,31,11,30,12,31}
    fmt.Scan(&gg,&mm,&aa)
    a:= gg > 0
    b:= mm>0 && mm<=12
    c:= aa>0
    if !a || !b || !c {
      fmt.Println("data non valida")
    } else {
      for i:=0;i<len(mesi);i+=2 {
        if mm == mesi[i] {
          if isBi(aa) {
            if gg > 29 {
              fmt.Println("data non valida")
            } else {
              fmt.Println("data valida")
            }
          } else if gg > mesi[i+1] {
            fmt.Println("data non valida")
          } else {
            fmt.Println("data valida")
          }
        }
      }
    }
    }
