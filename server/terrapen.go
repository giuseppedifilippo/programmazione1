package main

/*
per compilare
go mod init cartella
go mod tidy
(solo una volta)
*/
import (
	"image"
	"math"
	"net/http"

	"github.com/holizz/terrapin"
)
func dragon(t *terrapin.Terrapin, len float64, liv int ) {
	if  liv == 0 {
		t.Forward(len)
		return
	}
	c :=0
	dragon(t, len, liv -1)
	if c%2==0 {
		t.Right(math.Pi / 2 )
		t.Forward(len)
		c= 1 
	} else {
		t.Left(math.Pi / 2)
		t.Forward(len)
		c = 0 
	}

}
func stampa(w http.ResponseWriter, r *http.Request) {
	campo := image.NewRGBA(image.Rect(0,0,500,500))
	t := terrapin.NewTerrapin(campo, terrapin.Position{250.0,450,0})
	
}
