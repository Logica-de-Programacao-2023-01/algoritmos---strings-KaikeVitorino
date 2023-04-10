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

	Palavras := strings.Fields(str)
	fmt.Printf("Essa string tem %d palavras", len(Palavras))
}
