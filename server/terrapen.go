package main

/*
per compilare
go mod init cartella
go mod tidy
(solo una volta)
*/
import (
	"fmt"
	"image"
	"image/png"
	"math"
	"net/http"

	"github.com/holizz/terrapin"
)
func dragon(t *terrapin.Terrapin, len float64, liv int ) {
	if  liv == 0 {
		t.Forward(len)
		return
	}
	var c int 
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
	t := terrapin.NewTerrapin(campo, terrapin.Position{250.0,450.0})
	dragon(t, 100.0, 10)
	png.Encode(w, campo)
}
func palle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<body style = "background-color:white;">
	<doctype html>
	<h1 style = "color:black;">
	<title> sto cazzo</title>
	
	<p>Ciao, merda
	
	`))
}
func main(){
	fmt.Println("Listening on http://localhost:3000")
	http.HandleFunc("/main", stampa)
	http.ListenAndServe(":3000", nil)
}
