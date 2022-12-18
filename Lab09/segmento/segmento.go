package main

import "fmt"

type Segmento struct {
	ex, in      byte
	orizzontale bool
	l           int
}

func (x Segmento) String() string {
	dis := string(x.ex)
	for i := 0; i < x.l-2; i++ {
		if !x.orizzontale {
			dis += "\n"
		}
		dis += string(x.in)
	}
	if !x.orizzontale{
		dis += "\n"
	}
	dis += string(x.ex)
	return dis
}

func (x Segmento) allunga(n int) {
	x.l += n
}

func main() {
	var s Segmento
	fmt.Scanf("%c %c %t %d", &s.ex, &s.in, &s.orizzontale, &s.l)
	fmt.Println(s.String())
}//fatto