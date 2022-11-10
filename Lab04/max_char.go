package main
import "fmt"
func main() {
var a,s rune
for i:=0; i<5;i++ {
  fmt.Println("i vale", i)
  fmt.Scanf("%c\n", &s)
  fmt.Println(int(s))
  if int(s)>int(a) {
    a=s
  }

}
fmt.Printf("%c", a)
} // fatto
