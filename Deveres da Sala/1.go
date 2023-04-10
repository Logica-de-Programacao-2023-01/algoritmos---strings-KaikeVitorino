package main

import "fmt"

func main() {
	var s1 string
	var s2 string
	fmt.Println("Digite uma string(SOMENTE UMA)")
	fmt.Scan(&s1)
	fmt.Println("Digite outra string(SOMENTE UMA)")
	fmt.Scan(&s2)
	s3 := s1 + " " + s2
	fmt.Println(s3)
}
