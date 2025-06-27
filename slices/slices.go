package main

import (
	"fmt"
	"reflect"
)

func main() {
	exibeNomes()
}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Bernardo"} //slice é uma abstração do array
	fmt.Println(reflect.TypeOf(nomes)) // -> func que permite ver o tipo de uma variável
	nomes = append(nomes, "Augusto") //slice permite adicionar mais itens ao array
	//por baixo dos panos o slice é uma abstração que dobra a capacidade do array para permitir a inserção de mais itens
	fmt.Println(nomes)
}