package main

// Go suporta funções recursivas. Iremos demonstrar aqui
// um exemplo clássico utilizando o fatorial.
import "fmt"

// Essa função retorna ela mesma até alcançar
// o caso de parada, que nesse exemplo, seria
// fact(0).
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
