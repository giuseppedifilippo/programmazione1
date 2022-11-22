//data una data stabilire che giorno della settimana è
package main
import "fmt"
type data struct {
  g, m, a int
}
func isBisestile(a int ) bool {
  return a%4==0 && (a%100==0 || a%400==0)
}
func lunghezzaMese(m,a int ) int {
  switch m {
    case 11,4,6,9: return 30
  case 2 :
    if isBisestile(a){
      return 29
    } else {
      return 28
    }
    default: return 31
    }
  }
  func lunghezzaAnno(a int ) int {
    if isBisestile(a) {
      return 366
    }else {
      return 365
    }
  }

  func giorniDaEpoca(d data ) int {

  }
  func dataToStr(d data) string{
    return fmt.Sprintf("%02d/%02d/%04d", d.g,d.m,d.a)
  }

  func giornoDellaSettimana(d data) string {
    g:= giorniDaEpoca(d)
    switch g%7 {
    case 0 : return "giovedi"
    }
  }
  func main() {
    var x string
  }
