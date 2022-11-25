package main
import (
      "fmt"
      "os"
)
func main () {
  input  := os.Args[1:]
  fmt.Printf("hai scritto %q\n", input)
  fmt.Printf("%T", input )
}
