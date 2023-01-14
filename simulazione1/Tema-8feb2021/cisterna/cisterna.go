package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Cisterna struct {
	b,l,h, livello float64
}

func variazione(cisterna *Cisterna, lt int ) error {
	vol := float64(lt*10e3)
	base := cisterna.b*cisterna.l
	temp := cisterna.livello + vol/base
	if cisterna.h - temp < 0 || cisterna.h + temp > cisterna.h {
		return errors.New("errore di input")
	}else {
		cisterna.livello += vol/base
		return nil 
	}
}

func contenuto(cisterna Cisterna) ( litri int) {
	return int(cisterna.b*cisterna.l*cisterna.livello*10e-3)
}

func (cisterna *Cisterna) String () {
	fmt.Printf("cisterna %f cm x %f cm x %f cm\n livello attuale: %f, litri %d\n", cisterna.b, cisterna.l, cisterna.h,cisterna.livello,contenuto(*cisterna))
}

func main() {
	var cisterna Cisterna
	var err error
	cisterna.b, err = strconv.ParseFloat(os.Args[1],64)
	cisterna.l ,err = strconv.ParseFloat(os.Args[2],64)
	cisterna.h , err = strconv.ParseFloat(os.Args[3], 64)
	if err!= nil {
		os.Exit(2)
	}
	var lt int 
	for {
		cisterna.String()
		fmt.Print(">>>")
		fmt.Scan(&lt)
		if lt == 9999 {
			break
		}
		err1 := variazione(&cisterna,lt)
		if err1 != nil {
			fmt.Println("palle")
		}
	}
}