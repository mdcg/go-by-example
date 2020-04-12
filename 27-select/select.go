package main

// O "select" permite que você aguarde várias operações em um
// "channel". A combinação de "goroutines" e "channels" com o
// "select" é um recurso poderoso em Go.
import (
	"fmt"
	"time"
)

func main() {
	// Para o nosso exemplo, iremos utilizar
	// o "select" através de dois "channels".
	c1 := make(chan string)
	c2 := make(chan string)

	// Cada "channel" irá receber um valor depois
	// de um certo periodo de tempo, para simular
	// operações RPC bloqueantes executando concorrentemente
	// em "goroutines".
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Nós utilizaremos o "select" para aguardar esses dois
	// valores simultaneamente, imprimindo cada um quando chegar.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}
