package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cerchio struct {
	nome         string
	x, y, raggio float64
}

func newCircle(descr string) Cerchio {
	sl := strings.Split(descr, " ")
	nums := []float64{}
	for _, el := range sl[1:] {
		x, _ := strconv.ParseFloat(el, 64)
		nums =append(nums, x)
	}
	return Cerchio{nome:sl[0],x:nums[0],y:nums[1], raggio: nums[2]}
}

func distanzaPunti(x1,y1,x2,y2 float64) float64 {
	deltax := math.Pow(x2-x1, 2)
	deltay := math.Pow(y2-y1,2)
	return math.Sqrt(deltax + deltay)
}

func tocca(cerc1, cerc2 Cerchio) bool {
	if math.Abs(distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y) - (cerc1.raggio+ cerc2.raggio)) < 1e-6 {
		return true
	} else if math.Abs(math.Abs(cerc1.raggio - cerc2.raggio) - distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y)) < 1e-6 {
		return true
	} else {
		return false
	}
}

func boolString(cerc1, cerc2 Cerchio) string {
	touch := tocca(cerc1, cerc2)
	mag := maggiore(cerc1, cerc2)
	switch {
	case touch && mag : return "tangente, maggiore"
	case touch && !mag : return "tangente, minore o uguale"
	case !touch && mag : return "non tangente, maggiore"
	default : return "non tangente, minore o uguale"
	}
}

func maggiore(cerc1, cerc2 Cerchio) bool {
	return cerc1.raggio - cerc2.raggio < 1e-6
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cerc1 := Cerchio{}
	for scanner.Scan() {
		cerc2 := newCircle(scanner.Text())
		fmt.Printf("%v %s\n", cerc2, boolString(cerc1, cerc2))
		cerc1 = cerc2
	}
}//passa tutti i test