package main

// A "ramificação" com "if" e "else" em Go é direta.

import "fmt"

func main() {
	// Um exemplo básico.
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Você pode ter uma declaração "if" sem um "else".
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Uma declaração pode preceder condicionais. Quaisquer variáveis
	// declaradas nesta instrução estão disponíveis em todas as ramificações.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// Note que você não precisa de parenteses ao redor de condicionais em Go.
	// Contudo, as chaves são necessárias.

	// Não há operador ternário em Go, então você precisar utilizar a declaração
	// completa das condicionais até mesmo em condições básicas.
}
