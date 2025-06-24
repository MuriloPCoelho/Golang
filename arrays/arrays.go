package main

import "fmt"

func main() {
	// Arrays precisam ter tamanho definido
	var cars [5]string
	// Depois de criado, o array "aloca" a quantidade de espaços definidos na sua inicialização
	cars[0] = "Sandero"
	cars[1] = "Celta"
	cars[2] = "March"
	cars[3] = "320i"

	// Não é possível aumentar o array uma vez que ele foi criado
	// var newCars [6]string
	// cars = [5]string(newCars)

	// Se tentar dar mostrar informações de um espaço existente no array o código irá compilar, mesmo que esse espaço não possua valor atribuído 
	fmt.Println(cars[4])

	// Entretanto o mesmo não acontece caso tente acessar um valor de um espaço inexiste no array
	// fmt.Print(cars[5])

	fmt.Println(cars)
}