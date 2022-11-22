package main
import "fmt"
const DIM = 5
func reverse (x *[DIM]int) {
  for i:=0;i<DIM/2;i++ {
    (*x)[i] ,(*x)[DIM-1-i]= (*x)[DIM-1-i], (*x)[i]
  }
}
func switchFirstLast(x *[DIM]int) {
  (*x)[0],(*x)[DIM-1]= (*x)[DIM-1],(*x)[0]
}
func main() {
  var arr [DIM]int
  for i:=0;i<DIM;i++ {
    arr[i]=i
    }
    fmt.Println(arr)
    reverse(&arr)
    fmt.Println(arr)
    switchFirstLast(&arr)
    fmt.Println(arr)
}// fatto
