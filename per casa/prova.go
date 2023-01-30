package main

import "fmt"

type f func(x int)  {
	fmt.Println(x)
}
func main() {
	f(2)
}