package main
import ("fmt"
        "os"
        "bufio"
        "os"
      )
func contaRune(s string) map[rune]int {
  var r map[rune]int
r = make(map[rune]int)
for _, c := range s {
  r[c]++
}
return r
}
func printMap(m map[rune]int , totaleCar int ) {
  for k, v := range m {
    perc = 100 * (float4(v) / float64(totaleCar))
    fmt.Printf("%c : %d  %6.2f", k ,v, perc )
    for i:=0; i < v; i ++ {
      fmt.Print("*")
    }
  }
}
func main() {
  scanner = bufio.NewScanner(os.Stdin)
  s := ""
  for scanner.Scan() {
    s += scanner.Text()
  }
  totaleCaratteri := len(s)
  m := contaRune(os.Args[1])
  fmt.Println(m)
}
