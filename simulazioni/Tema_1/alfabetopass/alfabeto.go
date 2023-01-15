package main

import (
	"fmt"
	"os"
	"strings"
)

type Dizionario map[string]string

func main() {
	sl := strings.Split(strings.ToLower(os.Args[1]),"")
	dizionario := make(Dizionario)
	for _,el := range sl {
		_, ok := dizionario[el]
		if !ok {
			var temp string
			fmt.Printf("%s? ", el)
			fmt.Scanf("%s\n", &temp)
			dizionario[el] = temp
		}
	}
	for _, el := range sl {
		fmt.Printf("- %s ", dizionario[el])
	}
	fmt.Println()
}// il test non parte ma il programma da come output quello attesso