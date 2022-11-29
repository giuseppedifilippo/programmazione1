//data una stringa produrre un aslice dei  suoi anagrammi senza ripetizioni
package main
import "fmt"
import "os"
func anag(s string) []string {
  if len(s)==0 {
    return []string{""}
  }
  var res []string
  for i:=0; i <len(s);i++ {
    primalettera := string(s[i])
    resto:= s[:i] + s[i+1:]
    z:= anag(resto)
    for _, x:=range z {
      res = append(res, primalettera + x)
    }
  }
  return res
}
func main() {
  fmt.Println(anag(os.Args[1]))
}//funziona