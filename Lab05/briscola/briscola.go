package main
import "fmt"
func punti(s string) int {
  var c int
  for _, ch:=range s {
    switch string(ch) {
    case "A": c+=11
    case "3" : c+=10
    case "K" : c+=4
    case "Q" : c+=3
    case "J" : c+=2
    default : c+=0
    }
  }
  return c
}
func main() {
var s string
fmt.Scan(&s)
fmt.Println(punti(s))
}// fatto
