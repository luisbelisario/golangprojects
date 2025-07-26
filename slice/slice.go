package main

import "fmt"

func main() {

	// slice é como o ArrayList do Java, array sem determinacao de tamanho maximo
	slice := []string{}

	slice = append(slice, "Luis")

	fmt.Println(slice[0])

	fmt.Println(slice[1]) // ERRO

}
