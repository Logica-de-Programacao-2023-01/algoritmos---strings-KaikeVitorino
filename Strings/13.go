package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma sequência numérica: ")
	scanner.Scan()
	input := scanner.Text()

	// separa os números da entrada em um slice de strings
	nums := strings.Split(input, " ")

	// verifica se cada string é um número válido e se a sequência é crescente
	prev := -1
	for _, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("A sequência contém caracteres inválidos.")
			return
		}

		if num <= prev {
			fmt.Println("A sequência não é crescente.")
			return
		}

		prev = num
	}

	fmt.Println("A sequência é crescente.")
}
