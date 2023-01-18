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
func LunghezzeLati(A,B,C Punto) [3]float64 {
	var arr [3]float64
	arr[0] = distanzaPunti(A,B)
	arr[1] = distanzaPunti(B,C)
	arr[2] = distanzaPunti(C,A)
	return arr
}

func newTriangolo(A,B,C Punto) (Triangolo, error) {
	triangolo := Triangolo{}
	if !(LunghezzeLati(A,B,C)[0] < LunghezzeLati(A,B,C)[1] + LunghezzeLati(A,B,C)[2]) {
		return triangolo, errors.New("non è un triangolo")
	}else {
		triangolo = Triangolo{P1:A/*Punto{x:A.x, y:A.y}*/, P2:B/*Punto{x:B.x,y:B.y}*/, P3:C /*Punto{x:C.x, y:}*/}
		return triangolo, nil
	}
}

func tipoTriangolo(triangolo Triangolo) int {
	arr := LunghezzeLati(triangolo.P1, triangolo.P2, triangolo.P3)
	var count int 
	for i:= 0 ; i< 2; i++ {
		for _, el := range arr[i+1:] {
			if math.Abs(arr[i]-el) < 1e-6 {
				count++
			}
		}
	}
	if count == 0 {
		return count
	}
	return count+1
}

func main() {
	var A,B,C Punto 
	fmt.Scanf("%f %f\n", &A.x, &A.y)
	fmt.Scanf("%f %f\n", &B.x, &B.y)
	fmt.Scanf("%f %f\n", &C.x, &C.y)
	triangolo, err := newTriangolo(A,B,C)
	fmt.Printf("l'output è:\n%.2f\n", LunghezzeLati(A,B,C))
	if err != nil {
		fmt.Println("non è un triangolo")
	} else {
		fmt.Printf("triangolo %.2f\ntipo %d\n", triangolo, tipoTriangolo(triangolo))
	}
}