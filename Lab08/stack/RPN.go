package main 
import ("fmt"
        "bufio"
        "os"
        "strconv"
        )
func operate(op string, x *[]float64) {
  if len(*x) <2 {
    fmt.Println("not enough data")
    return
  }
  a,b := (*x)[len(*x)-2], (*x)[len(*x)-1]
  *x = (*x)[:len(*x)-2]
  switch op {
    case "+": (*x) = append(*x, a+b)
    case "-": (*x) = append(*x, a-b)
    case "*": (*x) = append(*x, a*b)
    default : (*x) = append(*x, a/b)
  }
}
func main() {
  var s []float64
  s=make([]float64,0)
  scanner := bufio.NewScanner(os.Stdin)
  myloop :
  for {
    fmt.Println("Next? +, -, *, /,o un numero?")
    scanner.Scan()
    switch scanner.Text(){
        case "+","-","*","/" : operate(scanner.Text(), &s)
        case "quit": break myloop
        default : b,_ := strconv.ParseFloat(scanner.Text(), 64);s=append(s,b)
    }
    fmt.Println(s)
  }
}