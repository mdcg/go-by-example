package main

// Em Go, um "array" é uma sequencia numerada de elementos com um comprimento específico.
import "fmt"

func main() {

	// Aqui nós estamos criando uma "array" que irá conter exatamente 5 inteiros.
	// O tipo dos elementos e o comprimento fazem parte do tipo do "array".
	// Por padrão um "array" tem valor zero, o que para inteiros significam que todos
	// os elementos serão 0s.
	var a [5]int
	fmt.Println("emp:", a)

	// Nós podemos atribuir um valor a um indice do array utilizando a sintaxe "array[indice] = valor".
	// De maneira similar, podemos pegar um valor de um indice do array utilizando "array[indice]".
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// A função "builtin" "len" retorna o comprimento do "array".
	fmt.Println("len:", len(a))

	// Utilize a sintaxe abaixo para declarar e inicializar um "array" em uma linha.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// "arrays" são uni-dimensionais, contudo, você pode compor tipos para criar estruturas multi-dimensionais.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
