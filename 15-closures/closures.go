package main

// Go suporta "funções anônimas", que podem formar "closures".
// Funções anônimas são úteis quando você precisa definir
// uma função embutida sem a necessidade de nomeá-la.
import "fmt"

// Essa função intSeq retorna outra função, que definimos
// "anonimamente" no corpo de intSeq. A função retornada
// "fecha" sobre a variável i para formar um "closure".
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// Nós chamamos o intSeq, atribuindo o retorno (uma função) ao nextInt.
	// Esse valor da função "captura" seu próprio valor i, que será atualizado
	// sempre que chamarmos o nextInt.
	nextInt := intSeq()

	// Veja o efeito de uma closure chamando o nextInt algumas vezes.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Para comprovar que o estado é único para essa função em particular, vamos
	// criar e testar uma nova.
	newInts := intSeq()
	fmt.Println(newInts())
}
