package main

import "fmt"

type ContaCorrente struct {
	titular string
	numeroAgencia int
	numeroConta int
	saldo float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo = 
	}
}

func main() {
	cliente1 := ContaCorrente{"Murilo", 531, 642332, 245.32}
	cliente2 := ContaCorrente{"Murilo", 531, 642332, 245.32}

	// var cliente2 *ContaCorrente
	// cliente2 := new(ContaCorrente)
	// cliente2.titular = "Cris"
	// cliente2.saldo = 500


	fmt.Println(&cliente1 == &cliente2)
	fmt.Println(cliente1 == cliente2)
}
