package main

import (
	"errors"
	"fmt"
	"math"
)

type Punto struct {
	x,y float64
}

type Triangolo struct {
	P1,P2,P3 Punto
}
func distanzaPunti(p1,p2 Punto) (dist float64) {
	dxsq := math.Pow(p2.x - p1.x, 2)
	dysq := math.Pow(p2.y - p1.y, 2 )
	dist = math.Sqrt(dxsq + dysq)
	return dist 
}
func lunghezzeLati(A,B,C Punto) [3]float64 {
	var arr [3]float64
	arr[0] = distanzaPunti(A,B)
	arr[1] = distanzaPunti(B,C)
	arr[2] = distanzaPunti(C,A)
	return arr
}

func newTriangolo(A,B,C Punto) (Triangolo, error) {
	triangolo := Triangolo{}
	arr := lunghezzeLati(A,B,C)
	max := arr[0]
	var sum float64
	for _, el := range arr {
		if el > max {
			max = el
		}
		sum += el
	}
	sum -= max
	if math.Abs(max - sum) < 1e-6   {
		return triangolo, errors.New("non è un triangolo")
	}else {
		triangolo = Triangolo{P1:A, P2:B, P3:C }
		return triangolo, nil
	}
}

func tipoTriangolo(triangolo Triangolo) int {
	arr := lunghezzeLati(triangolo.P1, triangolo.P2, triangolo.P3)
	var count int 
	for i, el := range arr {
		for j, el2 := range arr {
			if i< j && math.Abs(el-el2) < 1e-6 {
				count++
			}
		}
	}
	if count ==1 {
		return 2
	}
	
	return count
}

func main() {
	var A,B,C Punto 
	fmt.Scanf("%f %f\n", &A.x, &A.y)
	fmt.Scanf("%f %f\n", &B.x, &B.y)
	fmt.Scanf("%f %f\n", &C.x, &C.y)
	triangolo, err := newTriangolo(A,B,C)
	fmt.Printf("%g\n", lunghezzeLati(A,B,C))
	if err != nil {
		fmt.Println("non è un triangolo")
	} else {
		fmt.Printf("triangolo %g\ntipo %d\n", triangolo, tipoTriangolo(triangolo))
	}
}//passato