package main

// "slices" são tipos de dados chave no Go, fornecendo uma
// interface mais poderosa para sequência do que "arrays".

import "fmt"

func main() {
	// Diferentemente dos "arrays", os "slices" podem crescer
	// indefinidamente. Por exemplo, quando um "array" declara que
	// possui três elementos, esse por sua vez vai possuir sempre
	// três elementos em sua estrutura. Já um "slice", pode inicializar
	// com três elementos, mas isso não é um fator que determina que ele
	// terá no máximo essa quantidade.

	// Para criar um "slice" vazio com comprimento zero, use a função
	// builtin "make". Abaixo fazemos um slice de comprimento 3 (inicialmente
	// com valor zero).
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Algumas operações são similares aos "arrays", como por exemplo
	// para atribuir e capturar um elemento do "slice".
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// Como esperado, a função "len" retorna
	// o comprimento do "slice".
	fmt.Println("len:", len(s))

	// Além das operações básicas, os "slices" possuem várias outras,
	// que a tornam mais ricas que os "arrays". Uma delas é a função
	// builtin "append", que retorna um slice contendo um ou mais novos
	// valores. Perceba que o "append" retorna um novo "slice", dessa
	// forma, devemos fazer a reatribuição do slice em questão.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Os "slices" também podem ser copiados. Para isso, precisamos
	// criar um slice vazio com o mesmo comprimento do "slice" que
	// queremos copiar, e então, usamos a função builtin "copy"
	// para fazermos a copia do "slice".
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Um "slice" consegue fatiar ele mesmo utilizando a sintaxe
	// slice[low:high]. Por exemplo, o código abaixo irá pegar os
	// elementos do slice[2], slice[3] e slice[4]. (Excluindo o
	// elemento slice[5]).
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Analogamente, aqui vai pegar do slice[0] - slice[4],
	// lembrando novamente, excluindo o último elemento "high".
	l = s[:5]
	fmt.Println("sl2:", l)

	// Contudo, quando um elemento "low" é definido, nós iremos
	// incluí-lo em nossa fatia resultante. Dessa maneira, o
	// código abaixo teremos slice[2] - slice[6] (último elemento).
	l = s[2:]
	fmt.Println("sl3:", l)

	// Assim como "arrays", nós podemos declarar e inicializar uma
	// variável que será um "slice" em apenas uma linha.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Os "slices" também podem ser compostos em estrutuas de dados
	// multidimensionais. Uma coisa interessante é que o comprimento
	// das fatias internas pode variar, diferente dos "arrays"
	// multidimensionais.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// PS: Por mais que "slices" e "arrays" sejam tipos diferentes,
	// eles são impressos da mesma forma quando utilizamos o comando
	// fmt.Println.
}
