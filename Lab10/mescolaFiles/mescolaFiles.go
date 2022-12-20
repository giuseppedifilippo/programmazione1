package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var rem []string
	if len(os.Args) != 3 {
		fmt.Println("Inserire DUE nomi di file")
		os.Exit(1)
	}
	file1, file2 := os.Args[1], os.Args[2]
	f1, err1 := os.Open(file1)
	f2, err2 := os.Open(file2)
	if err1 != nil || err2 != nil {
		fmt.Println("inserire DUE nomi di file")
	}
	defer f1.Close()
	defer f2.Close()
	scanner1 := bufio.NewScanner(f1)
	scanner2 := bufio.NewScanner(f2)
	scanner1.Scan()
	sl1 := strings.Split(scanner1.Text(), " ")
	scanner2.Scan()
	sl2 := strings.Split(scanner2.Text(), " ")
	if len(sl1) != len(sl2) {
		if len(sl1) > len(sl2) {
			r := len(sl1) - len(sl2)
			rem = sl1[r+1:]
			sl1 =sl1[:r+1]
		} else {
			r := len(sl2) - len(sl1)
			rem = sl2[r+1:]
			sl2 =sl2[:r+1]
		}
	}
	//fmt.Println(sl1)
	for i, el := range sl1 {
		fmt.Println(el)
		fmt.Println(sl2[i])
	}
	//fmt.Println(sl2[len(sl2)-1])
	for _, el := range rem {
		fmt.Println(el)
	}
}//fatto