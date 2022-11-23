package main // programma funzionante ma manca di gestione di input errati
import (    // ma fallisce una dei test
      "fmt"
      "time"
)
type  Clock struct {
  hour, min, sec int
}

func updateHour(x *Clock) {
  if (*x).hour == 0 {
    return
  }
  (*x).hour -= 1
  (*x).min = 59
  (*x).sec = 60
}
func updateMin(x *Clock) { // spero che questa funzione sia terminata
  if (*x).min == 0 {
    updateHour(x)
    return
  }
  (*x).min -= 1
  (*x).sec = 59
}
func countdown(x *Clock)  {
  for {
    fmt.Println(*x)
    if (*x).sec == 0 {
      updateMin(x)
    }
    if (*x).sec== 0 && (*x).min == 0 && (*x).hour == 0 {
      break
    }

  time.Sleep(time.Duration(1) * time.Second)
  (*x).sec -= 1
  }
  return
  }

  func main() {
    var c Clock
    //fmt.Scan(&c.hour, &c.min, &c.sec)
    /*hhour := (*x).hour < 0     // se sono nel formato erraro ritenta la scansione fino a che non soddisfa i requisiti
    hmin := (*x).min < 0 || (*x).min >= 60
    hsec := (*x).sec < 0 || (*x).sec >= 60*/
    for {
      var hhour, hmin, hsec bool
      c = Clock{ 0 , 0 , 0 }
      fmt.Println(c)
      fmt.Scan(&c.hour, &c.min, &c.sec)
      
      hhour = c.hour < 0     // se sono nel formato erraro ritenta la scansione fino a che non soddisfa i requisiti
      hmin = c.min < 0 || c.min >= 60
      hsec = c.sec < 0 || c.sec >= 60
      if !(hhour && hmin && hsec) {
        fmt.Println("inserimento nel formato errato, riprova")
      //  fmt.Scan(&(*x).hour, &(*x).min, &(*x).sec)
      } else {
        break
      }
    }
    countdown(&c)
  }
