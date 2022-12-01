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
func pippoFunc(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(`
  <body style = "background-color:white;">
<!doctype html>
<h1 style = "color:black;"> 
<title> pagina di saluto</title>

<p>Ciao, coglione 

<p><img src=""https://www.google.com/imgres?imgurl=https%3A%2F%2Fi.kym-cdn.com%2Fphotos%2Fimages%2Fnewsfeed%2F001%2F858%2F345%2Fbcc&imgrefurl=https%3A%2F%2Fknowyourmeme.com%2Fmemes%2Ffemboy-hooters&tbnid=y7VG_7Tu-_rDXM&vet=12ahUKEwja64nFiNb7AhVBIMUKHZIzBHUQMygAegUIARCXAQ..i&docid=OejfLM2Bl_29uM&w=680&h=428&q=femboy%20hooters&ved=2ahUKEwja64nFiNb7AhVBIMUKHZIzBHUQMygAegUIARCXAQ">
  `))
}
func Heighway
func mainImage(w http.ResponseWriter, r *http.Request) {
  campo := image.NewRGBA(image.Rect(0,0,500,500))
  t := terrapin.NewTerrapin(campo, terrapin.Position{250.0,450.0})
  t.Forward(100)
  t.Right(math.Pi / 2.0)
  t.Forward(100)
  t.Right(math.Pi / 2.0)
  t.Forward(100)
  t.Right(math.Pi / 2.0)
  t.Forward(100)
  t.Right(math.Pi / 2.0)
  
  png.Encode(w, campo)

  
}
func main() {
  fmt.Println("Listening on http://localhost:3000")
  http.HandleFunc("/pippo", pippoFunc)
  http.HandleFunc("/main", mainImage)
  http.ListenAndServe(":3000", nil)
}