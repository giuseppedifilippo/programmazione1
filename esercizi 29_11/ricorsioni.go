//funzone che dato n restituisce un sliced istringheci tutte l posibili
//stringhe di 0 e 1 dilunghezza n 
package main
import ( "fmt"
        "strconv"
        "os"
        )
func f (n int ) []string {
  if n == 0 {
    return []string{""} // caso base 
  }
s:= f(n-1)
  var res []string 
  for _, x:=range s {
    res = append(res, x+"0",x+"1")
}
  return res
}
func main() {
  n,_ := strconv.Atoi(os.Args()[1:])
  fmt.Println(f(n))
}