package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Carrello struct {
	pos, carico int 
}

const MAX = 15 
//stringing del carrello 
func (c Carrello) String() string {
	return fmt.Sprintf("carrello: posizione %d, carico %d\n", c.pos, c.carico)
}

//aggiorna la situazione del carrello con i parametri passati
func aggiornaStato(c *Carrello, posizione, carico int ) bool {
	if posizione < 0 || carico < 0 {
		return false 
	}
	c.pos = posizione
	c.carico = carico 
	return true 
}
//la slice deve avere un elemento in piu alla fine 
//al fine di evitare il panic durante il checking 
//rimuovere l elemento dureante il reverse parse 

//converte la stringa di input in una slice per un iterazione piu facile
func parseString(s string) []int {
	out := []int{}
	s = strings.ReplaceAll(s, " ", "0")
	sl := strings.Split(s, "|")
	sl = sl[1:]
	for i:=0; i<len(sl); i++ {
		temp, _:= strconv.Atoi(sl[i])
		 out = append(out, temp)
	} 
	return out 
}
//porta tutti i valori della slice compresi tra start ed end a zero 
func azzera(sl []int,start, end int ) {
	for i:=start  ; i<= end ; i++ {
		sl[i] = 0 
	}
}

//stampa la slice nel formato richiesto 
func reverseParse(c Carrello, sl[]int) {
	fmt.Print("|")
	sl = sl[:len(sl)-1]
	for _, n := range sl {
		if n != 0 {
			fmt.Print(n,"|")
		} else {
			fmt.Print(" |")
		}
	}
	fmt.Println()
	fmt.Print(c.String())
}


func main() {
	var starting_point int 
	update:= true  
	listapesi := map[int]int{}
	keylist := []int{}
	 var  count ,  peso_max int 
	var c Carrello
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil || len(os.Args)<2  {
		fmt.Println("manca nome file")
		os.Exit(1)
	}
	defer f.Close()
	scanner:= bufio.NewScanner(f)
	scanner.Scan()
	sl:=parseString(scanner.Text())
	for c.pos = 0 ; c.pos < len(sl)-1; c.pos++{
		if sl[c.pos] != 0 {
			
			if update {
				starting_point = c.pos
				update = false 
			}
			listapesi[sl[c.pos]]++//mappa che contiene i carichi e le loro occorrenze 
			if sl[c.pos] > peso_max { // controllo del peso massimo 
				peso_max = sl[c.pos]
			}
			aggiornaStato((&c), c.pos, sl[c.pos]+c.carico)// carico sul carrello il carico nella poszione attuale 
		}
			
		// effettua lo scarico del carrello quando ha raggiunto il massimo che puo trasportare
		if sl[c.pos+1] + c.carico >= 15 {
			count++
			reverseParse(c, sl)
			azzera(sl, starting_point, c.pos)
			sl[0] += c.carico
			aggiornaStato((&c), 1, 0)
			update = true
		}
	}
	//ultima iterazione della 
	
	sl[0] += c.carico
	aggiornaStato((&c), 0 ,0 )
	azzera(sl, starting_point, len(sl)-1)
	reverseParse(c, sl)
	fmt.Printf("n viaggi: %d\n", count+1)
	fmt.Printf("peso max: %d\n", peso_max)
	fmt.Println("oggetti trovati:")
	for k := range listapesi {
		keylist = append(keylist, k)
	}
	sort.Ints(keylist)
	for _, el := range keylist {
		fmt.Printf("%d ogg. di peso %d\n", listapesi[el], el)
	}
}
// il programma funziona alla perfezione ma i test sono fatti col culo 
//basta mi arrendo
//fixo un problema e ne esce un altro .
//oddio come lo faccio io 
//funziona de cristo 
//ma sti kitammuorto di test non passano per l arcicazzo
//trentini morpungo siete bravi oddio 
//ma cristo sti test sono scritti buttando la fava sulla tastiera 