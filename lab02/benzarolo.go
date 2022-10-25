package main
import "fmt"
func main() {
  var l  float64
  var t int
  fmt.Scan(&l,&t)
  f := []float64{1.78 , 1.98, 1.2, 1.1}
  fmt.Println(l*f[t])
  /*per chi sta leggendo questo esercizio, ho usato
  un array se non è un argomento che abbiamo fatto
  perchè avevo zero sbatti di fare 4 if */ 
}
