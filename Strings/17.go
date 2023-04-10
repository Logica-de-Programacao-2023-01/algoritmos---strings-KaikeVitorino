package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite uma string: ")
	scanner.Scan()
	str := scanner.Text()

	// cria um mapa para contar as ocorrências de cada caractere
	mapa := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		mapa[str[i]]++
	}

	// imprime as letras únicas
	fmt.Println("Letras únicas na string:")
	for i := 0; i < len(str); i++ {
		if mapa[str[i]] == 1 {
			fmt.Printf("%c ", str[i])
			mapa[str[i]] = 0 // marca como já impresso
		}
	}
	fmt.Println()
}
