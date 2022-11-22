//data una stringa stabilire se ci sono caratteri ripetuti
package main
import "fmt"
func ripetizioni( x string) bool {
  for i:=0; i<len(x); i++ {
    for j:=i+1; j<len(x); j++{
      if x[i] == x[j]{
        return true
      }
    }
  }
  return false

func main() {
  var x string
  fmt.Scan(&x)
  if ripetizioni(x){
    fmt.Println("ci sono ripetizioni")
    } else {
      fmt.Println("non ci sono ripetizioni")
    }
}
