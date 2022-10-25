package main
import "fmt"
import "math"
func main () {
var n float64
fmt.Scan(&n)
p :=  math.Pow(10, -6)
y := int(n)
 fmt.Println(y == 10)
 fmt.Println(y != 10)
 fmt.Println(y != 10 && y != 20)
 fmt.Println( y != 10 || y != 20)
 fmt.Println( y >= 10)
 fmt.Println(y >= 10 && y <= 20)
 fmt.Println(y > 10 && y < 20)
 fmt.Println(y >= 10 && y < 20)
 fmt.Println(y < 10 && y > 20)
 fmt.Println((y > 10 && y < 20) || (y > 30 && y < 40))
 fmt.Println(y%4 == 0 && y%100 != 0)
 fmt.Println(y%2==1 && y > 0 && y<100)
 fmt.Println(n > (10.0-p) && (n < 10.0 + p))
}
