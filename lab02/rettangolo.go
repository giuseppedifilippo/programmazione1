package main
import "fmt"
import "math"
func main (){
var x1,y1,x2,y2,x3,y3 int
fmt.Scan(&x1,&y1,&x2,&y2,&x3,&y3)
xi, xii :=  math.Min(float64(x1), float64(x2)), math.Max(float64(x1), float64(x2))
yi,yii := math.Min(float64(y1),float64(y2)), math.Max(float64(y1),float64(y2))
xiii, yiii := float64(x3), float64(y3)
if xiii > xi && xiii<xii && yiii > yi && yiii < yii {
  fmt.Println("Interno")
} else if xiii == xi || xiii == xii || yiii==yi || yiii==yii {
  fmt.Println("Perimetrale")
} else {
  fmt.Println("Esterno")
}

}
