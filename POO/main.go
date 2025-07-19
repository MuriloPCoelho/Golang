package main

import (
	"fmt"
)

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) {
	if valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		fmt.Printf("Saque realizado com sucesso. O saldo atual é de: %.2f\n", c.saldo)
	} else {
		fmt.Println("Saldo indisponível")
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		fmt.Printf("Deposito realizado com sucesso. O saldo atual é de: %.2f\n", c.saldo)
	} else {
		fmt.Println("Valor de deposito inválido")
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
	cliente1.Sacar(200)
	cliente1.Depositar(-500)
	fmt.Println(cliente1)
}
