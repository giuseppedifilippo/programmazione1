package main
import "fmt"
func main() {
var n,i int
fmt.Scan(&n)
for i=1;n>0 ; i++ {
  if (n%10)%2==0 {
    fmt.Println("posizione della cifra pari", i)
    break
  }
  n/=10
}
if n==0 {
  fmt.Println(-1)
}
}
//fatto
