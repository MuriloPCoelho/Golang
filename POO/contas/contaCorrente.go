package contas

import (
	"example/hello/POO/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) {
	if valorDoSaque <= c.Saldo {
		c.Saldo -= valorDoSaque
		fmt.Printf("Saque realizado com sucesso. O Saldo atual é de: %.2f\n", c.Saldo)
	} else {
		fmt.Println("Saldo indisponível")
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) {
	if valorDeposito > 0 {
		c.Saldo += valorDeposito
		fmt.Printf("Deposito realizado com sucesso. O Saldo atual é de: %.2f\n", c.Saldo)
	} else {
		fmt.Println("Valor de deposito inválido")
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) {
	if c.Saldo > valorTransferencia && valorTransferencia > 0 {
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		fmt.Println("Transferência realizada com sucesso")
	} else {
		fmt.Println("Não foi possível realizar a transferência")
	}
}
