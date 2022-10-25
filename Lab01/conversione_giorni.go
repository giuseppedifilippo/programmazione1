package main
import "fmt"
func main () {
  var g, anni, sett, gg int
  fmt.Scan(&g)
  anni = g/365
  sett = (g%365)/7
  gg = (g%365)%7
  fmt.Println(anni, "anni ;", sett,"settimane ;", gg, "giorni")
}
