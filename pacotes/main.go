package main

import (
	"fmt"
	"my-module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Printando da main")
	auxiliar.Printar()

	erro := checkmail.ValidateFormat("luis")
	fmt.Println(erro)
}
