package main
import "fmt"
func main () {
  var n1, n2, n3 int
  fmt.Scan(&n1, &n2, &n3)
  if n1 < n2 && n2 < n3 {
    fmt.Println(n1 , n2 , n3)
  } else if n3 < n1 && n1 < n2  {
    fmt.Println(n3 , n1 ,n2 )
  } else if n2 < n3 && n3 < n1 {
    fmt.Println(n2 , n3 , n1)
  }

} // fatto 
