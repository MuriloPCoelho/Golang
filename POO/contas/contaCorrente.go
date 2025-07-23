package contas

import (
	"example/hello/POO/clientes"
	"fmt"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) {
	if valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		fmt.Printf("Saque realizado com sucesso. O saldo atual é de: %.2f\n", c.saldo)
	} else {
		fmt.Println("saldo indisponível")
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

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) {
	if c.saldo > valorTransferencia && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		fmt.Println("Transferência realizada com sucesso")
	} else {
		fmt.Println("Não foi possível realizar a transferência")
	}
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
