package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[2]
	n, _ := strconv.Atoi(os.Args[1])
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("file error")
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	sl := []string{}
	for scanner.Scan() {
		sl = append(sl, strings.Split(strings.ReplaceAll(scanner.Text(),"\n", " ")," ")...)
	}
	for _, el := range sl {
		var str string
		if len(str)+len(el) < n {
			fmt.Println(str)
			str = ""
		}
		str += el
	} 

	
}//non so come fare i test .sh