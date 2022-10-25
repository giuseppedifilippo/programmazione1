package main
import "fmt"
func main() {
  var d, l, consumo, resa  float64
  fmt.Println("inserire distanza e poi carburante usato")
  fmt.Scan(&d, &l)
  consumo = l/d
  resa = d/l
  fmt.Println("consumo medio:", consumo)
  fmt.Println("resa", resa)

}
