package main

import "fmt"

func main() {
	a, b := 0, 0

	a = 2
	b = 4
	if a > b {
		fmt.Println("A maior que B")
	} else if a < b {
		fmt.Println("A menor que B")
	} else {
		fmt.Println("A é igual a B")
	}

	if shouldEscape('@') {
		fmt.Println("Caractere permitido")
	} else {
		fmt.Println("Caractere não permitido")
	}
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&':
		return true
	}
	return false
}
