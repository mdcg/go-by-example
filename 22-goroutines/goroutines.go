package main

// Uma "goroutine" é uma thread de execução muito leve.
import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Suponha que tenhamos uma chamada de função "f(s)".
	// Veja como chamamos isso da maneira usual, executando
	// da forma síncrona.
	f("direct")

	// Para invocar uma função em "goroutine", use "go f(s)".
	// Essa nova "goroutine" será executada simultaneamente com a chamada.
	go f("goroutine")

	// Você também pode iniciar uma "goroutine" por uma chamada
	// de função anônima.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Nossas duas chamadas de função estão sendo executadas
	// assincronamente em goroutines separadas. Aguarde o término
	// (para uma abordagem mais robusta, use um WaitGroup)
	time.Sleep(time.Second)
	fmt.Println("done")

	// Quando executamos este programa, vemos primeiro a saída da
	// chamada de bloqueio (sincrono), depois, a saída intercalada
	// das duas "goroutines". Essa intercalação reflete as goroutines
	// sendo executadas simultaneamente em tempo de execução.
}
