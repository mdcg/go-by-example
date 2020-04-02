package main

// O "for" é a única estrutura de loop em Go. Abaixo teremos três tipos
// básicos de tipos de loops "for".

import "fmt"

func main() {
	// O tipo mais básico, com uma única condição.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Um clássico initial/condition/after.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Um "for" sem condição irá rodar repetidamente até você pará-lo utilziando
	// um "break" ou "return" da "função de fechamento" (enclosing function)
	for {
		fmt.Println("loop")
		break
	}

	// Você também pode utilização o "continue" para ir para a próxima iteração
	// do loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
