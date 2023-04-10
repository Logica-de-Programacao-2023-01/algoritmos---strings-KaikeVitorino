package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite uma string: ")
	scanner.Scan()
	str := scanner.Text()

	fmt.Print("Digite o caractere a ser substituído: ")
	scanner.Scan()
	Caractere1 := scanner.Text()

	fmt.Print("Digite o novo caractere: ")
	scanner.Scan()
	Caractere2 := scanner.Text()

	newStr := strings.ReplaceAll(str, Caractere1, Caractere2)

	fmt.Printf("String resultante: %s\n", newStr)
}
