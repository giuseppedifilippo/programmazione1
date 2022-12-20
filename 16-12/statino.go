//prende in input un file di testo che contiene nominativi e voti e ne elabora la media pesata
package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
  "os"
)
type studente struct {
  nome string
  voti []int
}
//fa la media dei voti in trentesimi
func avg(voti []int) float64 {
  s := 0 
  for _, voto := range voti {
    s+= voto
  }
  return float64(s)/float64(len(voti))
}

//converte la media da trentesimi in centesimi
func convert(media float64) float64 {
  return media*10.0/3.0
}
//elabora riga a riga del csv
func parseLine(riga string) (s studente, ok bool) {
  x := strings.Split(riga, ",")
  if len(x) < 2 {
    return
  }
  s.nome = x[0]
  if len(s.nome) == 0 {
    return
  }
  for i:=1; i<len(x); i++ {
    voto, err := strconv.Atoi(x[i])
    if err != nil {
      return 
    }
    if voto < 18 || voto >30 && voto != 32 {
      return
    }
    s.voti= append(s.voti, voto)
  }
  ok = true 
}
//elabora il file 
func parseFile(file string)  (s []studente, ok bool) {
  f,err := os.Open(file)
  defer f.Close()
  if err != nil{
    return
    fmt.Fprintf(os.Stderr, "")
  }
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    riga  := scanner.Text()
  studente, ok := parseLine(riga)
  if !ok {
    fmt.Fprint(os.Stderr, )
    continue
  }
    s = append(s, studente)
  }
  return s , true
}
//crea il file di output
func writeOut(file string, s []studente) (ok bool) {
  f, err := os.Create(file)
  if err != nil {
    return
  }
  for _, studente := range s {
    fmt.Fprintf(f, "%s\t%.2f\n",
               studente.nome,
               convert(avg(studente.voti)))
  }
  return
}
func main() {
  fileDiInput := os.Args[1]
  fileDiOutput := os.Args[2]
}