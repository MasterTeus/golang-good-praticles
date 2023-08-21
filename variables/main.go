package main

import (
	"fmt"
)

func main() {

	//TODO: VARIÁVEIS

	//criação de variável com tipagem explícita
	const variable1 string = "Olá sou uma variável"
	//Com inferência de tipos
	variable2 := "Olá sou uma variável 2"

	// chamando um valor destro de um logger
	fmt.Println(variable1, variable2)

	//TODO: Criação em série

	//EXPLÍCITAS
	var (
		name string = "Mateus"
		age  int    = 26
	)
	println(name, age)

	//IMPLÍCITAS (POR INFERÊNCIA DE TIPO)
	name2, age2 := "Mateus Cazuza", 27
	fmt.Println(name2, age2)

}
