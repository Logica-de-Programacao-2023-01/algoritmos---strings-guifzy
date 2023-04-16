package main

import (
	"fmt"
	"strings"
)

func main() { //Escreva um programa que receba uma string e substitua todas as ocorrências desse caractere
	// na string por outro caractere especificado pelo usuário. Informe ao usuário o resultado.
	var str, subs, novaStr string

	fmt.Print("Digite nomes: ")
	fmt.Scanln(&str)

	fmt.Print("Digite o some a ser removido: ")
	fmt.Scanln(&subs)

	fmt.Print("Digite o nome a ser adicionado: ")
	fmt.Scanln(&novaStr)

	replace := strings.ReplaceAll(str, subs, novaStr)
	fmt.Println("Novos nomes:", replace)
}
