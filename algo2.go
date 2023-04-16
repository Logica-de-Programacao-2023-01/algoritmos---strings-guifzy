package main

import (
	"fmt"
	"strings"
)

func main() { //Escreva um programa que receba uma string e remova todos os
	// espaços em branco. Informe ao usuário o resultado.
	var str string
	fmt.Println("Digite um texto: ")
	fmt.Scan(&str)

	x := strings.ReplaceAll(str, " ", "")
	fmt.Println("Novo texto sem espaços em branco:", x)
}
