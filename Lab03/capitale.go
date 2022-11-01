package main
import "fmt"
func main() {
var cap, in, obj float64
var i int
fmt.Scan(&cap, &in, &obj)
x:=cap
for i=0;x<obj ;i++ {
  x*=in
}
fmt.Println(i,"anni")
}
//fatto
