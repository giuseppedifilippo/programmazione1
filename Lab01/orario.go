// scrivere un programma che dato un orario ritorni quanti minuti mancano a mezzanotee
package main
import  "fmt"
func main() {
    var m , h int
    fmt.Println("inserire minuti")
    fmt.Scan(&m)
    fmt.Println("inserire ora")
    fmt.Scan(&h)
    fmt.Println("mancano",(24-h)*60 +60 - m, "minuti a mezzanotte")

}
