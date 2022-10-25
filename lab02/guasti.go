package main
import "fmt"
func main () {
var x1,x2,x3 int
fmt.Scan(&x1,&x2,&x3)
v := []int{x1,x2,x3}
v1 :=[]string{"caricamento acqua", "scarico acqua", "sistema di riscaldamento"}
if x1!=0 && x2!=0 && x3 !=0 {
  for i:=0;i<3;i++ {
    if v[i]>0 && v[i]<=3 {
      fmt.Println(v1[v[i]-1])
    }
  }
}
}
/*odio gli if quindi ho usato altri array*/
