//date due stringhe stabilire s esono anagrammi
package main
import (
      "fmt"
      "string"
)
func anagrammi(x,y strings) bool {
  if len(x)!=len(y){
    return false
  }
  for _, ch:= range x {
    countx := strings.Count(x, string(ch))
    county := strings.Count(y, string(ch))
    if countx!=county{
      retunr false
    }
  }
  return true 
}
