package main

// Go suporta por padrão o retorno múltiplo de valores. Essa feature
// é utilizada com frequência em códigos Go, por exemplo, para retornar
// tanto o valor do resultado quanto o do erro em uma função.

import "fmt"

// O (int, int) na assinatura da função mostra que a função
// retornará dois inteiros.
func vals() (int, int) {
	return 3, 7
}

func main() {
	// Aqui nós utilizamos atribuímos os dois diferentes valores
	// de retorno da chamada utilizando o "multiple assignment".
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Caso você deseje utilizar apenas um dos valores do retorno,
	// use o blank identifier.
	_, c := vals()
	fmt.Println(c)
}
