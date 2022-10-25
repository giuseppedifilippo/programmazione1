package main
import "fmt"
func main (){
var min , max int
fmt.Scan(&min , &max)
if min > max {
  min = min + max
  max = min - max
  min = min - max
}
fmt.Println(max)
}
