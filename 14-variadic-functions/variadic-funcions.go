package main

// "variadic functions", ou funções variáveis, são invocadas
// com qualquer número de argumentos. Por exemplo, o famoso
// fmt.Println é uma função variável comum.
import "fmt"

// Um exemplo de uma função variável é essa abaixo.
// Ela irá receber um número arbitrário de inteiros
// como argumento.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	// As funções variáveis podem ser chamadas da maneira
	// usal com os elementos individuais.
	sum(1, 2)
	sum(1, 2, 3)

	// Se você tiver múltiplos argumentos em um "slice",
	// para aplicá-los em uma função variável, nós precisamos
	// utilizar a sintaxe func(slice...), como no exemplo abaixo.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
