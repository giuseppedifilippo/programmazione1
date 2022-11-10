package main
import "fmt"
func main() {
var key int
var s string
fmt.Scan(&s)
fmt.Scan(&key)
for _, ch := range s {
  fmt.Printf("%c", string(int(ch)+key)%97)
}
}
