package main

import (
	"fmt"
	"importModules/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Ola Mundo")

	auxiliar.Escrever()

	error := checkmail.ValidateFormat("mateus.cazuza@gmail.com")

	fmt.Println(error)
}
