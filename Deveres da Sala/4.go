package main

import (
	"fmt"
	"strings"
)

func main() {
	var string1 string
	fmt.Println("Escreva uma string:")
	fmt.Scan(&string1)
	string2 := strings.ToUpper(string1)
	fmt.Print("Sua string dada com todos os caracteres maiusculos: ", string2)
}
