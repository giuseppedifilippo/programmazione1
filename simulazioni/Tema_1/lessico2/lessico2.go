package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func PrintMenu() {
	fmt.Println("+ (legge e memorizza)")
	fmt.Println("? (numeri di riga in cui è comparsa la parola data)")
	fmt.Println("m (stampa le lunghezze min e max)")
	fmt.Println("p (stampa le parole memorizzate)")
}

func PrintDict(m map[string][]int) {
	keys := []string{}

	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, ":", m[k])
	}
}

func AddToDict(line string, n int, dict map[string][]int) {
	line_arr := strings.Split(line, " ")
	line_arr = line_arr[1:]
	for _, v := range line_arr {
		dict[v] = append(dict[v], n)
	}
}

func Stats(dict map[string][]int) (int, int) {
	keys := []string{}

	for k := range dict {
		keys = append(keys, k)
	}

	min := len(keys[0])
	max := 0
	for _, w := range keys {
		if len(w) < min {
			min = len(w)
		}

		if len(w) > max {
			max = len(w)
		}
	}

	return min, max
}

func main() {
	dict := map[string][]int{}
	PrintMenu()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		switch line[0] {
		case '+':
			row++
			fmt.Println("alimento il dizionario")
			AddToDict(line, row, dict)
			break
		case '?':
			fmt.Println("consulto il dizionario")
			fmt.Println("parola:", line[2:])
			fmt.Println("righe", dict[line[2:]])
			break
		case 'm':
			fmt.Println("lunghezza min e max")
			fmt.Println(Stats(dict))
			break
		case 'p':
			fmt.Println("stampo il dizionario ordinato")
			PrintDict(dict)
			break
		default:
			fmt.Println("opzione non valida")
			break
		}
	}
}
