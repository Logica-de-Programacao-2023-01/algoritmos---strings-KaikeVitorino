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

	str = strings.ReplaceAll(str, "a", "")
	str = strings.ReplaceAll(str, "e", "")
	str = strings.ReplaceAll(str, "i", "")
	str = strings.ReplaceAll(str, "o", "")
	str = strings.ReplaceAll(str, "u", "")
	str = strings.ReplaceAll(str, "A", "")
	str = strings.ReplaceAll(str, "E", "")
	str = strings.ReplaceAll(str, "I", "")
	str = strings.ReplaceAll(str, "O", "")
	str = strings.ReplaceAll(str, "U", "")
	fmt.Printf(str)
}
