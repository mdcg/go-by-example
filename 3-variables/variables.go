package main

// Em Go, variáveis são explicitamente declaradas e usadas
// pelo compilador para, por exemplo, checar a correção do
// tipo de chamadas de função.

import "fmt"

func main() {
	// "var" declara uma ou mais variáveis.
	var a = "initial"
	fmt.Println(a)

	// você pode declarar múltiplas variáveis de uma só vez.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go irá inferir o tipo das variáveis inicializadas.
	var d = true
	fmt.Println(d)

	// Variáveis declaradas sem um valor correspondente de
	// inicialização são "zero-valued". Por exemplo, o "zero-value"
	// para o inteiro é 0.
	var e int
	fmt.Println(e)

	// A sintaxe := é um atalho para declarar e inicializar a variável,
	// por exemplo, para var f string = "apple" no caso abaixo.
	f := "apple"
	fmt.Println(f)
}
