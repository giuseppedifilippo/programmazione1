package main

import (
	"bufio"
	"fmt"
	"os"
)
type condizione struct {
	timestamp string 
	temp float64
	hum float64
}
func parseFile(filename string) (sl []condizione) {
	sl = make([]condizione, 20000)
	f, err := os.Open(filename) 
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "errore di apertura del file %s\n", filename)
	}
	os.Stdin = f 
	scanner := bufio.NewScanner(f)
	for i:=0 ; scanner.Scan();i++ {
		fmt.Scanf("%s,%f,%f", &sl[i].timestamp,&sl[i].temp,&sl[i].hum )
	}
	return sl
}

func tempHandle(sl []condizione) (min,max condizione ) {
	mint, maxt := sl[0], sl[0]
	for  i:=1; i<len(sl); i++ {
		if sl[i].temp > maxt.temp {
			maxt = sl[i]
		}
		if sl[i].temp < mint.temp  {
			mint = sl[i]
		}
	}
	return 
}

func humHandle(sl []condizione) (min,max condizione ) {
	minh, maxh := sl[0], sl[0]
	for  i:=1; i<len(sl); i++ {
		if sl[i].hum > maxh.hum {
			maxh = sl[i]
		}
		if sl[i].hum < minh.hum  {
			minh = sl[i]
		}
	}
	return 
}

func main() {
	slice := parseFile("temp.csv")
	minTemp,  maxTemp := tempHandle(slice)
	minHumid, maxHumid := humHandle(slice)
	fmt.Print(minTemp, maxTemp, minHumid, maxHumid)// manca da sistemare il timestamp 
}//dovevo usare < per la ridirezione 
//non aprire il file dal programma
