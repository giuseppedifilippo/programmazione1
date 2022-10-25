package main
import "fmt"
func main () {
var h int
fmt.Scan(&h)
if h<0 || h > 23 {
  fmt.Println("orario invalido")
} else {
  if h>=0 && h<=6 {
    fmt.Println("notte")
  } else if h>=7 && h<=13 {
    fmt.Println("mattino")
  } else if h>=14 && h<=18 {
    fmt.Println("pomeriggio")
  } else {
    fmt.Println("sera")
  }

  }
}//fatto
