package main

import "fmt"
func main() {
var s string
fmt.Scanf("%s", &s)
for c:=0; c<len(s); c++ {
  fmt.Println("ciao")
  for i:=-1;i<2;i++ {
    fmt.Print(string(int(s[c])+i))
  }
  fmt.Println()
}
}
//fatto
