package main
import "fmt"
func main() {
var n,count int
nv:=1
for {
  fmt.Scan(&n)
  if n==2 {
    break
  } else if n==1 && nv==0{
    count++
  }
  nv=n
}
fmt.Println(count)
}//fatto
