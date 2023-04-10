package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var n int

	fmt.Print("Digite uma string: ")
	fmt.Scan(&str)

	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	if n > len(str) {
		fmt.Println("Número inválido!")
		return
	}
	str = strings.ToUpper(str[:n]) + str[n:]

	fmt.Println("String resultante:", str)
}
