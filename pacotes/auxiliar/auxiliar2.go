package auxiliar

import "fmt"

func printar2() {
	fmt.Println("Printando da auxiliar 2")
	// como essa função esta escrita com inicial minuscula
	// ela esta invisivel para fora do pacote
	// essa eh a forma do go dizer que uma função eh privada
	// so consigo acessa-la dentro do pacote auxiliar
}
