package main

import "fmt"

func main() {
	var s1 string
	fmt.Println("Digite uma string(SOMENTE UMA)")
	fmt.Scan(&s1)
	s2 := len(s1)
	fmt.Printf("Numero de caracteres: %d", s2)

}
