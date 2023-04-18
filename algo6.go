package main

import (
	"fmt"
	"strings"
)

func main() { //Escreva um programa que receba uma string e conte quantas palavras ela contém. Informe ao usuário o resultado.
	var str string
	fmt.Println("Digite uma frase:")
	fmt.Scanln(&str)

	num := strings.Fields(str)
	largura := len(num)
	fmt.Printf("A frase tem %s palavra(s).", largura)
	//não sei resolver

}
