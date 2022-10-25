package main
import "fmt"
func main () {
  var y int
  fmt.Scan(&y)
  if y % 4 == 0 {
    if y%100 != 0 || y%400 == 0 {
      fmt.Println("anno bisestile")
    } else {
      fmt.Println("non bisestile")
    }
  } else {
    fmt.Println("non bisestile")
  }
} // fatto
