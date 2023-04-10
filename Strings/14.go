package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Ctrl c - Ctrl V No algoritmo anterior kkkkkkkkkk so mudei os sinais
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma sequência numérica: ")
	scanner.Scan()
	Sequencia := scanner.Text()

	// separa os números da entrada em um slice de strings
	nums := strings.Split(Sequencia, " ")

	// verifica se cada string é um número válido e se a sequência é decrescente
	prev := 1000000000
	for _, numStr := range nums {
		num, erro := strconv.Atoi(numStr)
		if erro != nil {
			fmt.Println("A sequência contém caracteres inválidos.")
			return
		}

		if num >= prev {
			fmt.Println("A sequência não é decrescente.")
			return
		}

		prev = num
	}

	fmt.Println("A sequência é decrescente.")
}
