// per questo progetto ho deciso di deviare dalle specifiche e creare un programma che simula una partita di blackjack senza badare esattamente
// alle funzioni descritte nel comando
// tra l utente e la CPU, infatti sono andato un pochino overkill oltre che il programma può certamente essere ottimizzato
// ma buona parte del programma si basa su un programma modificato che a sua volta nasce da un altro programma modificato
// p.s. tutte le funzioni di os/exec non so come funzionino bene ma servono soltanto a non farti premere invio ogni volta che vuoi inserire un comando
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)
type Carta struct {
    valore, seme string
}


//dato un numero ritorna la carta corrispondente in forma di struct
func carta(n int) (Carta, bool) { // la boolena in questo caso è inutile dato che i numeri che gli passiamo sono già da 0 a 52 ma le specifiche ne richiedono la presenza
  var c Carta
  var s, v string
  if n<0 || n>52 {
    return c , false
  }
  switch n/13 {
    case 0 : s = "♥"
    case 1 : s = "♦"
    case 2 : s= "♣"
    case 3 : s= "♠"
  }
  switch n%13 {
    case 0 : v = "A"
    case 10 : v ="J"
    case 11 : v ="Q"
    case 12 : v="K"
    default : v = strconv.Itoa(n%13 + 1 )
      }
      c.valore = v
      c.seme = s
      return c , true
}


func getValoreBJ(x Carta) int {
  var i int
  switch x.valore {
  case "K", "Q", "J" : i = 10
  case "A" : i =11
  default : i, _ = strconv.Atoi(x.valore)
  }
  return i
}



//calcola il valore della mano 
func mano(x []Carta)  int {
  var c,p   int 
  for _ ,  y:= range x {
    if y.valore == "A" {
      c++ 
      continue
    }
    p+= getValoreBJ(y)
  }
  if p>12 && c!=0 {
    p+= 1*c
  } else {
    p+= 11 * c 
  }
  return p
}


func mazzoBJ() []Carta {
  var x []Carta
  x = make([]Carta, 104)
  for i:=0 ; i<104; i++{
    x[i], _ = carta(i/2)
  }
  return x 
}


func mischia(x []Carta) {
  for i:=0; i<104; i++ {
    r:= rand.Intn(104-i) + i
    x[i], x[r] = x[r], x[i]
  }

}


//ritorna una carta dalla cima del mazzo e la rimuove dal mazzo 
func preleva(x *[]Carta) (c Carta) {
  c = (*x)[len(*x)-1]
  *x=(*x)[:len(*x)-1]
  return c
}

//handling del gioco del player 
func giocatore(x *[]Carta, m *[]Carta) {
  var b []byte = make([]byte, 1)
  var pun int 
  fmt.Println(*x)
  for {
    fmt.Printf("questa è  la tua mano\n")
    fmt.Println(*x)
    fmt.Println("pescare un'altra carta [Y/N]?")
    os.Stdin.Read(b)
    if string(b)=="Y" || string(b)=="y" {
      *x = append((*x),preleva(m))
    } else if string(b)=="N" || string(b)=="n"{
      break
    }
    pun = mano(*x)
    if pun == 21 {
      fmt.Println("BLACKJACK")
      break
    } else if pun > 22 {
      fmt.Printf("%d punti, hai sballato\n", pun)
      exec.Command("stty", "-F", "/dev/tty", "echo").Run()//se non 
      os.Exit(1)
    }
  }
}
//GESTIONE DEL GIOCO DEL MAZZIERE 
func dealer(x *[]Carta, m *[]Carta) {
  var pun int 
  for {
    pun = mano(*x)
    if pun < 17 {
      (*x)= append((*x), preleva(m))
    } else if pun >17 && pun<21 {
      break
    } else if pun == 21 {
      fmt.Println("BLACKJACK DEL DEALER")
    } else {
      fmt.Println("il dealer ha sballato, SEI IL VINCITORE")
      exec.Command("stty", "-F", "/dev/tty", "echo").Run()//se non 
      os.Exit(1)
    }
  }
}



func main() {
  //inizializzazione delle funzion propedeutiche 
  //generazione pseudorandomica
  rand.Seed(time.Now().Unix())
  //disabilitazazione dell input buffering per rendere il gioco più veloce 
  exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

// inizializzazione delle mani
  var player, cpu []Carta
  player, cpu = make([]Carta, 0), make([]Carta,0)

// crea e mischia il mazzo 
  maz := mazzoBJ()
  mischia(maz)
  //pesca delle mani per giocatore e dealer 
  for i:=0; i<2;i++ {
    player= append(player, preleva(&maz))
    cpu= append(cpu, preleva(&maz))
  }
  giocatore(&player, &maz)
  fmt.Println("turno del mazziere")
  dealer(&cpu,&maz)
  if mano(player) > mano(cpu) {
    fmt.Printf("HAI VINTO\n")
  } else if mano(player) < mano(cpu) {
    fmt.Printf("HAI PERSO\n")
  } else {
    fmt.Println("PAREGGIO")
  }
  exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}
// IL PROGRMMA è TECNICAMENTE TERMINATO SI POTREBBE ABBELLIRE MA SONO STANCO PEACE OUT 
