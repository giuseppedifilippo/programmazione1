package main
import "fmt"
func main() {
var totLen, obj int
var s string
p:= ""
fmt.Scan(&obj)
for obj>len(p) {
  fmt.Scan(&s)
  fmt.Println(totLen)
  totLen=totLen + len(s)
  fmt.Println(p)
  p=p+s
}
fmt.Println(totLen, p)
}//fatto 
