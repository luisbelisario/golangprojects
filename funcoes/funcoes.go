package main

import "fmt"

// depois dos parametros vem o retorno da funcao
func somar(n1 int, n2 int) int {
	return n1 + n2
}

// é possivel que uma funcao do go retorne varios valores
// aqui estou dizendo que essa funcao retorna dois inteiros
func calculosMatematicos(n1 int, n2 int) (int, int) {
	var soma int = n1 + n2
	var sub int = n1 - n2
	return soma, sub
}

// é possivel tambem retornar tipos diferentes
func tiposDiferentes() (string, bool) {
	return "tipos diferentes", true
}

func main() {
	var soma int = somar(2, 3)
	fmt.Println(soma)

	// é possivel atribuir funcoes a variaveis no Go

	// se a funcao nao retornar nada (void) nao coloco nada no retorno
	var funcao = func() {
		fmt.Println("Oi")
	}

	funcao()

	var funcaoComParametros = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := funcaoComParametros("Dentro da funcao com parametros")
	fmt.Println(resultado)

	var soma2, sub int = calculosMatematicos(10, 5)
	fmt.Println(soma2, sub)

	stringDafuncao, booleanoDaFuncao := tiposDiferentes()
	fmt.Println(stringDafuncao, booleanoDaFuncao)

	// se eu nao quiser usar uma das variaveis retornadas posso usar o underline para recebe-la
	stringDafuncao2, _ := tiposDiferentes()
	fmt.Println(stringDafuncao2)
}
