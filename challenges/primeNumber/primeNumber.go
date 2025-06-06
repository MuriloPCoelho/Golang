package main

import "fmt"

func main() {
	var input int

	fmt.Println("Digite um número inteiro e descubra se ele é primo ou não:")
	fmt.Scanln(&input)

	//o número é primo se ele só for divisível por 1 ou por ele mesmo

	for i := 2; i <= input; i++ {
		if input%i == 0 && i != input {
			fmt.Println("Não é primo")
			break
		}

		if i == input {
			fmt.Println("É primo")
		}
	}
}
