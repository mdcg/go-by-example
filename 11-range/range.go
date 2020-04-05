package main

// O "range" itera sobre elementos de uma variedade de estruturas de dados.
// Veremos como utilizar o "range" em algumas das estruturas de dados que
// já aprendemos.

import "fmt"

func main() {
	// Aqui nós utilizamos o "range" para somar números em um "slice".
	// "arrays" funcionam da mesma maneira.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// Utilizar o "range" em "arrays" e "slices" retorna tanto o indice quanto
	// o valor de cada entrada. No exemplo acima, nós não precisamos do indice,
	// então o ignoramos com o "blank identifier". Às vezes nós realmente iremos
	// precisar do indice.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// o "range" em "maps" itera sobre os pares chave/valor.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Uma coisa interessante, é que o "range" também pode iterar
	// apenas nas chaves de um "map".
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// O "range" em strings itera sobre o valor do unicode da letra.
	// O primeiro valor é o indice, o segundo é a "rune" propriamente dita.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
