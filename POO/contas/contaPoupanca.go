package contas

import (
	"example/hello/POO/clientes"
	"fmt"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) {
	if valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		fmt.Printf("Saque realizado com sucesso. O saldo atual é de: %.2f\n", c.saldo)
	} else {
		fmt.Println("saldo indisponível")
	}
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		fmt.Printf("Deposito realizado com sucesso. O saldo atual é de: %.2f\n", c.saldo)
	} else {
		fmt.Println("Valor de deposito inválido")
	}
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) PagarBoletoContaCorrente(valorBoleto float64) {
	if c.saldo > valorBoleto {
		c.saldo -= valorBoleto
		fmt.Println("Boleto pago com sucesso")
	} else {
		fmt.Println("Saldo insuficiente")
	}
}
