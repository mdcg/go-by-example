package main

// Os "timeouts" são importantes para programas que se conectam
// a recursos externos ou que, de alguma forma, precisam limitar
// o tempo de execução. A implementação de "timeouts" é simples,
// fácil e elegante, graças aos canais e select.
import (
	"fmt"
	"time"
)

func main() {
	// Para o nosso exemplo, suponha que estamos executando uma
	// chamada externa que retorna seu resultado em um "channel"
	// c1 após 2 segundos. Observe que o canal é armazenado em um "buffer"
	// e, portanto, o envio da "goroutine" não é bloqueado. Esse é um padrão
	// comum para evitar vazamentos de "goroutines", caso o canal nunca
	// seja lido.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// A seguir, o a implementação do "select" com "timeout".
	// "res:=<-Time.After" aguarda o envio após o tempo limite
	// que nesse caso, é de 1 segundo. Como o "select" prossegue
	// com o primeiro recebimento que estiver "pronto", o "case" "timeout"
	// irá acontecer se a operação demorar mais do que os 1s permitidos.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// Se permitirmos um tempo de "timeout" maior que 3 segundos, então o recebimento
	// do c2 será bem-sucedido e imprimiremos o resultado.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
