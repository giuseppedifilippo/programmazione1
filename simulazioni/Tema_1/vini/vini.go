package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BottigliaVino struct {
	nome string
	anno int
	grado float32
	quantita int
}

func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool) {
	if nome != "" && anno > 0 && grado > 0 && quantita > 0 {
		return BottigliaVino{nome : nome, anno : anno, grado : grado, quantita : quantita} , true
	}
	return BottigliaVino{} , true
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	sl := strings.Split(riga, ",")
	anno , _ := strconv.Atoi(sl[1])
	grado, _ := strconv.ParseFloat(sl[2],64)
	quantita, _ := strconv.Atoi(sl[3])
	return CreaBottiglia(sl[0],anno, float32(grado),quantita)
}

func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino) {
	*cantina = append(*cantina, bot)
}

func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino) {
	bottiglia , ok := CreaBottigliaDaRiga(bot)
	if  ok {
		*cantina = append(*cantina, bottiglia)
	}
}

func ContaPerTipo(nome string, cantina []BottigliaVino) int {
	var count int
	for _, el := range cantina {
		if el.nome == nome {
			count++
		}
	}
	return count
}

func (bot BottigliaVino) String() string {
	return fmt.Sprintf("%s,%d,%.f,%d", bot.nome,bot.anno, bot.grado,bot.quantita)
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("file error")
		os.Exit(1)
	}
	defer f.Close()
	cantina := []BottigliaVino{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if len(scanner.Text()) <= 1 {
			continue
		}
		bot, ok := CreaBottigliaDaRiga(scanner.Text())
		if ok {
			AggiungiBottiglia(bot,&cantina)
			fmt.Println(bot.String())
		}
	}
}