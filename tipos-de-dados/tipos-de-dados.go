package main

import (
	"errors"
	"fmt"
)

func main() {

	var num int = 10000000
	// o Go usa int8, int16, int32, int64 podendo ser criados int com essas quantidade de bits
	// se eu nao especificar o numero o go se baseia na arquitetura do meu computador
	// se meu computador for 32 bits o go criara int32 se for 64 bits criara int64
	fmt.Println(num)

	var num2 uint = 32
	fmt.Println(num2)
	// uint é o numero sem sinal (negativo ou positivo)

	var num3 float32 = 32.1
	var num4 float64 = 78019.2
	// para numeros float tenho que indicar obrigatoriamente se sera 32 ou 64 bits
	fmt.Println(num3)
	fmt.Println(num4)

	var str string = "Texto"
	fmt.Println(str)

	var booleano bool = true
	fmt.Println(booleano)

	// existe um tipo especificamente para representar erros no GO
	var erro1 error
	fmt.Println(erro1)

	// para criar um erro em go tenho que usar o pacote errors
	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}
