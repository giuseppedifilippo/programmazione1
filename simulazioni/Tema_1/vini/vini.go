package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
)

type  BottigliaVino struct {
	nome string
	anno int
	grado float32
	quantita int
}

func (b BottigliaVino) String() string {
	out := fmt.Sprintf("%s,%d,%.1f,%d", b.nome, b.anno, b.grado, b.quantita)
	pos := strings.Index(out, ".0,")

	if pos != -1 {
		out = out[:pos] + "," + out[pos+3:]
	}

	return out
}

func CreaBottiglia(nome string, anno int, grado float32, quantita int) (BottigliaVino, bool) {
	if nome == "" || anno <= 0 || grado <= 0 || quantita <= 0 {
		return BottigliaVino{}, false
	} else {
		return BottigliaVino{nome, anno, grado, quantita}, true
	}
}

func CreaBottigliaDaRiga(riga string) (BottigliaVino, bool) {
	line := strings.Split(riga, ",")
	if len(line) != 4 {
		return BottigliaVino{},false
	}
	anno, _ := strconv.Atoi(line[1])
	grd, _ := strconv.ParseFloat(line[2], 32)
	grado := float32(grd)
	qnt, _ := strconv.Atoi(line[3])
	return CreaBottiglia(line[0], anno, grado, qnt)
}

func AggiungiBottiglia(bot BottigliaVino, cantina *[]BottigliaVino) {
	*cantina = append(*cantina, bot)
}

func AggiungiBottigliaDaRiga(bot string, cantina *[]BottigliaVino) {
	bottiglia, ok := CreaBottigliaDaRiga(bot)
	if ok {
		AggiungiBottiglia(bottiglia, cantina)
	}
}

func ContaPerTipo(nome string, cantina []BottigliaVino) int {
	count := 0
	for _, tipo := range cantina {
		if tipo.nome == nome {
			count++
		}
	}

	return count
}

func main() {
	cantina := []BottigliaVino{}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("file non trovato")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if scanner.Text() != "" {
			AggiungiBottigliaDaRiga(scanner.Text(), &cantina)
		}
	}

	for _, b := range cantina {
		fmt.Println(b.String())
	}
}