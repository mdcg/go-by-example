package main

// Go suporta constantes de caractéres, strings, booleanos e valores numéricos.

import (
	"fmt"
	"math"
)

// "const" declara um valor contante.
const s string = "constant"

func main() {
	fmt.Println(s)

	// Uma declaração de "const" pode aparecer em qualquer lugar, da mesma forma
	// que uma declaração "var" pode.
	const n = 500000000

	// Expressões contantes desempenham valores aritméticos com precisão arbitrária.
	const d = 3e20 / n
	fmt.Println(d)

	// Uma constante numérica não tem tipo até receber um, como por meio de uma
	// conversão explícita.
	fmt.Println(int64(d))

	// Um número pode receber um tipo usando-o em um contexto que requer um, como
	// uma atribuição de variável ou chamada de função. Por exemplo, o math.Sin espera
	// um float64.
	fmt.Println(math.Sin(n))
}
