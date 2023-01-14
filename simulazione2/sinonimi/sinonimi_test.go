/*

Sinonimi
========

Si scriva un programma (il file deve chiamarsi 'sinonimi.go') che gestisca come spiegato qui di seguito l'implementazione di un dizionario dei sinonimi da applicare ad un testo.

Il programma dovrà lavorare su due file di testo, i cui nomi verranno passati a riga di comando, il primo conterrà il dizionario dei sinonimi, il secondo il testo da elaborare (privo di punteggiatura).

Il file del dizionario dei sinonimi conterrà un elenco di parole corredate del proprio insieme di sinonimi seguendo la seguente struttura:

	parola: sinonimo1, sinonimo2, sinonimo3, ...

Nota: si può assumere che la struttura delle righe del dizionario dei sinonimi sia sempre questa senza fare nessun controllo

Il file da elaborare è un semplice testo su più righe, privo di punteggiatura.

Il programma 'sinonimi.go' dovrà accettare come primo parametro il nome del file dei sinonimi e come secondo parametro il nome del file da elaborare.
In caso di mancanza di uno dei due parametri il programma dovrà stampare "2 file names, please".
In caso di assenza di uno dei due file (i cui nomi sono stati passati a linea di comando) il programma dovrà stampare rispettivamente "file 1 not found" e "file 2 not found".

L'elaborazione consiste nel produrre un nuovo testo in cui tutte le parole che possiedono un sinonimo (secondo il file dei sinonimi) vengono arricchite dei loro sinonimi racchiusi tra parentesi quadre e in ordine alfabetico, ad esempio:

$ ./sinonimi sinonimi.text promessi.input

produce l'output (riportato parzialmente, solo a titolo esemplificativo):

Quel ramo del lago[pozza specchio stagno] di Como che volge a mezzogiorno tra due catene non interrotte di monti tutto a seni e a golfi a seconda dello sporgere e del rientrare di quelli vien quasi a un tratto[pezzo porzione segmento spezzone] a ristringersi e a prender corso[cammino percorso tragitto] e figura di fiume tra un promontorio[capo punta] a destra e ampia costiera dall altra parte e il ponte che ivi congiunge le due rive par che renda ancor più sensibile all occhio questa trasformazione e segni il punto in cui il lago[pozza specchio stagno] cessa e l Adda rincomincia per ripigliar poi nome di lago[pozza specchio stagno] dove le rive allontanandosi di nuovo lascian l acqua distendersi e rallentarsi in nuovi golfi e in nuovi seni
...


*/

package main

import "testing"

var prog = "./sinonimi"

func TestEsBreve(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"es.output",
		"sinonimi_es.txt",
		"es.input")
}

func TestPromessi(t *testing.T) {
	LanciaGenericaConFileOutAtteso(
		t,
		prog,
		"NIENTE",
		"promessi.output",
		"sinonimi.text",
		"promessi.input")
}

func TestMainNoArgs(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"2 file names, please\n")
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"2 file names, please\n",
		"uno_solo")
}

func TestMainFileNotFound(t *testing.T) {
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"file 1 not found\n",
		"disicurononce",
		"sinonimi.text")
	LanciaGenerica(
		t,
		prog,
		"NIENTE",
		"file 2 not found\n",
		"sinonimi.text",
		"disicurononce")
}
