package main
import "fmt"
func main (){
const giorno = 24*60*60
const ore = 60*60
const minuti = 60
var s,d,h,m,r int
fmt.Scan(&s)
d = s/giorno
h = (s%giorno)/ore
m = (s%giorno)%ore
r = ((s%giorno)%ore)/minuti
fmt.Println(d,"giorni", h,"ore", m, "minuti", r, "secondi")


}
