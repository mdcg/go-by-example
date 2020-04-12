package main

// Quando usamos "channel" como parâmetros de funções, podemos
// especificar se o "channel" deve apenas enviar ou receber valores.
// Essa especificidade aumenta o "type-safety" do programa.
import "fmt"

// A função abaixo aceita apenas um canal para enviar valores.
// Seria um erro em tempo de compilação tentar receber valores por
// este canal.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// A função pong, aceita um canal para recebimento (pings)
// e outro para envio (pong).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
