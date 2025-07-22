package main

import (
	"example/hello/POO/contas"
	"fmt"
)

func main() {
	cliente1 := contas.ContaCorrente{Titular: "Murilo", NumeroAgencia: 531, NumeroConta: 642332, Saldo: 245.32}
	cliente2 := contas.ContaCorrente{Titular: "Murilo", NumeroAgencia: 531, NumeroConta: 642332, Saldo: 245.32}

	fmt.Println(&cliente1 == &cliente2)
	fmt.Println(cliente1 == cliente2)
	cliente1.Sacar(100)
	cliente1.Depositar(-500)

	cliente1.Transferir(100, &cliente2)
	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
