package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite a primeira string: ")
	scanner.Scan()
	str1 := scanner.Text()

	fmt.Print("Digite a segunda string: ")
	scanner.Scan()
	str2 := scanner.Text()

	if strings.Contains(str1, str2) {
		fmt.Println("A segunda string é uma substring da primeira.")
	} else {
		fmt.Println("A segunda string não é uma substring da primeira.")
	}
}
