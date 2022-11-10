package main
import (
      "fmt"
      "unicode"
)
func main() {
var s string
var cmin, cmai int
fmt.Scanf("%s", &s)
for _, ch := range s {
  if unicode.IsLower(ch) {
    cmin++
  } else if unicode.IsUpper(ch) {
    cmai++
  }
}
if cmin!=0 && cmai!=0 {
  fmt.Println("sia minuscole che maiuscole")
} else if cmin!=0 {
  fmt.Println("solo minuscole")
} else {
  fmt.Println("solo maiuscole")
}
}
// fatto
