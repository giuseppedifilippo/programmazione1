package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func greater(m, n int) int {
	if m > n {
		return m
	} else {
		return n
	}
}

func recursiveMax(list []int) int {
	if len(list) == 1 {
		return (list)[0]
	}
	max := []int{greater(list[0], list[1])}
	max = append(max, list[2:]...)
	list = max
	return recursiveMax(list)
}

func main() {
	var list []int
	list = make([]int, 0 )
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		list = append(list, n)
		fmt.Println(list)
	}
	fmt.Println(recursiveMax(list))
}//fatto 
