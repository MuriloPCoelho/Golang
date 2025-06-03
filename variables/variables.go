package main

import (
	"fmt"
	"log"
	"strconv"
)

/*	
	DATA TYPES
		bool -> booleanos (true ou false)
		string -> textos
		int,int8/16/32/64 -> integer/números inteiros (zero, positivos e negativos)
		byte -> equivalente a uint8 (usado para strings)
		uint, uint8/16/32/64 -> integer (zero e positivos)
		rune -> equivalente ao int32 (usado para unicode)
		float32,float64 -> float/números flutuantes (com casas decimais)
		complex64/128 -> parte real e imaginária em float32/64
		uintptr -> tipo inteiro sem sinal para representar endereços (em geral, não usar)
*/
func main() {
	const ATIVO string = "true"
	boolValue, err := strconv.ParseBool(ATIVO)
	if (err != nil) {
		log.Fatal(err)
	}
	nome := "Coelho"
	ativo := boolValue
	idade := 23

	fmt.Printf("Nome: %s \nIdade: %d \nAtivo: %s", nome, idade, strconv.FormatBool(ativo))
}