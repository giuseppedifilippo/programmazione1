package choba
import ("fmt"
        "strings"
        )
func Choba(n int ) {
  switch n {
    case 0 : fmt.Println("choba")
    default : fmt.Println(strings.Repeat("choba", n))
  }
}