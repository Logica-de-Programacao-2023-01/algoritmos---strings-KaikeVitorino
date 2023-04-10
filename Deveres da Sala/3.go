package main

import "fmt"

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	var char rune
	fmt.Print("Digite um caractere: ")
	fmt.Scanf("%c", &char)
	count := 0
	for _, c := range str {
		if c == char {
			count++
		}
	}
	fmt.Printf("O caractere '%c' ocorre %d vezes na string '%s'.\n", char, count, str)
}
