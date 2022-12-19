package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)
type Carrello struct {
	pos, carico int 
}
const MAX = 15
func (c Carrello) String() string {
	return fmt.Sprintf("posizione %d, carico %d\n", c.pos, c.carico)
}

func aggiornaStato(c *Carrello, posizione, carico int ) bool {
	if posizione <0 || carico < 0 {
		return false 
	}
	c.pos = posizione
	c.carico += carico 
	return true 
}

func parseString(s string) []int {
	out := []int{}
	s = strings.ReplaceAll(s, " ", "0")
	sl := strings.Split(s, "|")
	for i:=0; i<len(sl); i++ {
		temp, _:= strconv.Atoi(sl[i])
		 out = append(out, temp)
	} 
	return out 
}

func reverseParse(c Carrello, sl []int) {
	fmt.Print("|")
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
	//gestione input 
	listaPesi := map[int]int{}
	var count, peso_max  int 
	sl := parseString("| | | |12|4| | | |10| | | | |4| | | | |5| |12| | | | |3| |")
	var c Carrello 
	for c.pos = 0 ; c.pos <len(sl)-1 ; c.pos++ {
		if sl[c.pos] !=0 {
			if sl[c.pos] > peso_max {
				peso_max = sl[c.pos]
			}
			if c.carico+ sl[c.pos]< MAX {
				aggiornaStato((&c), c.pos, sl[c.pos])
				listaPesi[sl[c.pos]]++
				sl[c.pos]=0
			} else {
				reverseParse(c, sl)
				sl[0]+= c.carico
				c.carico = 0 
				aggiornaStato((&c), 1, 0 )
				count++
			}
		}
	}
	reverseParse(c, sl)
	sl[0]+= c.carico
	c.carico = 0 
	aggiornaStato((&c), 1, 0 )
	reverseParse(c, sl)
	fmt.Printf("n viaggi: %d\n", count)
	fmt.Printf("Peso max : %d\n", peso_max)
	list := []int{}
	for k  := range listaPesi {
		list= append(list, k)
	}
	sort.Ints(list)
	for _, el := range list {
		fmt.Printf("%d ogg. di peso %d\n", listaPesi[el] , el)
	}
}