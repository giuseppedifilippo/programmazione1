package main
import "fmt"
func main() {
var m1,q1,m2,q2, float64
fmt.Scan(&m1,&q1,&m2,&q2,)
x := (q1-q2)/(m2-m1)
y := m1*x + q1
fmt.Println(x, ",", y)
}
