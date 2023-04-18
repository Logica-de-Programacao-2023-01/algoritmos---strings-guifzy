package main

import (
	"fmt"
	"strconv"
)

func main() { //Escreva um programa que receba uma string e verifique se ela é um número válido em ponto flutuante.
	// Informe ao usuário se é ou não.
	var str string
	fmt.Println("Digite um número:")
	fmt.Scan(&str)

	_, err := strconv.ParseFloat(str, 64)
	if err == nil {
		fmt.Println("É um número flutuante")
	} else {
		fmt.Println("O valor não é um número válido em ponto flutuante")
	}

}
