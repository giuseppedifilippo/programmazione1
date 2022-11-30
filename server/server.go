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
	"net/http"
  "math"
	"github.com/holizz/terrapin"
  "image/png"
)
func pippoFunc(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte(`
<!doctype html>
<title> pagina di saluto</title>

<p>Ciao, coglione 

<p><img src=""/mainImage.png">
  `))
}
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