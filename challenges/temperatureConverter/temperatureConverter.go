package main

import "fmt"

func main() {
	var optInput int
	var degrees, convertedDegrees float64
	fmt.Println("Selecione a medida inicial para temperatura:")
	fmt.Println("1- Celsius \n2- Fahrenheit \n3- Kelvin")
	fmt.Scanln(&optInput)
	fmt.Println("Digite os graus")
	fmt.Scanln(&degrees)
	fmt.Println("Para qual medida você deseja converter?")
	fmt.Println("1- Celsius \n2- Fahrenheit \n3- Kelvin")
	fmt.Scanln(&optInput)
	fmt.Println("-- RESULTADO ", convertedDegrees, "--")
}
