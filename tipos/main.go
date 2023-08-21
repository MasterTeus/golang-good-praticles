package main

import (
	"errors"
	"fmt"
)

func main() {

	//TODO: INTEIROS

	//Intervalo: -128 a 127. int8 é o conjunto de todos os inteiros de 8 bits assinados.
	const inteiro8 int8 = 127
	//Intervalo: -32768 a 32767. int16 é o conjunto de todos os inteiros de 16 bits assinados.
	const inteiro16 int16 = 30000
	//Intervalo: -2147483648 a 2147483647. int32 é o conjunto de todos os inteiros de 32 bits assinados.
	const inteiro32 int32 = 2147483647
	//Faixa: -9223372036854775808 a 9223372036854775807. int64 é o conjunto de todos os inteiros de 64 bits assinados.
	const inteiro64 int64 = 9223372036854775807

	inteiro := 16

	println("Inteiros", inteiro8,
		inteiro16,
		inteiro32,
		inteiro64,
		inteiro)

	//TODO: FLUTUANTES

	//float32 é o conjunto de todos os números de ponto flutuante IEEE-754 de 32 bits.
	const flutuante32 float32 = 127.00
	//float64 é o conjunto de todos os números de ponto flutuante IEEE-754 de 64 bits.
	const flutuante64 float64 = 12700000.00

	flutuante := 12700000.01

	println(flutuante32, flutuante64, flutuante)

	//TODO: STRINGS

	var texto1 string = "Sou uma string"
	texto2 := "Sou uma string"

	// aspas simples retorna o valor numérico da string na tabela ASCII
	char := 'B'

	println(texto1, texto2, char)

	//TODO: ERROR

	var erro error = errors.New("Error aqui")

	fmt.Println(erro)

}
