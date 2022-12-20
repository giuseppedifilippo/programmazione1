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
func (c Carrello) String() string {
	return fmt.Sprintf("carrello: posizione %d, carico %d\n", c.pos, c.carico)
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
	sl = sl[1:len(sl)-1]
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

func azzerra(sl []int, end int) {
	for i:=1; i <= end ; i++ {
		sl[i] = 0 
	} 
}


func main() {
	listaPesi := map[int]int{}
	first := true 
	var count, peso_max  int 
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("manca nome file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	
	//gestione input 
	
	sl := parseString(scanner.Text())//"| | | |12|4| | | |10| | | | |4| | | | |5| |12| | | | |3| |"
	var c Carrello 
	for c.pos = 0 ; c.pos <len(sl)-2 ; c.pos++ {
		if sl[c.pos+1] !=0 {
			if sl[c.pos+1] > peso_max {
				peso_max = sl[c.pos+1]
			}
			
			if c.carico+ sl[c.pos+1]<= MAX {
				listaPesi[sl[c.pos+1]]++
				aggiornaStato((&c), c.pos , sl[c.pos+1])
				sl[c.pos] = 0 
			} else {
				if first {
					reverseParse(c, sl)
					azzerra(sl, c.pos)
					first = false 
				}
				sl[0]+= c.carico
				reverseParse(c, sl)
				azzerra(sl, c.pos)
				c.carico = 0 
				aggiornaStato((&c), 1, 0 )
				count++
			}
			
		}
	}	
	//aggiornaStato((&c), c.pos+1, 0)
	c.pos += 1
	reverseParse(c, sl)
	sl[0]+= c.carico
	c.carico = 0 
	azzerra(sl, c.pos)
	aggiornaStato((&c), 0, 0 )
	reverseParse(c, sl)
	fmt.Printf("n viaggi: %d\n", count+1)
	fmt.Printf("peso max : %d\n", peso_max)
	list := []int{}
	for k  := range listaPesi {
		list= append(list, k)
	}
	sort.Ints(list)
	for _, el := range list {
		fmt.Printf("%d ogg. di peso %d\n", listaPesi[el] , el)
	}
}