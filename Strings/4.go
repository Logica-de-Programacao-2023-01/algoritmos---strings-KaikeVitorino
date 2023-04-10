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
	str1 := scanner.Text()

	fmt.Print("Digite uma string: ")
	scanner.Scan()
	str2 := scanner.Text()

	if str1 == str2 {
		fmt.Printf("Strings iguais")
	} else {
		fmt.Printf("String diferentes")
	}

}
