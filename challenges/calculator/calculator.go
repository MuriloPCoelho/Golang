package main

import (
	"fmt"
)

/*
OPERAÇÕES
	1- Somar
	2- Substrair
	3- Multiplicar
	4- Dividir
	5- Resultado
*/

func main() {
	var opInput int
	var n1, n2, result float64

	fmt.Println("Digite dois números")
	fmt.Scanln(&n1, &n2)

	fmt.Println("-- OPERAÇÕES --\n 1- Somar\n 2- Substrair\n 3- Multiplicar\n 4- Dividir\n 5- Resultado")
	fmt.Println("\nQual operação você deseja realizar? ")
	fmt.Scanln(&opInput)

	switch opInput {
	case 1:
		result = n1 + n2
	case 2:
		result = n1 - n2
	case 3:
		result = n1 * n2
	case 4:
		result = n1 / n2
	}

	fmt.Printf("\n--RESULTADO: %f-- ", result)
}
