package main

// "structs" de Go são coleções tipadas de campos. Elas são
// úteis para agrupar dados para formar registros.

import "fmt"

// Essa "struct" do tipo pessoa possui campos de nome e idade.
type person struct {
	name string
	age  int
}

// A função abaixo, retorna um "struct" person com o nome
// passado como parâmetro.
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	// Você pode retornar com segurança um ponteiro para a variável local
	// pois uma variável local sobreviverá ao escopo da função.
	return &p
}

func main() {
	// Essa sintaxe cria uma nova "struct".
	fmt.Println(person{"Bob", 20})
	// Você pode nomear os campos quando for inicializar um nova "struct".
	fmt.Println(person{name: "Bob", age: 30})
	// Campos omitidos serão "zero-valued".
	fmt.Println(person{name: "Fred"})
	// Utilizando o prefixo &, você irá criar um ponteiro para um "struct".
	fmt.Println(&person{name: "Ann", age: 40})
	// É idiomático encapsular a criação de uma nova "struct" em funções
	// construtoras.
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	// Você pode acessar campos de uma "struct" utilizando o ponto.
	fmt.Println(s.name)

	// Você também pode usar o ponto em ponteiros de "structs", o ponteiro
	// será automaticamente desreferenciado.
	sp := &s
	fmt.Println(sp.age)

	// "structs" são mutáveis.
	sp.age = 51
	fmt.Println(sp.age)
}
