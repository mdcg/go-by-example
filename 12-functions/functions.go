package main

// Funções são peças-chave em Go. Iremos aprender um pouco
// sobre elas com alguns exemplos diferentes.

import "fmt"

// Aqui estamos declarando uma função que recebe dois inteiros
// e retorna a soma entre eles, que por sinal, também é um inteiro.
func plus(a int, b int) int {
	// Go requer um retorno explicito, ou seja, ela não irá retornar
	// automaticamente o valor da última expressão.
	return a + b
}

// Quando temos múltiplos parâmetros consecutivos de um mesmo tipo,
// podemos omitir o nome do tipo até o parâmetro final que declara o tipo.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	// Para invocarmos a função como você esperava, basta utilizarmos
	// a sintaxe name(args).
	res := plus(1, 2)
	fmt.Println("1+2 = ", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 = ", res)
}
