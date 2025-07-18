package main

import (
	"fmt"
	"math"
)

func main() {
	x1,x2 := bhaskara(1, -5, 6)
	fmt.Printf("x1= %.2f | x2= %.2f", x1, x2)
	
	total := sum(1, 4, 2.4, 6, 2, 7.6, 11, 43.3, 9)
	fmt.Printf("Total: %.2f", total)
}

//Go permite criar funções com múltiplos retornos
func bhaskara(a float64, b float64, c float64) (float64, float64) {
	//x = (-b ± √(b² - 4ac)) / 2a
	delta := delta(a, b, c)
	
	x1 := ((-b + math.Sqrt(delta)) / (2*a))
	x2 := ((-b - math.Sqrt(delta)) / (2*a))

	return x1, x2
}

func delta(a float64, b float64, c float64) float64 {
	delta := math.Pow(b, 2) - 4*a*c
	return delta
}

//Função Variádica, permite estabelecer uma função com número de parâmetros indefinido
func sum(numbers ...float64) float64 {
	result := 0.0
	for _, number := range numbers {
		result += number
	}

	return result
}
