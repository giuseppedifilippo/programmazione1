package main
import "fmt"
func main() {
  var n int
  fmt.Println("inserire numero")
  fmt.Scan(&n)
  if n >= 1000 {
    if n%1000 == 0 {
      fmt.Println("termina con tre zeri")
      } else {
        fmt.Println("non termina con tre zeri")
      }
} else {
  fmt.Println("non termina con tre zeri")
}

} // fatto 
