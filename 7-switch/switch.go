package main

// instruções "switch" expressão condicionais em vários ramos.

import (
	"fmt"
	"time"
)

func main() {
	// Um exemplo de "switch" simples.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Você pode utilizar virgulas para separar múltiplas expressões
	// em uma mesma declaração de "case". Nós utilizamos para esse exemplo
	// um "default" que é opcional.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// "switch" sem uma expressão é um jeito alternativo de expressar
	// a lógica if/else. Aqui mostramos também como o "switch" pode ser
	// "não constante".
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// um "tipo switch" compara tipos ao invés de valores.
	// Você pode usar isso para descobrir o tipo de um valor
	// de interface. Neste exemplo, a variável "t" terá o tipo
	// correspondente à sua cláusula.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey!")
}
