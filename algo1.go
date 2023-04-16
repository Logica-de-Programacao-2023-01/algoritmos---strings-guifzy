package main

import (
	"fmt"
	"strings"
)

func main() { //Escreva um programa que receba uma string e converta todas as
	// letras minúsculas para maiúsculas. Informe ao usuário o resultado.
	var nome string
	fmt.Println("Digite um nome em letra minusculas:")
	fmt.Scan(&nome)

	maiusculo := strings.ToUpper(nome)
	fmt.Println("Seu nome em maiusculo: ", maiusculo)
}
