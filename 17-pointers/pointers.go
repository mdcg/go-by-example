package main

// Go suportar ponteiros, permitindo que você passe referências
// a valores e registros dentro do seu programa.

import "fmt"

// Nós iremos mostrar como ponteiros funcionam em contraste
// com valores com duas funções: "zeroval" e "zeroptr".
// "zeroval" possui um parâmetro int, portanto, os argumentos
// serão passados a ele por valor. O zeroval obterá uma cópia
// do ival distinta daquela na função chamada.
func zeroval(ival int) {
	ival = 0
}

// O "zeroptr", por outro lado, possui um parâmetro *int, o que
// significa que é necessário um ponteiro int. O código *iptr no corpo
// da função, desreferencia o ponteiro do endereço de memória para o valor
// atual nesse endereço. Atribuir um valor a uma ponteiro não referenciado
// altera o valor do endereço referenciado.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial", i)

	zeroval(i)
	fmt.Println("zeroval", i)

	// A sintaxe &i retorna o endereço de memória de i, ou seja, um
	// ponteiro para i.
	zeroptr(&i)
	fmt.Println("zeroptr", i)

	// Ponteiros também podem ser printados.
	fmt.Println("pointer", &i)

	// A função "zeroval" não muda o i principal, porém "zeroptr" o faz
	// porque ele possuia a referência de memória para essa variável.
}
