package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Vocabolario map[string]int


func (v Vocabolario) String() {
	lista := []string{}
	for k := range v {
		lista  = append(lista, k)
	}
	sort.Strings(lista)
	for _, el := range lista {
		fmt.Printf("%s : %d\n", el , v[el])
	}
}

func selezione(voc Vocabolario, char rune) []string {
	lista := []string{}
	for k := range voc {
		if k[0] == byte(char) {
			lista = append(lista, k )
		}
	}
	return lista 
}

func statistiche(voc Vocabolario) (int,int, float64) {
	var max, min, tot  int 
	first := true 
	for _, v := range voc {
		tot += v 
		if first {
			max , min = v, v 
			first = false 
		}else {
			if v > max {
				max = v 
			}
			if v < min {
				min = v 
			}
		}

	}
	return max, min, float64(tot)/float64(len(voc))
}

func main() {
	input , char := os.Args[1], []rune(os.Args[2])[0]
	f ,err := os.Open(input)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var voc Vocabolario
	voc = make(Vocabolario)
	for scanner.Scan() {
		sl := strings.Split(strings.Trim(strings.ReplaceAll(scanner.Text(), "	", " "), " "), " ")
		//fmt.Println(sl)
		for _ , el := range sl {
			if len(el) != 0 {
			voc[el]++
			}
		}
	}
	//delete(voc, "	")
	voc.String()
	max, min ,avg := statistiche(voc)
	fmt.Printf("%d %d %g\n", max, min, avg)
	fmt.Println(selezione(voc, char))
	//fmt.Print(char)
	//fmt.Println(selezione(voc, char))
}//output come previsto ma il test non si avvia e non so il perchè