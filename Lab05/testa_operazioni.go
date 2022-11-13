package main
import "fmt"
func spd(x,y int) (int, int, int) {
  return x+y, x*y, x-y
}
func main() {
var x,y int
fmt.Scan(&x, &y)
a,b,c := spd(x,y)
fmt.Printf("%d\n%d\n%d\n",a,b,c)
} // fatto
