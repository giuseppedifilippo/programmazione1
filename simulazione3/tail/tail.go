package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, filename := os.Args[1] , os.Args[2]
	N, _ := strconv.Atoi(n)
	f ,err := os.Open(filename)
	if err != nil {
		fmt.Println("file error")
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	list := []string{}
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	for i:= len(list)-N; i<=len(list)-1; i++ {
		fmt.Println(list[i])
	}
}//funziona