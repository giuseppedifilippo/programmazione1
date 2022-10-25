C:\Users\giuse\OneDrive\Desktop\UNI primo anno\programmazione1package main
import "fmt"
func main() {
var red int
var coniugato bool
const nc = 32000
const c = 64000
fmt.Scan(&red,&coniugato)
if coniugato {
  if red>0 && red<c {
    fmt.Println("10%")
  } else {
    fmt.Println("24%")
  }
} else {
    if red > 0 && red<nc {
      fmt.Println("10%")
  } else {
    fmt.Println("24%")
  }
}

}
