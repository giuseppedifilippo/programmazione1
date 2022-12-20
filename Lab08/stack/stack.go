package main

import (
	"bufio"
	"fmt"
	"os"
)

//prepend di un valore alla stack
func Push(x *[]float64) {
	var n float64 
	fmt.Println("valore?")
	fmt.Scan(&n)
	s := []float64{n}
	*x = append(s, *x...)
	fmt.Println(*x)
}

//rimuove il valore in testa alla stack e lo restituisce 
func Pop(x *[]float64) {
	fmt.Println("testa", (*x)[0])
	*x = (*x)[1:]
	fmt.Println(*x)
}
func Top(x *[]float64) {
	fmt.Println("testa", (*x)[0])
	fmt.Println(*x)	
}


//controlla se la stack è vuota 
func Empty(x *[]float64) {
	fmt.Println(len(*x)==0)
}


func main() {
	var stack []float64
	stack = make([]float64, 0 )
	scanner := bufio.NewScanner(os.Stdin)
	myloop :
	for {
		fmt.Println("Operazione (push/pop/top/empty/quit)?")
		scanner.Scan()
		switch scanner.Text() {
			case "push": Push(&stack)
			case "pop":  Pop(&stack) 	
			case "top": Top(&stack)	
			case "empty": Empty(&stack)	
			case "quit": break myloop
			default : fmt.Println("inserimento incoretto")
		}
	}
}//fatto