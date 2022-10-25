package main
import "fmt"
func main () {
var a int
var stud bool
fmt.Println("età")
fmt.Scan(&a)
fmt.Println("studente?")
fmt.Scan(&stud)
age := []int{0,9,18,26,65,2000}
tariffe:=[]string{"gratis","5","7.5","10","7.5"}
if a>= 18 && stud {
  fmt.Println("ingresso 5 euro")
} else {
  for i:=0;i<6; i++ {
    if a>age[i]&& a<age[i+1] {
      fmt.Println("ingresso ",tariffe[i],"euro")
      break
    }
  }
}
}
//odio gli if ripetuti
