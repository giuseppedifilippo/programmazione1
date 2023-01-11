/**
 * "libreria" di test per gli esami, non modificare questo file!
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func lanciaGenerica(
	t *testing.T,
	nomeProg string,
	inputString string,
	expectedOutString string,
	args ...string) {

	fmt.Println("### Questo test verifica output atteso!")
	fmt.Println("### Presuppone che l'eseguibile da testare (", nomeProg, ") sia già stato compilato!")

	subproc := exec.Command(nomeProg, args...)
	subproc.Stdin = strings.NewReader(inputString)
	stdout, err := subproc.CombinedOutput()

	if err != nil {
		// verificare se uscito per esplicito os.Exit()
		//t.Errorf("Fallito: %s\n", stderr)

		//fmt.Printf(">>> Attenzione! Uscito con errore: %s\n>>> (non è un test fallito se l'esercizio richiedeva uscita esplicita con exit code)\n", err)
		fmt.Printf(">>> Attenzione! Uscito con errore: %s\n>>> (non è un test fallito se si termina il programma con un esplicito os.Exit\n", err)
	}

	fmt.Printf("+++ Args:\n%s\n", args)
	fmt.Printf("+++ Input:\n%s\n", inputString)
	fmt.Printf("+++ Output:\n%s\n", string(stdout))
	fmt.Printf("+++ Expected output:\n%s\n", expectedOutString)
	if string(stdout) != expectedOutString {
		fmt.Println(">>> FAIL!")
		t.Fail()
	}

	subproc.Process.Kill()
}

func lanciaGenericaConFileOutAtteso(
	t *testing.T,
	nomeProg string,
	inputString string,
	expectedOutFilename string,
	args ...string) {

	content, err := ioutil.ReadFile(expectedOutFilename)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)
	//fmt.Println(text)

	lanciaGenerica(t, nomeProg, inputString, text, args...)
}
