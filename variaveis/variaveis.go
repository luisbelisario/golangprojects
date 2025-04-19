package main

// para rodar uma aplicacao na linha de comando
// tenho que obrigatoriamente definir aquele pacote como o main

import "fmt"

func main() {

	var var1 string = "var1"
	var2 := "var2"

	// mesmo com a tipagem implicita nao posso alterar o tipo da variavel
	// nao posso fazer var2 = 2

	fmt.Println(var1, var2)

	var (
		var3 string = "var3"
		var4 string = "var4"
	)

	fmt.Println(var3, var4)

	var5, var6 := "var5", "var6"

	fmt.Println(var5, var6)

	const const1 string = "const1"
	// constantes não geram erro de compilacao se nao forem usadas
	// variaveis geram erro de compilacao se nao forem usadas

	var5, var6 = var6, var5

	fmt.Println(var5, var6)
}
