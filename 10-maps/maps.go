package main

// "maps" são tipos associativos de dados em Go (Comumente chamado de
// "hashes" ou "dicts" em outras linguagens).

import "fmt"

func main() {
	// Para criar um "map" vazio, nós utilizamos a builtin "make":
	// make(map[key-type]val-type)
	m := make(map[string]int)

	// Atribui um conjunto chave/valor usando a típica sintaxe
	// "name[key] = val"
	m["k1"] = 7
	m["k2"] = 13

	// Quando você printar um "map", com um fmt.Println por exemplo,
	// ele irá mostrar todos os conjuntos chave/valor do "map" em questão.
	fmt.Println("map:", m)

	// Caso você queira o valor de uma chave, basta utilizar
	// o comando abaixo.
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Analogamente aos "arrays" e "slices", o builtin "len" irá trazer
	// o número de conjuntos chave/valor em um "map".
	fmt.Println("len:", len(m))

	// Para excluir um valor de um "map", utilizamos o builtin "delete".
	delete(m, "k2")
	fmt.Println("map:", m)

	// O segundo valor de retorno (que é opcional) ao tentarmos obter
	// o valor de uma chave de um "map", indica se a chave existe na mesma.
	// Como não precisamos do valor em si, o ignoramos com o identificador
	// em branco (blank identifier) "_".
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Bem parecido com "arrays" e "slices", também podemos declarar
	// e inicializar um novo "map" em uma única linha.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
