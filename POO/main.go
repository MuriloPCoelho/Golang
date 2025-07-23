package main

import (
	"example/hello/POO/clientes"
	"example/hello/POO/contas"
	"fmt"
)

func main() {
	cliente1 := contas.ContaCorrente{Titular: clientes.Titular{Nome: "Murilo", CPF: "1231231112", Profissao: "Analista Juridico"}, NumeroAgencia: 531, NumeroConta: 642332, Saldo: 245.32}
	clienteJairo := clientes.Titular{Nome: "Jairo", CPF: "1231231112", Profissao: "Politico"}
	contaJairo := contas.ContaCorrente{Titular: clienteJairo, NumeroAgencia: 531, NumeroConta: 642332, Saldo: 245.32}

	fmt.Println(&cliente1 == &contaJairo)
	fmt.Println(cliente1 == contaJairo)
	cliente1.Sacar(100)
	cliente1.Depositar(-500)

	cliente1.Transferir(100, &contaJairo)
	fmt.Println(cliente1)
	fmt.Println(contaJairo)
}
