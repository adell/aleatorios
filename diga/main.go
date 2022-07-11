package main

import (
	"fmt"
	"os"
	"strings"
)

var qtd_barra int

func diga(diga string) string {
	qtd_barra = len(diga) + 1
	return diga
}

func main() {
	var fale string
	if len(os.Args) > 1 {
		fale = diga(os.Args[1])
	} else {
		fale = diga("Digite um texto para que eu fale!")
	}
	barra := strings.Repeat("-", qtd_barra)

	fmt.Println(" ", barra)
	fmt.Println(" (", fale, ")")
	fmt.Println(" ", barra)
	fmt.Println("       \\   ^__^")
	fmt.Println("        \\_ (oo)\\_______")
	fmt.Println("           (__)\\       )\\")
	fmt.Println("               ||----w | \\")
	fmt.Println("               ||     ||  s")
}
